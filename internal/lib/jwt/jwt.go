package jwt

import (
	"time"

	"github.com/Evensee/user-service/internal"
	"github.com/Evensee/user-service/internal/domain/user"
	"github.com/golang-jwt/jwt/v5"
)

func NewToken(
	user *user.User, 
	appConfig *internal.AppConfig, 
	duration time.Duration,
) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(duration).Unix()

	tokenString, err := token.SignedString([]byte(appConfig.Secret))

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GenerateAccessToken(u *user.User, appConfig *internal.AppConfig) (string, error) {
	accessTokenLifetime := time.Second * time.Duration(appConfig.AccessTokenLifetimeSeconds)
	return NewToken(u, appConfig, accessTokenLifetime)
}

func GenerateRefreshToken(u *user.User, appConfig *internal.AppConfig) (string, error) {
	refreshTokenLifetime := time.Second * time.Duration(appConfig.RefreshTokenLifetimeSeconds)
	return NewToken(u, appConfig, refreshTokenLifetime)
}

func GenerateOAuthTokens(u *user.User, appConfig *internal.AppConfig) (
	accessToken,
	refreshToken string,
	err error,
) {
	accessToken, err = GenerateAccessToken(u, appConfig)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = GenerateRefreshToken(u, appConfig)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
