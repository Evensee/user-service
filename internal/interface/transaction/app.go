package transaction

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Ctx = context.Context

type AppTransaction interface {
	GetOrmTx() *gorm.DB
	GetMemoryTx() redis.Pipeliner
	CloseTransaction(ctx Ctx)
}

type CreateAppTransactionResolver interface {
	CreateAppTransaction(Ctx) AppTransaction
}

type CreateAppTransaction func(Ctx) AppTransaction
