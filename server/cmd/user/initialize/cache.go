package initialize

import (
	"context"
	"net"
	"project/server/cmd/user/config"
	"strconv"

	"github.com/redis/go-redis/v9"

	"github.com/redis/go-redis/extra/redisotel/v9"

	"github.com/cloudwego/kitex/pkg/klog"
)

func InitCache() *redis.Client {
	c := config.GlobalServerConfig.RedisInfo
	rc := redis.NewClient(&redis.Options{
		Addr:     net.JoinHostPort(c.Host, strconv.Itoa(c.Port)),
		Password: c.Password,
		DB:       c.DB,
	})
	pong, err := rc.Ping(context.Background()).Result()
	if err != nil {
		klog.Fatalf("user redis connect ping failed, err:", err)
	}
	klog.Info("user redis init ping resp:", pong)
	err = redisotel.InstrumentTracing(rc)
	if err != nil {
		klog.Fatalf("user redisotel.InstrumentTracing(rc) failed, err:", err)
	}

	return rc
}
