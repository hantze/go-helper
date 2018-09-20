package redis

import (
	"strconv"

	"time"

	goredis "github.com/go-redis/redis"
)

// Manager ...
type Manager struct {
	client *goredis.Client
}

// NewManager ...
func NewManager(addr string, password string) (*Manager, error) {
	client := goredis.NewClient(&goredis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return &Manager{client: client}, nil
}

// GetClient ...
func (manager Manager) Client() *goredis.Client {
	return manager.client
}

// UpdateValue ...
func (manager Manager) UpdateValue(key, value string) error {
	err := manager.client.Set(key, value, 0).Err()
	return err
}

// StoreValue ...
func (manager Manager) StoreValue(key, value string) error {
	err := manager.client.Set(key, value, 0).Err()
	return err
}

// StoreIntValue ...
func (manager Manager) StoreIntValue(key string, value int) error {
	result := strconv.Itoa(value)
	return manager.StoreValue(key, result)
}

// GetValue ...
func (manager Manager) Value(key string) (string, error) {
	val, err := manager.client.Get(key).Result()
	return val, err
}

// Remove ...
func (manager Manager) Remove(key string) (int64, error) {
	result, err := manager.client.Del(key).Result()
	return result, err
}

// Expire ...
func (manager Manager) Expire(key string, duration time.Duration) bool {
	result, _ := manager.client.Expire(key, duration).Result()
	return result
}

// GetIntValue ...
func (manager Manager) IntValue(key string) (int, error) {
	val, err := manager.Value(key)
	if err != nil {
		return 0, err
	}
	intVal, err := strconv.Atoi(val)
	if err != nil {
		return 0, err
	}
	return intVal, nil
}

// FlushAll ...
func (manager Manager) FlushAll() (string, error) {
	return manager.client.FlushAll().Result()
}
