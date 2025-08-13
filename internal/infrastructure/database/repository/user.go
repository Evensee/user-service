package repository

import (
	"context"
	"errors"
	"github.com/Evensee/user-service/internal/domain/user"
	"github.com/Evensee/user-service/internal/infrastructure/database/mapper"
	"github.com/Evensee/user-service/internal/infrastructure/database/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetById(ctx context.Context, id string) (user.User, error) {
	var userOrm model.UserORMModel

	uid, err := uuid.Parse(id)

	if err != nil {
		return user.User{}, err
	}

	if err := r.db.WithContext(ctx).First(&userOrm, "id = ?", uid).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user.User{}, user.ErrorUserNotFound
		}

		return user.User{}, err
	}

	return mapper.MapToDomain(userOrm), nil
}

func (r *UserRepository) Create(ctx context.Context, userDomain user.User) (user.User, error) {
	userOrm := mapper.MapToOrm(userDomain)

	if err := r.db.WithContext(ctx).Create(&userOrm).Error; err != nil {
		return user.User{}, err
	}

	return mapper.MapToDomain(userOrm), nil
}
