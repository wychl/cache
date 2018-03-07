package redis

import (
	"errors"

	"github.com/garyburd/redigo/redis"
	"github.com/wychl/cache"
)

//ErrKeyNotExist key is not exit
var ErrKeyNotExist = errors.New("key is not exist")

type redisCache struct {
	Pool    *redis.Pool
	HaskKey string
}

//New return cache
func New(pool *redis.Pool, hashKey string) cache.Cache {
	if hashKey == "" {
		hashKey = "cache"
	}
	return &redisCache{Pool: pool, HaskKey: hashKey}
}

//Set storage key value
func (r *redisCache) Set(key string, value interface{}) error {
	conn := r.Pool.Get()
	_, err := conn.Do("HSET", r.HaskKey, key, value)
	return err
}

//Get query key
func (r *redisCache) Get(key string) (interface{}, error) {
	conn := r.Pool.Get()

	state, err := redis.Int(conn.Do("HEXISTS", r.HaskKey, key))
	if err != nil {
		return nil, err
	}

	if state == 0 {
		return nil, ErrKeyNotExist
	}

	return conn.Do("HGET", r.HaskKey, key)
}

//Delete delete key
func (r *redisCache) Delete(key string) error {
	conn := r.Pool.Get()
	_, err := conn.Do("HDEL", r.HaskKey, key)

	return err
}

//IsExist judge key exist
func (r *redisCache) IsExist(key string) bool {
	conn := r.Pool.Get()

	state, err := redis.Int(conn.Do("HEXISTS", r.HaskKey, key))
	if err != nil {
		return false
	}

	return state == 1
}

//ClearAll clear all cache
func (r *redisCache) ClearAll() error {
	conn := r.Pool.Get()
	_, err := conn.Do("DEL", r.HaskKey)

	return err
}
