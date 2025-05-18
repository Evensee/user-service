package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserORMModel struct {
	gorm.Model

	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`

	Email string

	HashedPassword string
}

func (UserORMModel) TableName() string {
	return "users"
}
