package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type CacheClient interface {
	Get(key string) (string, error)
	Set(key string, value string) error
}

type redisCache struct {
	host       string
	db         int
	expiration time.Duration
}

func NewRedisCacheClient(host string, db int, exp time.Duration) CacheClient {
	return &redisCache{
		host:       host,
		db:         db,
		expiration: exp,
	}
}

func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}

func (cache *redisCache) Get(key string) (string, error) {
	ctx := context.Background()
	value, err := cache.getClient().Get(ctx, key).Result()
	if err == redis.Nil {
		// La clave no existe en la caché
		return "", err
	} else if err != nil {
		// Se produjo un error al obtener el valor
		return "", err
	}
	return value, nil
}

func (cache *redisCache) Set(key string, value string) error {
	ctx := context.Background()
	err := cache.getClient().Set(ctx, key, value, cache.expiration).Err()
	if err != nil {
		return err
	}
	return nil
}
