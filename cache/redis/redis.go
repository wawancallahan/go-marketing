package cache

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"matsukana.cloud/go-marketing/config"
)

var (
	ctx = context.Background()
)

type RedisCache struct {
	Client *redis.Client
}

func NewRedis(config *config.Config) (*RedisCache, error) {
	redisDbNumber := config.GetInt("REDIS_DB_NUMBER")

	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.GetString("REDIS_HOST"), config.GetInt("REDIS_PORT")),
		Password: config.GetString("REDIS_PASSWORD"),
		DB:       redisDbNumber,
	})

	rds := &RedisCache{
		Client: redisClient,
	}

	err := rds.Ping()

	if err != nil {
		return nil, err
	}

	return rds, nil
}

func (r *RedisCache) Ping() error {
	return r.Client.Ping(ctx).Err()
}

func (r *RedisCache) Set(key string, value string, expiration *time.Duration) error {
	defaultExpiration := 1 * time.Hour

	if expiration == nil {
		expiration = &defaultExpiration
	}

	return r.Client.Set(ctx, key, value, *expiration).Err()
}

func (r *RedisCache) Get(key string) (string, error) {
	value, err := r.Client.Get(ctx, key).Result()

	if err == redis.Nil {
		return "", errors.New("Redis key does not exist")
	} else if err != nil {
		return "", err
	}

	return value, nil
}

func (r *RedisCache) Delete(key string) error {
	return r.Client.Del(ctx, key).Err()
}

func (r *RedisCache) Reset() error {
	return r.Client.FlushDB(ctx).Err()
}

func (r *RedisCache) Close() error {
	return r.Client.Close()
}
