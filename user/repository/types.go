package repository

import (
	"context"
	"tiktok_electric_business/user/domain"
)

//go:generate mockgen -source=./types.go -package=repomocks -destination=./mocks/user.mock.go UserRepository
type UserRepository interface {
	Create(ctx context.Context, user domain.User) error
	FindByPhone(ctx context.Context, phone string) (domain.User, error)
	FindByEmail(ctx context.Context, email string) (domain.User, error)
	UpdateNonZeroFields(ctx context.Context, user domain.User) error
	FindById(ctx context.Context, id int64) (domain.User, error)
	FindByWechat(ctx context.Context, openId string) (domain.User, error)
	DeleteById(ctx context.Context, id int64) error
}
