package failover

import (
	"context"
	"sync/atomic"
	"tiktok_electric_business/sms/service"
)

// TimeoutFailoverSmsService 基于连续超时故障转移的短信服务
type TimeoutFailoverSmsService struct {
	svcs      []service.SmsService
	idx       int64 // 当前使用的节点
	cnt       int64 // 连续超时数
	threshold int64 // 切换服务的阈值
}

var _ service.SmsService = (*TimeoutFailoverSmsService)(nil)

func NewTimeoutFailoverSmsService(svcs []service.SmsService, threshold int64) service.SmsService {
	return &TimeoutFailoverSmsService{
		svcs:      svcs,
		idx:       0,
		cnt:       0,
		threshold: threshold,
	}
}

func (s *TimeoutFailoverSmsService) Send(ctx context.Context, tplId string, args []string, numbers ...string) error {
	idx := atomic.LoadInt64(&s.idx)
	cnt := atomic.LoadInt64(&s.cnt)
	if cnt > s.threshold {
		newIdx := (idx + 1) % int64(len(s.svcs))
		if atomic.CompareAndSwapInt64(&s.idx, idx, newIdx) {
			atomic.StoreInt64(&s.cnt, 0)
		}
		idx = newIdx
	}
	svc := s.svcs[idx]
	err := svc.Send(ctx, tplId, args, numbers...)
	switch err {
	case nil:
		atomic.StoreInt64(&s.cnt, 0)
		return nil
	case context.Canceled, context.DeadlineExceeded:
		atomic.AddInt64(&s.cnt, 1)
	}
	return err
}
