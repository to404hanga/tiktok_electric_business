package ratelimiter

import (
	"context"
	"tiktok_electric_business/sms/service"

	"github.com/to404hanga/pkg404/limiter"
)

// RateLimiterSmsService 基于限流器的短信服务
type RateLimiterSmsService struct {
	svc   service.SmsService
	limit limiter.Limiter
	key   string
}

var _ service.SmsService = (*RateLimiterSmsService)(nil)

func NewRateLimiterSmsService(svc service.SmsService, limit limiter.Limiter) service.SmsService {
	return &RateLimiterSmsService{
		svc:   svc,
		limit: limit,
		key:   "teb:sms:ratelimiter",
	}
}

func (s *RateLimiterSmsService) Send(ctx context.Context, tplId string, args []string, numbers ...string) error {
	limited, err := s.limit.Limit(ctx, s.key)
	if err != nil {
		return err
	}
	if limited {
		return ErrLimited
	}
	return s.svc.Send(ctx, tplId, args, numbers...)
}
