package dependency

import (
	"context"

	"github.com/Evensee/user-service/internal/interface/transaction"
	"gorm.io/gorm"
)

func StartTransaction(db *gorm.DB, ctx context.Context) transaction.AppTransaction {
	ormTx := db.WithContext(ctx).Begin()

	return Transaction{
		ormTx: ormTx,
	}
}

type Transaction struct {
	ormTx *gorm.DB
}

func (t Transaction) GetOrmTx() *gorm.DB {
	return t.ormTx
}

func (t Transaction) rollback() {
	t.ormTx.Rollback()
}

func (t Transaction) CloseTransaction() {
	ormResult := t.ormTx.Commit()

	if recoverError := recover(); recoverError != nil {
		t.rollback()

		panic(recoverError)
	}

	if ormResult.Error != nil {
		t.rollback()

		panic(ormResult.Error)
	}
}
