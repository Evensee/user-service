package mapper

import (
	"database/sql"

	"github.com/Evensee/user-service/internal/domain/user"
	"github.com/Evensee/user-service/internal/infrastructure/database/model"
)

func MapToDomain(userOrm *model.UserORMModel) *user.User {
	var avatarUrl *string

	if userOrm.AvatarUrl.Valid {
		avatarUrl = &userOrm.AvatarUrl.String
	}

	return &user.User{
		ID:             userOrm.ID,
		Email:          userOrm.Email,
		HashedPassword: userOrm.HashedPassword,
		FirstName:      userOrm.FirstName,
		LastName:       userOrm.LastName,
		AvatarUrl:      avatarUrl,
	}
}

func MapToOrm(userDomain *user.User) model.UserORMModel {
	return model.UserORMModel{
		ID:             userDomain.ID,
		Email:          userDomain.Email,
		HashedPassword: userDomain.HashedPassword,
		FirstName:      userDomain.FirstName,
		LastName:       userDomain.LastName,
		AvatarUrl: sql.NullString{
			String: *userDomain.AvatarUrl,
			Valid:  userDomain.AvatarUrl != nil,
		},
	}
}

func CreateUserToOrm(createUser *user.CreateUser) *model.UserORMModel {
	return &model.UserORMModel{
		Email:          createUser.Email,
		FirstName:      createUser.FirstName,
		LastName:       createUser.LastName,
		HashedPassword: *createUser.Password,
	}
}
