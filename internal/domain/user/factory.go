package user

import (
	"github.com/Evensee/user-service/internal/lib/security"
	"github.com/google/uuid"
)

func NewUser(createUser *CreateUser) (*User, error) {
	hashedPassword, err := security.HashPassword(*createUser.Password)

	return &User{
		ID:             uuid.New(),
		Email:          createUser.Email,
		FirstName:      createUser.FirstName,
		LastName:       createUser.LastName,
		AvatarUrl:      createUser.AvatarUrl,
		HashedPassword: hashedPassword,
	}, err
}
