package resolver

import (
	"github.com/Evensee/user-service/internal/interface/service"
	"github.com/Evensee/user-service/internal/interface/transaction"
)

type AppResolver interface {
	service.CreateAppServiceResolver
	transaction.CreateAppTransactionResolver
}
