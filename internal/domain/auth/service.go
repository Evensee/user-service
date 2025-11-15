package auth

import (
	"context"

	"github.com/Evensee/user-service/internal"
	"github.com/Evensee/user-service/internal/domain/user"
	"github.com/Evensee/user-service/internal/lib/jwt"
	"github.com/Evensee/user-service/internal/lib/security"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Ctx = context.Context

type AuthService struct {
	tokenRepo Repository
	userRepo  user.Repository
	appConfig *internal.AppConfig
}

func NewAuthService(
	tokenRepo Repository,
	userRepo user.Repository,
	appConfig *internal.AppConfig,
) *AuthService {
	return &AuthService{
		tokenRepo: tokenRepo,
		userRepo:  userRepo,
		appConfig: appConfig,
	}
}

func (s *AuthService) LoginUser(ctx Ctx, email, password string) (Tokens, error) {
	u, err := s.userRepo.GetOne(&user.FindUser{
		Email: &email,
	})
	if err != nil {
		return Tokens{}, err
	}

	if !security.VerifyPassword(password, u.HashedPassword) {
		return Tokens{}, status.Error(codes.Unauthenticated, "invalid password")
	}

	access, err := jwt.GenerateAccessToken(u, s.appConfig)
	if err != nil {
		return Tokens{}, err
	}

	refresh, err := jwt.GenerateRefreshToken(u, s.appConfig)
	if err != nil {
		return Tokens{}, err
	}

	tokens := Tokens{
		AccessToken:  access,
		RefreshToken: refresh,
	}

	if err := s.tokenRepo.Save(ctx, u.ID, tokens); err != nil {
		return Tokens{}, err
	}

	return tokens, nil
}

func (s *AuthService) RefreshTokens(ctx Ctx, refreshToken string) (Tokens, error) {
	userID, err := s.tokenRepo.GetUserIDByRefreshToken(ctx, refreshToken)
	if err != nil {
		return Tokens{}, err
	}

	userIdFromBlocklist, err := s.tokenRepo.CheckRefreshTokenBlocked(ctx, refreshToken)

	if userIdFromBlocklist != nil && err == nil {
		panic(status.Error(codes.Unauthenticated, "refresh token was blocked"))
	}

	u, err := s.userRepo.GetOne(&user.FindUser{
		ID: &userID,
	})
	if err != nil {
		return Tokens{}, err
	}

	newAccess, err := jwt.GenerateAccessToken(u, s.appConfig)
	if err != nil {
		return Tokens{}, err
	}

	newRefresh, err := jwt.GenerateRefreshToken(u, s.appConfig)
	if err != nil {
		return Tokens{}, err
	}

	newTokens := Tokens{
		AccessToken:  newAccess,
		RefreshToken: newRefresh,
	}

	if err := s.tokenRepo.Save(ctx, u.ID, newTokens); err != nil {
		return Tokens{}, err
	}

	_ = s.tokenRepo.DeleteRefreshToken(ctx, refreshToken)
	s.tokenRepo.BlockRefreshToken(ctx, refreshToken, userID)

	return newTokens, nil
}

func (s *AuthService) LogoutUser(ctx Ctx, accessToken, refreshToken string) error {
	_ = s.tokenRepo.DeleteAccessToken(ctx, accessToken)
	_ = s.tokenRepo.DeleteRefreshToken(ctx, refreshToken)
	return nil
}

func (s *AuthService) ValidateAccessToken(ctx Ctx, accessToken string) (*user.User, error) {
	userId, err := jwt.ValidateAccessToken(accessToken, s.appConfig)
	if err != nil {
		return nil, err
	}

	u, err := s.userRepo.GetOne(&user.FindUser{
		ID: userId,
	})
	if err != nil {
		return nil, err
	}

	return u, nil
}
