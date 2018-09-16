package datastoreservice

import (
	"AP/common"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type CacheService interface {
	Get(key string) (*common.WeatherResponse, error)
	Set(key string, value *common.WeatherResponse, expiration time.Duration) error
}

func NewCacheService() CacheService {
	return &RedisCache{
		Client: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		}),
	}
}

type RedisCache struct {
	Client *redis.Client
}

func (rc *RedisCache) isCacheActive() (bool, error) {
	_, err := rc.Client.Ping().Result()
	if err != nil {
		return false, err
	}

	return true, nil
}

//Get returns a value mapped to key
func (rc *RedisCache) Get(key string) (*common.WeatherResponse, error) {
	if len(key) == 0 {
		return nil, fmt.Errorf("Get: Key's empty")
	}

	isCacheAlive, aliveErr := rc.isCacheActive()
	if aliveErr != nil {
		return nil, aliveErr
	}

	if !isCacheAlive {
		return nil, fmt.Errorf("Get:cache server not alive")
	}

	var value common.WeatherResponse
	valueBytes, valueErr := rc.Client.Get(key).Bytes()
	if valueErr == redis.Nil {
		return nil, fmt.Errorf("Get:key (%s) does not exist; Err:(%v)", key, valueErr)
	}

	if valueErr != nil {
		return nil, valueErr
	}

	err := json.Unmarshal(valueBytes, &value)
	return &value, err
}

//Set saves a key,value into redis
func (rc *RedisCache) Set(key string, value *common.WeatherResponse, expiration time.Duration) error {
	if key == "" {
		return fmt.Errorf("Set:Invalid key")
	}

	if value == nil {
		return fmt.Errorf("Set:Invalid value(%v) for key (%s)", value, key)
	}

	if rc.Client == nil {
		return fmt.Errorf("Set:Invalid client reference")
	}

	valueMarshaled, valMarshaledErr := json.Marshal(*value)
	if valMarshaledErr != nil {
		return valMarshaledErr
	}

	setValueErr := rc.Client.Set(key, valueMarshaled, expiration).Err()
	if setValueErr != nil {
		return setValueErr
	}

	fmt.Println("Saved value:", string(valueMarshaled))
	return nil
}
