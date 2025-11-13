package memory

import (
	"context"

	"github.com/Evensee/user-service/internal"
	"github.com/redis/go-redis/v9"
)

type C = redis.Client

func Connect(config *internal.RedisConfig) *C {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Host,
		Password: config.Password,
		DB:       config.Db,
	})
	return rdb
}

func CreateTransaction(client *C) redis.Pipeliner {
	tx := client.TxPipeline()

	return tx
}

func RollbackTransaction(ctx context.Context, pipe redis.Pipeliner) {
	pipe.Discard()
}

func CommitTransaction(ctx context.Context, pipe redis.Pipeliner) error {
	_, err := pipe.Exec(ctx)

	return err
}
