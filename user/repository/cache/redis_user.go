package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"tiktok_electric_business/user/domain"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisUserCache struct {
	cmd        redis.Cmdable
	expiration time.Duration
}

var _ UserCache = (*RedisUserCache)(nil)

func NewRedisUserCache(cmd redis.Cmdable, expiration time.Duration) UserCache {
	return &RedisUserCache{
		cmd:        cmd,
		expiration: expiration,
	}
}

// Get 根据 id 获取用户信息缓存
func (u *RedisUserCache) Get(ctx context.Context, id int64) (domain.User, error) {
	var user domain.User
	data, err := u.cmd.Get(ctx, u.key(id)).Result()
	if err != nil {
		return user, err
	}
	err = json.Unmarshal([]byte(data), &user)
	return user, err
}

// Set 设置用户信息缓存，并使用 RedisUserCache 提供的过期时间
func (u *RedisUserCache) Set(ctx context.Context, user domain.User) error {
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return u.cmd.Set(ctx, u.key(user.Id), data, u.expiration).Err()
}

// Del 根据 id 删除用户信息缓存
func (u *RedisUserCache) Del(ctx context.Context, id int64) error {
	return u.cmd.Del(ctx, u.key(id)).Err()
}

// key 根据 id 生成访问 redis 的键
func (u *RedisUserCache) key(id int64) string {
	return fmt.Sprintf("teb:user:info:%d", id)
}
