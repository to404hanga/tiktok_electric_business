package cache

import "context"

//go:generate mockgen -source=./types.go -package=cachemocks -destination=./mocks/code.mock.go CodeCache
type CodeCache interface {
	Set(ctx context.Context, biz, phone, code string) error
	Verify(ctx context.Context, biz, phone, code string) (bool, error)
}
