package redis

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/pstpmn/my-golang-hexagonal-template/internal/core/port"
	rd "github.com/redis/go-redis/v9"
)

type r struct {
	client *rd.Client
}

// Delete implements port.ICache.
func (r *r) Delete(ctx context.Context, key string) error {
	err := r.client.Del(ctx, key).Err()
	return err
}

func (r *r) CacheIgnoreDuplcateKey(ctx context.Context, key string, data string, expireTime time.Time) error {
	expirationDuration := time.Until(expireTime) // Calculate the duration until expiration

	// Use SET instead of SETNX to force updating the cache even if the key already exists
	err := r.client.Set(ctx, key, data, expirationDuration).Err()
	if err != nil {
		fmt.Printf("error in function save cacheIgnoreDuplcateKey adapter %s\n", err)
		return errors.New("error can't save cache data")
	}

	return nil // Cache was successfully updated
}

// Cache implements port.ICache.
func (r *r) Cache(ctx context.Context, key string, data string, expireTime time.Time) (bool, error) {
	expirationDuration := time.Until(expireTime) // Calculate the duration until expiration
	isSucc, err := r.client.SetNX(ctx, key, data, expirationDuration).Result()
	if err != nil {
		fmt.Printf("error in function save cacheIgnoreDuplcateKey adapter %s\n", err)
		return false, errors.New("error can't save cache data")
	}

	if !isSucc {
		return true, nil
	}

	return false, nil
}

// Get implements port.ICache.
func (r *r) Get(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()

	if err == rd.Nil {
		return "", errors.New(`key dose not exist`)
	} else if err != nil {
		fmt.Printf("error in function get cache adapter %s\n", err)
		return "", errors.New(`error can't get data in cache`)
	}
	return val, nil
}

// New creates a new instance of Redis
func NewRedis(uri string) (port.ICache, error) {
	opts, err := rd.ParseURL(uri)
	if err != nil {
		return &r{}, err
	}
	return &r{
		client: rd.NewClient(opts),
	}, nil
}
