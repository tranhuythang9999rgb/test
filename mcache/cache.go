package mcache

import (
	"ap_sell_products/common/configs"
	"sync"

	"github.com/redis/go-redis/v9"
)

var (
	rdb  *redis.Client
	cfg  *configs.Configs
	once sync.Once
	lock sync.Mutex
)

func Init(config *configs.Configs) {
	lock.Lock()
	defer lock.Unlock()

	if cfg == nil {
		cfg = config
	}

	if rdb == nil {
		once.Do(func() {
			rdb = redis.NewClient(&redis.Options{
				Addr:     cfg.AddressRedis,
				Password: cfg.PasswordRedis,
				DB:       cfg.DatabaseredisIndex,
			})
		})
	}
}

func GetRDB() *redis.Client {
	lock.Lock()
	defer lock.Unlock()
	return rdb
}
