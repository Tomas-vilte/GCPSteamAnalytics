package cache

import (
	"context"
	"encoding/json"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/persistence/entity"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

type RedisClient interface {
	Get(key string) (*entity.GameDetails, error)
	Set(key string, value string) error
}

type redisCache struct {
	host       string
	db         int
	expiration time.Duration
}

func NewRedisCacheClient(host string, db int, exp time.Duration) RedisClient {
	return &redisCache{
		host:       host,
		db:         db,
		expiration: time.Duration(exp.Seconds()),
	}
}

func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}

func (cache *redisCache) Get(key string) (*entity.GameDetails, error) {
	ctx := context.Background()
	value, err := cache.getClient().Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			// La clave no existe en la caché
			return nil, err
		}
		// Se produjo un error al obtener el valor
		return nil, err
	}

	// Ahora, convierte la cadena JSON en un objeto JSON parseado.
	var gameDetails entity.GameDetails
	if err := json.Unmarshal([]byte(value), &gameDetails); err != nil {
		return nil, err
	}

	return &gameDetails, nil
}

func (cache *redisCache) Set(key string, value string) error {
	ctx := context.Background()
	err := cache.getClient().Set(ctx, key, value, cache.expiration).Err()
	if err != nil {
		log.Printf("Error3: %v", err)
		return err
	}
	return nil
}
