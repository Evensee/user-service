package handler

import (
	"context"

	"github.com/Evensee/user-service/internal/interface/service"
	p "github.com/Evensee/user-service/protobuf_generated/user"
	"github.com/evensee/go-tl/api"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ServerApi) LoginUser(ctx context.Context, req *p.LoginRequest) (*p.TokensResponse, error) {
	handler := func(
		ctx Ctx,
		req *p.LoginRequest,
		appService service.AppService,
	) (*p.TokensResponse, error) {
		authService := appService.GetAuthService()

		tokens, err := authService.LoginUser(ctx, req.Email, req.Password)
		
		return &p.TokensResponse{
			AccessToken: tokens.AccessToken,
			RefreshToken: tokens.RefreshToken,
		}, err
	}

	resolvedHandler := api.GrpcApiHandlerFactory(
		handler,
		s.appResolver.CreateAppTransaction,
		s.appResolver.CreateAppService,
	)

	return resolvedHandler(ctx, req)
}
func (s *ServerApi) LogoutUser(ctx context.Context, req *p.LogoutRequest) (*emptypb.Empty, error) {
	handler := func(
		ctx Ctx,
		req *p.LogoutRequest,
		appService service.AppService,
	) (*emptypb.Empty, error) {
		authService := appService.GetAuthService()
		
		err := authService.LogoutUser(ctx, req.AccessToken, req.RefreshToken)
		
		return new(emptypb.Empty), err
	}

	resolvedHandler := api.GrpcApiHandlerFactory(
		handler,
		s.appResolver.CreateAppTransaction,
		s.appResolver.CreateAppService,
	)

	return resolvedHandler(ctx, req)
}
func (s *ServerApi) ValidateTokens(ctx context.Context, req *p.ValidateTokensRequest) (*p.ValidateTokensResponse, error) {
	handler := func(
		ctx Ctx,
		req *p.ValidateTokensRequest,
		appService service.AppService,
	) (*p.ValidateTokensResponse, error) {
		authService := appService.GetAuthService()
		
		authService.
	}

	resolvedHandler := api.GrpcApiHandlerFactory(
		handler,
		s.appResolver.CreateAppTransaction,
		s.appResolver.CreateAppService,
	)

	return resolvedHandler(ctx, req)
}
func (s *ServerApi) RefreshTokens(ctx context.Context, req *p.RefreshTokensRequest) (*p.TokensResponse, error) {
	handler := func(
		ctx Ctx,
		req *p.RefreshTokensRequest,
		appService service.AppService,
	) (*p.TokensResponse, error) {
		userService := appService.GetUserService()
	}

	resolvedHandler := api.GrpcApiHandlerFactory(
		handler,
		s.appResolver.CreateAppTransaction,
		s.appResolver.CreateAppService,
	)

	return resolvedHandler(ctx, req)
}
