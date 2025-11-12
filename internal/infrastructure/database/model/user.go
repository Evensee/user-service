package model

import (
	"database/sql"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserORMModel struct {
	gorm.Model

	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`

	FirstName string
	LastName  string

	Email string

	AvatarUrl sql.NullString

	HashedPassword string
}

func (UserORMModel) TableName() string {
	return "users"
}
