package jwt

import (
	"errors"
	"time"

	"github.com/Evensee/user-service/internal"
	"github.com/Evensee/user-service/internal/domain/user"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func NewToken(
	user *user.User,
	appConfig *internal.AppConfig,
	duration time.Duration,
) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.ID
	claims["exp"] = time.Now().Add(duration).Unix()

	tokenString, err := token.SignedString([]byte(appConfig.Secret))

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GenerateAccessToken(u *user.User, appConfig *internal.AppConfig) (string, error) {
	accessTokenLifetime := time.Second * time.Duration(appConfig.AccessTokenLifetimeSeconds)
	return NewToken(u, appConfig, accessTokenLifetime)
}

func GenerateRefreshToken(u *user.User, appConfig *internal.AppConfig) (string, error) {
	refreshTokenLifetime := time.Second * time.Duration(appConfig.RefreshTokenLifetimeSeconds)
	return NewToken(u, appConfig, refreshTokenLifetime)
}

func GenerateOAuthTokens(u *user.User, appConfig *internal.AppConfig) (
	accessToken,
	refreshToken string,
	err error,
) {
	accessToken, err = GenerateAccessToken(u, appConfig)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = GenerateRefreshToken(u, appConfig)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func ValidateAccessToken(
	tokenStr string,
	appConfig *internal.AppConfig,
) (
	*uuid.UUID,
	error,
) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&jwt.MapClaims{},
		func(token *jwt.Token) (any, error) {
			return []byte(appConfig.Secret), nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	userId, ok := claims["uid"].(string)
	if !ok {
		return nil, errors.New("invalid user id")
	}

	userUUID, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}

	return &userUUID, nil
}
