package ioc

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"github.com/to404hanga/pkg404/limiter"
	"github.com/to404hanga/pkg404/limiter/redisslidewindow"
)

func InitLimiter(cmd redis.Cmdable) limiter.Limiter {
	interval := viper.GetDuration("limiter.interval")
	rate := viper.GetInt("limiter.rate")
	return redisslidewindow.NewRedisSlidingWindowLimiter(cmd, interval, rate)
}
