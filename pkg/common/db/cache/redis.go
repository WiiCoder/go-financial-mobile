package cache

import (
	"github.com/redis/go-redis/v9"
	"github.com/wiiCoder/installment/pkg/common/config"
)

func NewRedis() *redis.Client {
	conf := config.Config.Redis
	return redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Password: conf.Password,
		DB:       conf.Db,
	})
}
