package repository

import (
	"context"

	"github.com/Evensee/user-service/internal/domain/user"
	"github.com/Evensee/user-service/internal/infrastructure/database/mapper"
	"github.com/Evensee/user-service/internal/infrastructure/database/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
	user.Repository
}

type Ctx = context.Context

func NewUserRepository(db *gorm.DB) user.Repository {
	return UserRepository{
		db: db,
	}
}

func (r UserRepository) CreateUser(
	createUser *user.User,
) (*user.User, error) {
	userModel := mapper.MapToOrm(createUser)

	result := r.db.Create(&userModel)

	return mapper.MapToDomain(&userModel), result.Error
}

func (r UserRepository) GetUsers(
	findUser user.FindUser,
) (*[]user.User, error) {
	users := make([]model.UserORMModel, 0)

	result := r.db.Find(&users)

	if result.Error != nil {
		panic(result.Error)
	}

	mappedUsers := make([]user.User, 0, len(users))

	for _, u := range users {
		mappedUsers = append(
			mappedUsers,
			*mapper.MapToDomain(&u),
		)
	}

	return &mappedUsers, result.Error
}

func (r UserRepository) GetUser(userId uuid.UUID) (*user.User, error) {
	u := model.UserORMModel{
		ID: userId,
	}

	result := r.db.First(&u)

	if result.Error != nil {
		panic(result.Error)
	}

	return mapper.MapToDomain(&u), result.Error
}

func (r UserRepository) UpdateUser(
	userId uuid.UUID,
	updateUser *user.UpdateUser,
) (*user.User, error) {
	u := model.UserORMModel{
		ID: userId,
	}

	result := r.db.Model(&u).Where("id = ?", userId).Updates(updateUser)
	r.db.First(&u)

	return mapper.MapToDomain(&u), result.Error
}

func (r UserRepository) DeleteUser(
	userId uuid.UUID,
) error {
	u := model.UserORMModel{
		ID: userId,
	}
	r.db.First(&u)
	result := r.db.Delete(&u)

	return result.Error
}
