package storages

import (
	"fmt"

	"github.com/morehao/golib/dbstore/dbredis"
	"github.com/morehao/golib/glog"
	"github.com/redis/go-redis/v9"
)

var (
	DemoRedis *redis.Client
	IAMRedis  *redis.Client
)

const (
	RedisServiceNameDemo = "demoapp"
	RedisServiceNameIAM  = "iam"
)

func InitMultiRedis(configs []dbredis.RedisConfig, logConfig *glog.LogConfig) error {
	if len(configs) == 0 {
		return fmt.Errorf("redis config is empty")
	}
	var opts []dbredis.Option
	if logConfig != nil {
		opts = append(opts, dbredis.WithLogConfig(logConfig))
	}
	for _, cfg := range configs {
		client, err := dbredis.InitRedis(&cfg, opts...)
		if err != nil {
			return err
		}
		switch cfg.Service {
		case RedisServiceNameDemo:
			DemoRedis = client
		case RedisServiceNameIAM:
			IAMRedis = client
		default:
			return fmt.Errorf("unknown redis service: %s", cfg.Service)
		}
	}
	return nil
}
