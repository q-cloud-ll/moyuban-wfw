package consts

import "time"

const (
	Moyuban = "Moyuban"
	Issuer
	User       = "User"
	ThirtyDays = time.Hour * 24 * 30

	HlogFilePath = "./tmp/hlog/logs/"
	KlogFilePath = "./tmp/klog/logs/"

	TCP             = "tcp"
	FreePortAddress = "localhost:0"
	CorsAddress     = "http://localhost:3000"

	ApiConfigPath  = "./server/cmd/api/config.yaml"
	UserConfigPath = "./server/cmd/user/config.yaml"

	ConsulCheckInterval                       = "7s"
	ConsulCheckTimeout                        = "5s"
	ConsulCheckDeregisterCriticalServiceAfter = "15s"

	MySqlDSN = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"

	IPFlagName  = "ip"
	IPFlagValue = "0.0.0.0"
	IPFlagUsage = "address"

	PortFlagName  = "port"
	PortFlagUsage = "port"

	UserSnowflakeNode = 2
)
