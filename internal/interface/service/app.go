package service

import (
	"github.com/Evensee/user-service/internal"
	"github.com/Evensee/user-service/internal/domain/auth"
	"github.com/Evensee/user-service/internal/domain/user"
	"github.com/Evensee/user-service/internal/interface/transaction"
)

type AppService interface {
	GetUserService() user.DomainUserService
	GetAuthService() auth.AuthService
}

type CreateAppServiceResolver interface {
	CreateAppService(transaction.AppTransaction, *internal.AppConfig) AppService
}

type CreateAppService func(transaction.AppTransaction, *internal.AppConfig) AppService
