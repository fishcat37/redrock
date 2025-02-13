package dao

import (
	"context"
	"github.com/redis/go-redis/v9"
	"redrock/config"
)

var Rdb *redis.Client

func InitRedis() error {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     config.Address,
		Password: config.Password,
		DB:       config.DB,
	})
	ctx := context.TODO()
	_, err := Rdb.Ping(ctx).Result()
	defer func(rdb *redis.Client) {
		if err := rdb.Close(); err != nil {
			panic(err)
		}
	}(Rdb)
	if err != nil {
		return err
	}
	return nil
}
