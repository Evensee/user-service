package api

import (
	"context"

	"github.com/Evensee/user-service/internal/interface/service"
	p "github.com/Evensee/user-service/protobuf_generated/user"
	"github.com/evensee/go-tl/api"
	"google.golang.org/protobuf/types/known/emptypb"
)

type (
	Ctx = context.Context
)

func (s *serverApi) GetUserById(ctx context.Context, req *p.GetUserByIdRequest) (*p.UserResponse, error) {
	handler := func(
		ctx Ctx,
		req *p.GetUserByIdRequest,
		appService service.AppService,
	) (*p.UserResponse, error) {
		userService := appService.GetUserService()
	}

	resolvedHandler := api.GrpcApiHandlerFactory(
		handler,
		s.appResolver.CreateAppTransaction,
		s.appResolver.CreateAppService,
	)

	return resolvedHandler(ctx, req)
}
func (s *serverApi) CreateUser(ctx context.Context, req *p.CreateUserRequest) (*p.UserResponse, error) {
	handler := func(
		ctx Ctx,
		req *p.CreateUserRequest,
		appService service.AppService,
	) (*p.UserResponse, error) {
		userService := appService.GetUserService()
	}

	resolvedHandler := api.GrpcApiHandlerFactory(
		handler,
		s.appResolver.CreateAppTransaction,
		s.appResolver.CreateAppService,
	)

	return resolvedHandler(ctx, req)
}
func (s *serverApi) LoginUser(ctx context.Context, req *p.LoginRequest) (*p.TokensResponse, error) {
	handler := func(
		ctx Ctx,
		req *p.LoginRequest,
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
func (s *serverApi) LogoutUser(ctx context.Context, req *p.LogoutRequest) (*emptypb.Empty, error) {
	handler := func(
		ctx Ctx,
		req *p.LogoutRequest,
		appService service.AppService,
	) (*emptypb.Empty, error) {
		userService := appService.GetUserService()
	}

	resolvedHandler := api.GrpcApiHandlerFactory(
		handler,
		s.appResolver.CreateAppTransaction,
		s.appResolver.CreateAppService,
	)

	return resolvedHandler(ctx, req)
}
func (s *serverApi) ValidateTokens(ctx context.Context, req *p.ValidateTokensRequest) (*p.ValidateTokensResponse, error) {
	handler := func(
		ctx Ctx,
		req *p.ValidateTokensRequest,
		appService service.AppService,
	) (*p.ValidateTokensResponse, error) {
		userService := appService.GetUserService()
	}

	resolvedHandler := api.GrpcApiHandlerFactory(
		handler,
		s.appResolver.CreateAppTransaction,
		s.appResolver.CreateAppService,
	)

	return resolvedHandler(ctx, req)
}
func (s *serverApi) RefreshTokens(ctx context.Context, req *p.RefreshTokensRequest) (*p.TokensResponse, error) {
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
