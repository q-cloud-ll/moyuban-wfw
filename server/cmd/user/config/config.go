package config

type MysqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Name     string `mapstructure:"db" json:"db"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
	Salt     string `mapstructure:"salt" json:"salt"`
}

type ConsulConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	BasePort int    `mapstructure:"base-port" json:"base-port"`
	Key      string `mapstructure:"key" json:"key"`
}

type OtelConfig struct {
	EndPoint string `mapstructure:"endpoint" json:"endpoint"`
}

type PasetoConfig struct {
	PubKey   string `mapstructure:"secret_key" json:"pub_key"`
	Implicit string `mapstructure:"implicit" json:"implicit"`
}

type RedisConfig struct {
	Host         string `mapstructure:"host" json:"host"`
	Port         int    `mapstructure:"port" json:"port"`
	Password     string `mapstructure:"password"`
	DB           int    `mapstructure:"db"`
	PoolSize     int    `mapstructure:"pool_size"`
	MinIdleConns int    `mapstructure:"min_idle_conns"`
}

type ServerConfig struct {
	Name       string       `mapstructure:"name" json:"name"`
	Host       string       `mapstructure:"host" json:"host"`
	RedisInfo  RedisConfig  `mapstructure:"redis" json:"redis"`
	PasetoInfo PasetoConfig `mapstructure:"paseto" json:"paseto"`
	MysqlInfo  MysqlConfig  `mapstructure:"mysql" json:"mysql"`
	OtelInfo   OtelConfig   `mapstructure:"otel" json:"otel"`
}
