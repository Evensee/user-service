package api

import (
	"context"

	"github.com/Evensee/user-service/internal/delivery/grpc/handler"
	"github.com/Evensee/user-service/internal/interface/resolver"
	"github.com/Evensee/user-service/protobuf_generated/user"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Ctx = context.Context

type serverApi struct {
	user.UnimplementedUserServiceServer
	appResolver resolver.AppResolver
	handler     handler.ServerApi
}

func New(appResolver resolver.AppResolver) *serverApi {
	return &serverApi{
		appResolver: appResolver,
		handler:     handler.New(appResolver),
	}
}

func Register(gRPC *grpc.Server, serverApiInstance *serverApi) {
	user.RegisterUserServiceServer(
		gRPC,
		serverApiInstance,
	)
}

func (s *serverApi) CreateUser(
	ctx Ctx,
	req *user.CreateUserRequest,
) (*user.UserResponse, error) {
	return s.handler.CreateUser(ctx, req)
}

func (s *serverApi) GetUserById(
	ctx Ctx,
	req *user.GetUserByIdRequest,
) (*user.UserResponse, error) {
	return s.handler.GetUserById(ctx, req)
}

func (s *serverApi) LoginUser(
	ctx Ctx,
	req *user.LoginRequest,
) (*user.TokensResponse, error) {
	return s.handler.LoginUser(ctx, req)
}

func (s *serverApi) LogoutUser(
	ctx Ctx,
	req *user.LogoutRequest,
) (*emptypb.Empty, error) {
	return s.handler.LogoutUser(ctx, req)
}

func (s *serverApi) RefreshTokens(
	ctx Ctx,
	req *user.RefreshTokensRequest,
) (*user.TokensResponse, error) {
	return s.handler.RefreshTokens(ctx, req)
}

func (s *serverApi) ValidateTokens(
	ctx Ctx,
	req *user.ValidateTokensRequest,
) (*user.ValidateTokensResponse, error) {
	return s.handler.ValidateTokens(ctx, req)
}
