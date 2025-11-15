package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserORMModel struct {
	ID             *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	FirstName      *string
	LastName       *string
	Email          *string `gorm:"unique"`
	AvatarUrl      *string
	HashedPassword *string

	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *gorm.DeletedAt `gorm:"index"`
}

func (UserORMModel) TableName() string {
	return "users"
}
