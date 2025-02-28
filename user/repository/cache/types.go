package cache

import (
	"context"
	"tiktok_electric_business/user/domain"
)

//go:generate mockgen -source=./types.go -package=cachemocks -destination=./mocks/user.mock.go UserCache
type UserCache interface {
	Get(ctx context.Context, id int64) (domain.User, error)
	Set(ctx context.Context, user domain.User) error
	Del(ctx context.Context, id int64) error
}
