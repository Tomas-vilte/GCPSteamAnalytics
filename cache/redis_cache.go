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
	host string
	db   int
}

func NewRedisCacheClient(host string, db int) RedisClient {
	return &redisCache{
		host: host,
		db:   db,
	}
}

func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:        cache.host,
		Password:    "",
		DB:          cache.db,
		DialTimeout: 100 * time.Millisecond,
		ReadTimeout: 100 * time.Millisecond,
	})
}

func (cache *redisCache) Get(key string) (*entity.GameDetails, error) {
	ctx := context.Background()
	value, err := cache.getClient().Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			log.Printf("La clave %s no existe en la caché.", key)
			return nil, err
		}
		log.Printf("Error al obtener el valor de la clave %s: %v", key, err)
		return nil, err
	}

	// Ahora, convierte la cadena JSON en un objeto JSON parseado.
	var gameDetails entity.GameDetails
	if err := json.Unmarshal([]byte(value), &gameDetails); err != nil {
		log.Printf("Error al analizar JSON de la clave %s: %v", key, err)
		return nil, err
	}
	log.Println(value)
	return &gameDetails, nil
}

func (cache *redisCache) Set(key string, value string) error {
	ctx := context.Background()
	err := cache.getClient().Set(ctx, key, value, 10*time.Second).Err()
	if err != nil {
		log.Printf("Error al establecer la clave %s en la caché: %v", key, err)
		return err
	}
	log.Printf("La clave %s se ha establecido en la caché con éxito.", key)
	return nil
}
