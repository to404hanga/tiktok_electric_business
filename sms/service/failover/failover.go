package failover

import (
	"context"
	"sync/atomic"
	"tiktok_electric_business/sms/service"

	"github.com/to404hanga/pkg404/logger"
)

// FailoverSmsService 带有故障转移的短信服务
type FailoverSmsService struct {
	svcs []service.SmsService
	idx  uint64
	l    logger.Logger
}

var _ service.SmsService = (*FailoverSmsService)(nil)

func NewFailoverSmsService(svcs []service.SmsService, l logger.Logger) service.SmsService {
	return &FailoverSmsService{
		svcs: svcs,
		idx:  0,
		l:    l,
	}
}

// Send 轮询所有服务商发送短信，当所有服务商都发送失败时，返回ErrAllFailed
func (s *FailoverSmsService) Send(ctx context.Context, tplId string, args []string, numbers ...string) error {
	idx := atomic.AddUint64(&s.idx, 1)
	length := uint64(len(s.svcs))
	for i := idx; i < idx+length; i++ {
		svc := s.svcs[i%length]
		err := svc.Send(ctx, tplId, args, numbers...)
		switch err {
		case nil:
			return nil
		case context.Canceled, context.DeadlineExceeded:
			return err
		}
		s.l.Error("Send SMS failure", logger.Error(err))
	}
	return ErrAllFailed
}
