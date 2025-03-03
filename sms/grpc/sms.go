package grpc

import (
	"context"
	smsv1 "tiktok_electric_business/api/proto/gen/sms/v1"
	"tiktok_electric_business/sms/service"

	"google.golang.org/grpc"
)

type SmsServiceServer struct {
	smsv1.UnimplementedSmsServiceServer
	svc service.SmsService
}

func NewSmsServiceServer(svc service.SmsService) *SmsServiceServer {
	return &SmsServiceServer{
		svc: svc,
	}
}

func (s *SmsServiceServer) Register(srv grpc.ServiceRegistrar) {
	smsv1.RegisterSmsServiceServer(srv, s)
}

func (s *SmsServiceServer) Send(ctx context.Context, req *smsv1.SmsSendRequest) (*smsv1.SmsSendResponse, error) {
	err := s.svc.Send(ctx, req.GetTplId(), req.GetArgs(), req.GetNumbers()...)
	return &smsv1.SmsSendResponse{}, err
}
