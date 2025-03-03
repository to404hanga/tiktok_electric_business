package tencent

import (
	"context"
	"fmt"
	"tiktok_electric_business/sms/service"

	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
	"github.com/to404hanga/pkg404/logger"
	"github.com/to404hanga/pkg404/stl/transform"
)

// TencentSmsService 对接腾讯的短信服务
type TencentSmsService struct {
	client   *sms.Client
	appId    *string
	signName *string
	l        logger.Logger
}

var _ service.SmsService = (*TencentSmsService)(nil)

func NewTencentSmsService(client *sms.Client, appId, signName string, l logger.Logger) service.SmsService {
	return &TencentSmsService{
		client:   client,
		appId:    &appId,
		signName: &signName,
		l:        l,
	}
}

func (s *TencentSmsService) Send(ctx context.Context, tplId string, args []string, numbers ...string) error {
	req := sms.NewSendSmsRequest()
	req.SetContext(ctx)
	req.SmsSdkAppId = s.appId
	req.SignName = s.signName
	req.TemplateId = &tplId
	req.TemplateParamSet = transform.SliceFromSlice[string, *string](args, func(i int, arg string) *string {
		return &arg
	})
	req.PhoneNumberSet = transform.SliceFromSlice[string, *string](numbers, func(i int, number string) *string {
		return &number
	})

	resp, err := s.client.SendSms(req)
	s.l.Debug("Request Tencent SMS interface", logger.Any("req", req), logger.Any("resp", resp))
	if err != nil {
		return err
	}

	for _, statusPtr := range resp.Response.SendStatusSet {
		if statusPtr == nil {
			// 不可能进入这个分支
			continue
		}
		status := *statusPtr
		if status.Code == nil || *(status.Code) != "Ok" {
			return fmt.Errorf("send sms message failure, code: %s, msg: %s", *status.Code, *status.Message)
		}
	}
	return nil
}
