package repository

import (
	"context"
	"tiktok_electric_business/sms/domain"
	"tiktok_electric_business/sms/repository/dao"

	"github.com/ecodeclub/ekit/sqlx"
)

type asyncSmsRepository struct {
	dao dao.AsyncSmsDAO
}

var _ AsyncSmsRepository = (*asyncSmsRepository)(nil)

func NewAsyncSmsRepository(dao dao.AsyncSmsDAO) AsyncSmsRepository {
	return &asyncSmsRepository{
		dao: dao,
	}
}

// Add 新增一条异步短信记录
func (s *asyncSmsRepository) Add(ctx context.Context, sms domain.AsyncSms) error {
	return s.dao.Insert(ctx, dao.AsyncSms{
		Config: sqlx.JsonColumn[dao.SmsConfig]{
			Val: dao.SmsConfig{
				TplId:   sms.TplId,
				Args:    sms.Args,
				Numbers: sms.Numbers,
			},
			Valid: true,
		},
		RetryMax: sms.RetryMax,
	})
}

// PreemptWaitingSMS 获取最近1分钟内的未发送的异步短信
func (s *asyncSmsRepository) PreemptWaitingSMS(ctx context.Context) (domain.AsyncSms, error) {
	sms, err := s.dao.GetWaitingSMS(ctx)
	if err != nil {
		return domain.AsyncSms{}, err
	}
	return domain.AsyncSms{
		Id:       sms.Id,
		TplId:    sms.Config.Val.TplId,
		Args:     sms.Config.Val.Args,
		Numbers:  sms.Config.Val.Numbers,
		RetryMax: sms.RetryMax,
	}, nil
}

// ReportScheduleResult 标记任务执行结果
func (s *asyncSmsRepository) ReportScheduleResult(ctx context.Context, id int64, success bool) error {
	if success {
		return s.dao.MarkSuccess(ctx, id)
	}
	return s.dao.MarkFailed(ctx, id)
}
