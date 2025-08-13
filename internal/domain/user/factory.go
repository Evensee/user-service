package user

import "github.com/google/uuid"

func NewUser(email, hashedPassword string) (*User, error) {
	return &User{
		ID:             uuid.NewString(),
		Email:          email,
		HashedPassword: hashedPassword,
	}, nil
}
