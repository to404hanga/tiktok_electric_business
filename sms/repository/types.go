package repository

import (
	"context"
	"tiktok_electric_business/sms/domain"
)

//go:generate mockgen -source=./types.go -package=repomocks -destination=./mocks/sms.mock.go AsyncSmsRepository
type AsyncSmsRepository interface {
	Add(ctx context.Context, sms domain.AsyncSms) error
	PreemptWaitingSMS(ctx context.Context) (domain.AsyncSms, error)
	ReportScheduleResult(ctx context.Context, id int64, success bool) error
}
