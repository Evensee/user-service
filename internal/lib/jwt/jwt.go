package jwt

import (
	"time"

	"github.com/Evensee/user-service/internal"
	"github.com/Evensee/user-service/internal/domain/user"
	"github.com/golang-jwt/jwt/v5"
)

func NewToken(user user.User, appConfig internal.AppConfig, duration time.Duration) (string, error) {
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
