package auth

import (
	"github.com/google/uuid"
)

type Repository interface {
	Save(Ctx, uuid.UUID, Tokens) error
	GetUserIDByAccessToken(Ctx, string) (uuid.UUID, error)
	GetUserIDByRefreshToken(Ctx, string) (uuid.UUID, error)
	DeleteAccessToken(Ctx, string) error
	DeleteRefreshToken(Ctx, string) error
	BlockAccessToken(Ctx, string, uuid.UUID)
	BlockRefreshToken(Ctx, string, uuid.UUID)
	CheckAccessTokenBlocked(Ctx, string) (*uuid.UUID, error)
	CheckRefreshTokenBlocked(Ctx, string) (*uuid.UUID, error)
}
