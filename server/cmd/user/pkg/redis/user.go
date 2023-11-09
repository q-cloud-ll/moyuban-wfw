package redis

import (
	"context"
	"project/server/shared/consts"
	"time"

	"github.com/redis/go-redis/v9"
)

type Manager struct {
	client *redis.Client
}

func NewManager(c *redis.Client) *Manager {
	return &Manager{client: c}
}

func (m *Manager) SavePhoneMsg(ctx context.Context, phone, code string) error {
	return m.client.Set(ctx, consts.GetRedisKey(phone), code, 2*time.Minute).Err()
}

func (m *Manager) GetPhoneMsg(ctx context.Context, phone string) (code string, err error) {
	return m.client.Get(ctx, consts.GetRedisKey(phone)).Result()
}
