package main

import (
	"context"
	"net"
	"project/server/cmd/user/config"
	"project/server/cmd/user/initialize"
	"project/server/cmd/user/pkg/mysql"
	"project/server/cmd/user/pkg/paseto"
	"project/server/cmd/user/pkg/redis"
	"project/server/shared/consts"
	"project/server/shared/kitex_gen/user/userservice"
	"strconv"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
)

func main() {
	initialize.InitLogger()
	initialize.InitConfig()
	IP, PORT := initialize.InitFlag()
	r, info := initialize.InitRegistry(PORT)
	db := initialize.InitDB()
	cache := initialize.InitCache()
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(config.GlobalServerConfig.Name),
		provider.WithExportEndpoint(config.GlobalServerConfig.OtelInfo.EndPoint),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())
	tg, err := paseto.NewTokenGenerator(config.GlobalServerConfig.PasetoInfo.PubKey, []byte((config.GlobalServerConfig.PasetoInfo.Implicit)))
	if err != nil {
		klog.Fatal(err)
	}
	svr := userservice.NewServer(&UserServiceImpl{
		UserMysqlManager: mysql.NewUserManager(db, config.GlobalServerConfig.MysqlInfo.Salt),
		TokenGenerator:   tg,
		RedisManager:     redis.NewManager(cache),
	},
		server.WithServiceAddr(utils.NewNetAddr(consts.TCP, net.JoinHostPort(IP, strconv.Itoa(PORT)))),
		server.WithRegistry(r),
		server.WithRegistryInfo(info),
		server.WithLimit(&limit.Option{MaxConnections: 2000, MaxQPS: 500}),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.GlobalServerConfig.Name}))

	err = svr.Run()

	if err != nil {
		klog.Fatal(err)
	}
}
