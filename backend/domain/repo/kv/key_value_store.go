package kv

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
)

type IKVStore interface {
	Set(key string, value string, expire_minutes int64) error
	Get(key string) (string, error)
	Delete(key string) error
}

type KVStore struct {
	client redis.Client
}

func NewKVStore(address string, password string) *KVStore {
	kv := KVStore{
		client: *redis.NewClient(&redis.Options{
			Addr:     address,
			Password: password,
		}),
	}
	return &kv
}

func (k *KVStore) Set(key string, value string, expire_minutes int64) error {
	ctx := context.TODO()
	_, err := k.client.Set(ctx, key, value, time.Duration(expire_minutes)*60*time.Second).Result()
	return err

}
func (k *KVStore) Get(key string) (string, error) {
	ctx := context.TODO()
	v, err := k.client.Get(ctx, key).Result()
	return v, err
}

func (k *KVStore) Delete(key string) error {
	return errors.New("Not implemented yet")
}
