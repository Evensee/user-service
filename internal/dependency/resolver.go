package dependency

import (
	"context"

	"github.com/Evensee/user-service/internal"
	"github.com/Evensee/user-service/internal/interface/resolver"
	"github.com/Evensee/user-service/internal/interface/service"
	"github.com/Evensee/user-service/internal/interface/transaction"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Resolver struct {
	db     *gorm.DB
	rdb    *redis.Client
	config internal.AppConfig
}

func NewResolver(db *gorm.DB, rdb *redis.Client, config internal.AppConfig) resolver.AppResolver {
	return Resolver{
		db:     db,
		rdb:    rdb,
		config: config,
	}
}

func (r Resolver) CreateAppService(appTransaction transaction.AppTransaction) service.AppService {
	return CreateAppService(appTransaction, &r.config)
}

func (r Resolver) CreateAppTransaction(ctx context.Context) transaction.AppTransaction {
	return StartTransaction(r.db, ctx, r.rdb)
}
