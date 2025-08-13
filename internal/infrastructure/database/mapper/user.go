package mapper

import (
	"github.com/Evensee/user-service/internal/domain/user"
	"github.com/Evensee/user-service/internal/infrastructure/database/model"
	"github.com/google/uuid"
)

func MapToDomain(userOrm model.UserORMModel) user.User {
	return user.User{
		ID:             userOrm.ID.String(),
		Email:          userOrm.Email,
		HashedPassword: userOrm.HashedPassword,
	}
}

func MapToOrm(userDomain user.User) model.UserORMModel {
	id, _ := uuid.Parse(userDomain.ID)

	return model.UserORMModel{
		ID:             id,
		Email:          userDomain.Email,
		HashedPassword: userDomain.HashedPassword,
	}
}
