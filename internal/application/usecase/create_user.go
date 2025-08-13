package usecase

import (
	"context"
	"github.com/Evensee/user-service/internal/application/contract"
	"github.com/Evensee/user-service/internal/domain/user"
)

type UserUseCase struct {
	userService *user.DomainUserService
}

func NewUserUseCase(userService *user.DomainUserService) *UserUseCase {
	return &UserUseCase{userService: userService}
}

func (userUseCase *UserUseCase) CreateUser(ctx context.Context, request contract.CreateUserContract) (contract.CreateUserResponse, error) {
	userEntity, err := user.NewUser(request.Email, request.Password)

	if err != nil {
		return contract.CreateUserResponse{}, err
	}

	createdUser, err := userUseCase.userService.Create(ctx, *userEntity)

	if err != nil {
		return contract.CreateUserResponse{}, err
	}

	return contract.CreateUserResponse{
		ID:    createdUser.ID,
		Email: createdUser.Email,
	}, nil
}
