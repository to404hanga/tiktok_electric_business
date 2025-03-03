package cache

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/to404hanga/pkg404/logger"
)

var (
	//go:embed lua/set_code.lua
	luaSetCode string
	//go:embed lua/verify_code.lua
	luaVerifyCode string
)

type RedisCodeCache struct {
	cmd redis.Cmdable
	l   logger.Logger
}

var _ CodeCache = (*RedisCodeCache)(nil)

func NewRedisCodeCache(cmd redis.Cmdable, l logger.Logger) CodeCache {
	return &RedisCodeCache{
		cmd: cmd,
		l:   l,
	}
}

func (c *RedisCodeCache) Set(ctx context.Context, biz, phone, code string) error {
	res, err := c.cmd.Eval(ctx, luaSetCode, []string{key(biz, phone)}, code).Int()
	if err != nil {
		return err
	}
	switch res {
	case 0:
		return nil
	case -1:
		return ErrCodeSendTooMany
	default:
		c.l.Error("An unknown error occurred while sending the captcha", logger.Int("error_code", res))
		return ErrUnknowForCode
	}
}

func (c *RedisCodeCache) Verify(ctx context.Context, biz, phone, code string) (bool, error) {
	res, err := c.cmd.Eval(ctx, luaVerifyCode, []string{key(biz, phone)}, code).Int()
	if err != nil {
		return false, err
	}
	switch res {
	case 0:
		return true, nil
	case -1:
		return false, ErrCodeVerifyTooMany
	default:
		return false, nil
	}
}

func key(biz, phone string) string {
	return fmt.Sprintf("phone_code:%s:%s", biz, phone)
}
