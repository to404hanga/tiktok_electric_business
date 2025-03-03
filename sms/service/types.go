package service

import "context"

//go:generate mockgen -source=./types.go -package=svcmocks -destination=./mocks/sms.mock.go SmsService
type SmsService interface {
	Send(ctx context.Context, tplId string, args []string, numbers ...string) error
}
