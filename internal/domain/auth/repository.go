package auth

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Save(context.Context, uuid.UUID, Tokens) error
	GetUserIDByAccessToken(context.Context, string) (uuid.UUID, error)
	GetUserIDByRefreshToken(context.Context, string) (uuid.UUID, error)
	DeleteAccessToken(context.Context, string) error
	DeleteRefreshToken(context.Context, string) error
}
