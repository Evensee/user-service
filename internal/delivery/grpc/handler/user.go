package handler

import (
	"context"

	"github.com/Evensee/user-service/internal/delivery/grpc/mapper"
	"github.com/Evensee/user-service/internal/domain/user"
	"github.com/Evensee/user-service/internal/interface/service"
	p "github.com/Evensee/user-service/protobuf_generated/user"
	"github.com/evensee/go-tl/api"
	"github.com/google/uuid"
)

type (
	Ctx = context.Context
)

func (s *ServerApi) GetUserById(ctx context.Context, req *p.GetUserByIdRequest) (*p.UserResponse, error) {
	handler := func(
		ctx Ctx,
		req *p.GetUserByIdRequest,
		appService service.AppService,
	) (*p.UserResponse, error) {
		userService := appService.GetUserService()

		userId, err := uuid.Parse(req.GetUserId())

		if err != nil {
			panic(err)
		}

		user, err := userService.GetOne(&user.FindUser{
			ID: &userId,
		})

		return mapper.MapUserDomainToGrpcModel(user), err
	}

	resolvedHandler := api.GrpcApiHandlerFactory(
		handler,
		s.appResolver.CreateAppTransaction,
		s.appResolver.CreateAppService,
	)

	return resolvedHandler(ctx, req)
}

func (s *ServerApi) CreateUser(ctx context.Context, req *p.CreateUserRequest) (*p.UserResponse, error) {
	handler := func(
		ctx Ctx,
		req *p.CreateUserRequest,
		appService service.AppService,
	) (*p.UserResponse, error) {
		userService := appService.GetUserService()

		createUserModel := mapper.MapCreateUserGrpcToDomainModel(req)
		user, err := userService.Create(createUserModel)
		if err != nil {
			panic(err)
		}
		return mapper.MapUserDomainToGrpcModel(user), nil
	}

	resolvedHandler := api.GrpcApiHandlerFactory(
		handler,
		s.appResolver.CreateAppTransaction,
		s.appResolver.CreateAppService,
	)

	return resolvedHandler(ctx, req)
}
