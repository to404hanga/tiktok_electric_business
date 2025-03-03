package local

import (
	"context"
	"tiktok_electric_business/sms/service"

	"github.com/to404hanga/pkg404/logger"
)

// LocalSmsService 本地短信服务，使用debug级别打印到日志文件
type LocalSmsService struct {
	l logger.Logger
}

var _ service.SmsService = (*LocalSmsService)(nil)

func NewLocalSmsService(l logger.Logger) service.SmsService {
	return &LocalSmsService{
		l: l,
	}
}

func (s *LocalSmsService) Send(ctx context.Context, tplId string, args []string, numbers ...string) error {
	for i := 0; i < len(args); i++ {
		s.l.Debug("Verification code", logger.String("code", args[i]), logger.SafePhoneZH(numbers[i]))
	}
	return nil
}
