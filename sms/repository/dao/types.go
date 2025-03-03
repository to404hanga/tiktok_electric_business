package dao

import (
	"context"

	"github.com/ecodeclub/ekit/sqlx"
)

//go:generate mockgen -source=./types.go -package=daomocks -destination=./mocks/sms.mock.go AsyncSmsDAO
type AsyncSmsDAO interface {
	Insert(ctx context.Context, sms AsyncSms) error
	GetWaitingSMS(ctx context.Context) (AsyncSms, error)
	MarkSuccess(ctx context.Context, id int64) error
	MarkFailed(ctx context.Context, id int64) error
}

const (
	asyncStatusWaiting uint8 = iota
	asyncStatusFailed
	asyncStatusSuccess
)

type AsyncSms struct {
	Id        int64
	Config    sqlx.JsonColumn[SmsConfig]
	RetryCnt  int
	RetryMax  int
	Status    uint8
	CreatedAt int64
	UpdatedAt int64 `gorm:"index"`
}

type SmsConfig struct {
	TplId   string
	Args    []string
	Numbers []string
}
