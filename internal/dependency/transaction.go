package dependency

import (
	"context"

	"github.com/Evensee/user-service/internal/infrastructure/memory"
	"github.com/Evensee/user-service/internal/interface/transaction"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type (
	Ctx = context.Context
)

func StartTransaction(
	db *gorm.DB,
	ctx context.Context,
	rdb *redis.Client,
) transaction.AppTransaction {
	ormTx := db.WithContext(ctx).Begin()
	rTx := memory.CreateTransaction(rdb)

	return Transaction{
		ormTx: ormTx,
		rTx:   rTx,
	}
}

type Transaction struct {
	ormTx *gorm.DB
	rTx   redis.Pipeliner
}

func (t Transaction) GetOrmTx() *gorm.DB {
	return t.ormTx
}

func (t Transaction) GetMemoryTx() redis.Pipeliner {
	return t.rTx
}

func (t Transaction) rollback(ctx context.Context) {
	t.ormTx.Rollback()
	memory.RollbackTransaction(ctx, t.rTx)
}

func (t Transaction) commitOrm(ctx Ctx) {
	ormResult := t.ormTx.Commit()

	if recoverError := recover(); recoverError != nil {
		t.rollback(ctx)

		panic(recoverError)
	}

	if ormResult.Error != nil {
		t.rollback(ctx)

		panic(ormResult.Error)
	}
}

func (t Transaction) commitMemory(ctx Ctx) {
	memory.CommitTransaction(ctx, t.rTx)
}

func (t Transaction) CloseTransaction(ctx Ctx) {
	t.commitOrm(ctx)
	t.commitMemory(ctx)
}
