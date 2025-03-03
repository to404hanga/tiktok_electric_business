package service

import (
	"context"
	"errors"
	"fmt"
	"tiktok_electric_business/user/domain"
	"tiktok_electric_business/user/repository"

	"github.com/to404hanga/pkg404/limiter"
	"github.com/to404hanga/pkg404/logger"
)

const prefix = "teb:user:limiter"

var ErrTooManyRequests = errors.New("too many requests, please try again later")

type LimitedUserService struct {
	BaseUserService
	limit limiter.Limiter
}

var _ UserService = (*LimitedUserService)(nil)

func NewLimitedUserService(repo repository.UserRepository, l logger.Logger, limit limiter.Limiter) UserService {
	return &LimitedUserService{
		BaseUserService: BaseUserService{
			repo: repo,
			l:    l,
		},
		limit: limit,
	}
}

// Login 带限流器的用户登陆（邮箱）
func (u *LimitedUserService) Login(ctx context.Context, email, password string) (domain.User, error) {
	limited, err := u.limit.Limit(ctx, fmt.Sprintf("%s:login:%s", prefix, email))
	if err != nil {
		u.l.Error("Limiter error", logger.Error(err))
	}
	if limited {
		u.l.Warn("Too many login attempts", logger.SafeEmail(email))
		return domain.User{}, ErrTooManyRequests
	}

	return u.BaseUserService.Login(ctx, email, password)
}
