package config

type PasetoConfig struct {
	PubKey   string `mapstructure:"pub_key" json:"pub_key"`
	Implicit string `mapstructure:"implicit" json:"implicit"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Key  string `mapstructure:"key" json:"key"`
}

type OtelConfig struct {
	EndPoint string `mapstructure:"endpoint" json:"endpoint"`
}

type RPCSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}

type ServerConfig struct {
	Name        string       `mapstructure:"name" json:"name"`
	Host        string       `mapstructure:"host" json:"host"`
	Port        int          `mapstructure:"port" json:"port"`
	GPTKey      string       `mapstructure:"gpt" json:"gpt"`
	ProxyURL    string       `mapstructure:"proxy" json:"proxy"`
	PasetoInfo  PasetoConfig `mapstructure:"paseto" json:"paseto"`
	OtelInfo    OtelConfig   `mapstructure:"otel" json:"otel"`
	UserSrvInfo RPCSrvConfig `mapstructure:"user_srv" json:"user_srv"`
}
