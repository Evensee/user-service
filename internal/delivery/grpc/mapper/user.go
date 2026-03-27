package mapper

import (
	"github.com/Evensee/user-service/internal/domain/user"
	p "github.com/Evensee/user-service/protobuf_generated/user"
)

func MapUserDomainToGrpcModel(user *user.User) *p.UserResponse {
	return &p.UserResponse{
		UserId:    user.ID.String(),
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		AvatarUrl: user.AvatarUrl,
	}
}

func MapCreateUserGrpcToDomainModel(req *p.CreateUserRequest) *user.CreateUser {
	return &user.CreateUser{
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		AvatarUrl: req.AvatarUrl,
		Password:  req.Password,
	}
}
