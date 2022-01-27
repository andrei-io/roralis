package kv

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
)

type IKVStore interface {
	Set(key string, value string, expire_minutes int64) error
	Update(key string, value string) error
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
	_, err := k.client.SetNX(ctx, key, value, time.Duration(expire_minutes)*60*time.Second).Result()
	return err

}
func (k *KVStore) Update(key string, value string) error {
	return errors.New("Not implemented yet")
}
func (k *KVStore) Delete(key string) error {
	return errors.New("Not implemented yet")
}
