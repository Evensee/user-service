package user

import (
	"github.com/google/uuid"
)

type Repository interface {
	CreateUser(*User) (*User, error)
	Update(uuid.UUID, *UpdateUser) (*User, error)
	GetAll(*FindUser) (*[]User, error)
	GetOne(*FindUser) (*User, error)
}
