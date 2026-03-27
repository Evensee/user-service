package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/Evensee/user-service/internal"
	"github.com/Evensee/user-service/internal/domain/auth"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type authTokenRepository struct {
	pipe redis.Pipeliner
	config internal.AppConfig
}

func NewAuthTokenRepository(pipe redis.Pipeliner) auth.Repository {
	return &authTokenRepository{
		pipe: pipe,
	}
}

func (r *authTokenRepository) Save(ctx context.Context, userID uuid.UUID, tokens auth.Tokens) error {
	r.pipe.Set(
		ctx,
		r.accessTokenKey(tokens.AccessToken),
		userID.String(),
		time.Second*time.Duration(r.config.AccessTokenLifetimeSeconds),
	)
	r.pipe.Set(
		ctx,
		r.refreshTokenKey(tokens.RefreshToken),
		userID.String(),
		time.Second*time.Duration(r.config.RefreshTokenLifetimeSeconds),
	)

	return nil
}

func (r *authTokenRepository) GetUserIDByAccessToken(ctx context.Context, accessToken string) (uuid.UUID, error) {
	val, err := r.pipe.Get(ctx, r.accessTokenKey(accessToken)).Result()
	if err != nil {
		if err == redis.Nil {
			return uuid.Nil, fmt.Errorf("access token not found")
		}
		return uuid.Nil, err
	}
	return uuid.Parse(val)
}

func (r *authTokenRepository) GetUserIDByRefreshToken(ctx context.Context, refreshToken string) (uuid.UUID, error) {
	val, err := r.pipe.Get(ctx, r.refreshTokenKey(refreshToken)).Result()
	if err != nil {
		if err == redis.Nil {
			return uuid.Nil, fmt.Errorf("refresh token not found")
		}
		return uuid.Nil, err
	}
	return uuid.Parse(val)
}

func (r *authTokenRepository) BlockAccessToken(ctx context.Context, accessToken string, userId uuid.UUID) {
	r.pipe.Set(
		ctx,
		r.accessTokenBlockKey(accessToken),
		userId.String(),
		time.Second*time.Duration(r.config.AccessTokenLifetimeSeconds),
	)
}

func (r *authTokenRepository) BlockRefreshToken(ctx context.Context, refreshToken string, userId uuid.UUID) {
	r.pipe.Set(
		ctx,
		r.refreshTokenBlockKey(refreshToken),
		userId.String(),
		time.Second*time.Duration(r.config.RefreshTokenLifetimeSeconds),
	)
}

func (r *authTokenRepository) CheckAccessTokenBlocked(ctx context.Context, accessToken string) (*uuid.UUID, error) {
	val, err := r.pipe.Get(ctx, r.accessTokenBlockKey(accessToken)).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, fmt.Errorf("access token not found")
		}
		return nil, err
	}
	userId, err := uuid.Parse(val)
	return &userId, err
}

func (r *authTokenRepository) CheckRefreshTokenBlocked(ctx context.Context, refreshToken string) (*uuid.UUID, error) {
	val, err := r.pipe.Get(ctx, r.accessTokenBlockKey(refreshToken)).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, fmt.Errorf("refresh token not found")
		}
		return nil, err
	}
	userId, err := uuid.Parse(val)
	
	return &userId, err
}

func (r *authTokenRepository) DeleteAccessToken(ctx context.Context, accessToken string) error {
	return r.pipe.Del(ctx, r.accessTokenKey(accessToken)).Err()
}

func (r *authTokenRepository) DeleteRefreshToken(ctx context.Context, refreshToken string) error {
	return r.pipe.Del(ctx, r.refreshTokenKey(refreshToken)).Err()
}

func (r *authTokenRepository) accessTokenKey(token string) string {
	return fmt.Sprintf("auth:access-token:%s", token)
}

func (r *authTokenRepository) refreshTokenKey(token string) string {
	return fmt.Sprintf("auth:refresh-token:%s", token)
}

func (r *authTokenRepository) accessTokenBlockKey(token string) string {
	return fmt.Sprintf("auth:blocklist:access-token:%s", token)
}

func (r *authTokenRepository) refreshTokenBlockKey(token string) string {
	return fmt.Sprintf("auth:blocklist:refresh-token:%s", token)
}
