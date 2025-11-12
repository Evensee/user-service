package user

import "github.com/google/uuid"

type CreateUser struct {
	Email     string  `json:"email" validate:"required"`
	FirstName string  `json:"firstName" validate:"required"`
	LastName  string  `json:"lastName" validate:"required"`
	AvatarUrl *string `json:"avatarUrl"`
	Password  *string `json:"password" validate:"required"`
}

type UpdateUser struct {
	ID        *uuid.UUID `json:"id" `
	Email     *string    `json:"email" `
	FirstName *string    `json:"firstName" `
	LastName  *string    `json:"lastName" `
	AvatarUrl *string    `json:"avatarUrl"`
	Password  *string    `json:"password" `
}

type FindUser struct {
	ID        *uuid.UUID `json:"id"`
	Email     *string    `json:"email" `
	FirstName *string    `json:"firstName" `
	LastName  *string    `json:"lastName" `
}
