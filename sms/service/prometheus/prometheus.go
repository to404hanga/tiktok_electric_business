package prometheus

import (
	"context"
	"tiktok_electric_business/sms/service"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

// PrometheusSmsService 接入prometheus的短信服务
type PrometheusSmsService struct {
	svc    service.SmsService
	vector *prometheus.SummaryVec
}

var _ service.SmsService = (*PrometheusSmsService)(nil)

func NewPrometheusSmsService(svc service.SmsService, opt prometheus.SummaryOpts) service.SmsService {
	return &PrometheusSmsService{
		svc:    svc,
		vector: prometheus.NewSummaryVec(opt, []string{"tpl_id"}),
	}
}

func (s *PrometheusSmsService) Send(ctx context.Context, tplId string, args []string, numbers ...string) error {
	start := time.Now()
	defer func() {
		duration := time.Since(start).Milliseconds()
		s.vector.WithLabelValues(tplId).Observe(float64(duration))
	}()
	return s.svc.Send(ctx, tplId, args, numbers...)
}
