package config

import "project/server/shared/kitex_gen/user/userservice"

var (
	GlobalServerConfig ServerConfig
	GlobalConsulConfig ConsulConfig

	GlobalUserClient userservice.Client
)
