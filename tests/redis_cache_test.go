package tests

import (
	"context"
	"encoding/json"
	"github.com/Tomas-vilte/GCPSteamAnalytics/cache"
	"github.com/Tomas-vilte/GCPSteamAnalytics/steamapi/persistence/entity"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func setupTestRedisClient() *redis.Client {
	options := &redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	}
	client := redis.NewClient(options)
	return client
}

func TestRedisCache_Get(t *testing.T) {
	client := setupTestRedisClient()
	defer client.Close()

	cacheClient := cache.NewRedisCacheClient("localhost:6379", 0)

	// Datos de prueba
	key := "test_key"
	gameDetails := &entity.GameDetails{
		Name: "Test Game",
	}

	value, err := json.Marshal(gameDetails)
	require.NoError(t, err)
	err = client.Set(context.Background(), key, value, 0).Err()
	require.NoError(t, err)

	retrievedGameDetails, err := cacheClient.Get(key)
	require.NoError(t, err)
	require.NotNil(t, retrievedGameDetails)

	assert.Equal(t, gameDetails.Name, retrievedGameDetails.Name)
}

func TestRedisCache_Set(t *testing.T) {
	client := setupTestRedisClient()
	defer client.Close()

	cacheClient := cache.NewRedisCacheClient("localhost:6379", 0)

	// Datos de prueba
	key := "test_key"
	gameDetails := &entity.GameDetails{
		Name: "Test Game",
	}

	value, err := json.Marshal(gameDetails)
	require.NoError(t, err)
	err = cacheClient.Set(key, string(value))
	require.NoError(t, err)

	// Verifica si el valor se almacenó correctamente en Redis
	retrievedValue, err := client.Get(context.Background(), key).Result()
	require.NoError(t, err)
	require.NotNil(t, retrievedValue)

	assert.Equal(t, string(value), retrievedValue)
}
