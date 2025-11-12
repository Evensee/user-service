package transaction

import (
	"context"

	"gorm.io/gorm"
)

type AppTransaction interface {
	GetOrmTx() *gorm.DB
	CloseTransaction()
}

type CreateAppTransactionResolver interface {
	CreateAppTransaction(context.Context) AppTransaction
}

type CreateAppTransaction func(context.Context) AppTransaction
