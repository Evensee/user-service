package service

import (
	"github.com/Evensee/user-service/internal/domain/user"
	"github.com/Evensee/user-service/internal/interface/transaction"
)

type AppService interface {
	GetUserService() user.DomainUserService
}

type CreateAppServiceResolver interface {
	CreateAppService(transaction.AppTransaction) AppService
}

type CreateAppService func(transaction.AppTransaction) AppService
