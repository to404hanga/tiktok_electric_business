package service

import (
	"context"
	"tiktok_electric_business/user/domain"
)

type UserService interface {
	FindOrCreateByPhone(ctx context.Context, phone string) (domain.User, error)
}
