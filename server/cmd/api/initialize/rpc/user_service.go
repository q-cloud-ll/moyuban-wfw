package rpc

import (
	"fmt"
	"project/server/cmd/api/config"
	"project/server/shared/kitex_gen/user/userservice"

	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	consul "github.com/kitex-contrib/registry-consul"
	"gorm.io/plugin/opentelemetry/provider"
)

func initUser() {
	r, err := consul.NewConsulResolver(fmt.Sprintf("%s:%d",
		config.GlobalConsulConfig.Host, 8500))
	if err != nil {
		klog.Fatalf("new consul client failed: %s", err.Error())

	}
	// init OpenTelemetry
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(config.GlobalServerConfig.UserSrvInfo.Name),
		provider.WithExportEndpoint(config.GlobalServerConfig.OtelInfo.EndPoint),
		provider.WithInsecure(),
	)

	c, err := userservice.NewClient(config.GlobalServerConfig.UserSrvInfo.Name,
		client.WithResolver(r),
		client.WithLoadBalancer(loadbalance.NewWeightedBalancer()),
		client.WithMuxConnection(1),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.GlobalServerConfig.UserSrvInfo.Name}))

	if err != nil {
		klog.Fatalf("ERROR: cannot init client: %v\n", err)
	}
	config.GlobalUserClient = c
}
