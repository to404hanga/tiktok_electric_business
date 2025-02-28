package service

import (
	"context"
	"tiktok_electric_business/user/domain"
)

//go:generate mockgen -source=./types.go -package=svcmocks -destination=./mocks/user.mock.go UserService
type UserService interface {
	FindOrCreateByPhone(ctx context.Context, phone string) (domain.User, error)
	FindOrCreateByWechat(ctx context.Context, info domain.WechatInfo) (domain.User, error)
	SignUp(ctx context.Context, user domain.User) error
	Login(ctx context.Context, email, password string) (domain.User, error)
	UpdateNonSensitiveInfo(ctx context.Context, user domain.User) error
	Profile(ctx context.Context, id int64) (domain.User, error)
}
