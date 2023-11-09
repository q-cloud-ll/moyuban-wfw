package consts

const (
	Prefix = "myb"
)

// GetRedisKey 给key加上前缀
func GetRedisKey(key string) string {
	return Prefix + key
}
