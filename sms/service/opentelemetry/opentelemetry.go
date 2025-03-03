package opentelemetry

import (
	"context"
	"tiktok_electric_business/sms/service"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type OpenTelemetrySmsService struct {
	svc    service.SmsService
	tracer trace.Tracer
}

var _ service.SmsService = (*OpenTelemetrySmsService)(nil)

func NewOpenTelemetrySmsService(svc service.SmsService, tracer trace.Tracer) service.SmsService {
	return &OpenTelemetrySmsService{
		svc:    svc,
		tracer: tracer,
	}
}

func (s *OpenTelemetrySmsService) Send(ctx context.Context, tplId string, args []string, numbers ...string) error {
	ctx, span := s.tracer.Start(ctx, "sms")
	defer span.End()
	span.SetAttributes(attribute.String("tpl", tplId))
	span.AddEvent("send message")
	err := s.svc.Send(ctx, tplId, args, numbers...)
	if err != nil {
		span.RecordError(err)
	}
	return err
}
