package redis

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/hadan/gogox/errorx"
)

// Redis implements Cache interface for redis memory storage
type Redis struct {
	client *redis.Client
}

// New creates new Redis implementation
func New(client *redis.Client) *Redis {
	return &Redis{client: client}
}

// Get fetch key value and then unmarshal it into dest. Will return gogox error with code CodeNotFound if key is not found.
func (r *Redis) Get(ctx context.Context, key string, dest interface{}) error {
	bytes, err := r.client.Get(ctx, key).Bytes()
	if errors.Is(err, redis.Nil) {
		return errorx.Wrap(err, errorx.CodeNotFound, "redis key not found")
	}

	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, dest)
}

// Set set value for key using value json marshal result. Expiration is set accordingly.
func (r *Redis) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	setBytes, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return r.client.Set(ctx, key, setBytes, expiration).Err()
}

// Del delete keys
func (r *Redis) Del(ctx context.Context, keys ...string) error {
	return r.client.Del(ctx, keys...).Err()
}
