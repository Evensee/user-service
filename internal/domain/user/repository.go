package user

import (
	"context"

	"github.com/google/uuid"
)

type (
	Ctx = context.Context
)

type Repository interface {
	Create(Ctx, *User) (*User, error)
	Update(Ctx, uuid.UUID, *UpdateUser) (*User, error)
	GetAll(Ctx, *FindUser) (*[]User, error)
	GetOne(Ctx, *FindUser) (*User, error)
}
