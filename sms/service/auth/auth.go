package auth

import (
	"context"
	"tiktok_electric_business/sms/service"

	"github.com/golang-jwt/jwt/v5"
)

// AuthSmsService 使用jwt认证的短信服务
type AuthSmsService struct {
	svc service.SmsService
	key []byte
}

var _ service.SmsService = (*AuthSmsService)(nil)

func NewAuthSmsService(svc service.SmsService, key []byte) service.SmsService {
	return &AuthSmsService{
		svc: svc,
		key: key,
	}
}

// Send 发送短信
func (s *AuthSmsService) Send(ctx context.Context, tplToken string, args []string, numbers ...string) error {
	claims := SmsClaims{}
	_, err := jwt.ParseWithClaims(tplToken, &claims, func(t *jwt.Token) (interface{}, error) {
		return s.key, nil
	})
	if err != nil {
		return err
	}
	return s.svc.Send(ctx, claims.Tpl, args, numbers...)
}

type SmsClaims struct {
	jwt.RegisteredClaims
	Tpl string
}
