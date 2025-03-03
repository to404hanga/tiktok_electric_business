package async

import (
	"context"
	"fmt"
	"sync"
	"tiktok_electric_business/sms/domain"
	"tiktok_electric_business/sms/repository"
	"tiktok_electric_business/sms/service"
	"time"

	"github.com/to404hanga/pkg404/limiter"
	"github.com/to404hanga/pkg404/logger"
)

const prefix = "teb:sms:async"

// AsyncSmsService 异步发送短信，当一段时间内错误次数达到阈值后，触发异步发送，持续一分钟后退出
type AsyncSmsService struct {
	svc         service.SmsService
	repo        repository.AsyncSmsRepository
	l           logger.Logger
	limit       limiter.Limiter
	shouldAsync bool
	lock        sync.RWMutex
}

var _ service.SmsService = (*AsyncSmsService)(nil)

func NewAsyncSmsService(svc service.SmsService, repo repository.AsyncSmsRepository, l logger.Logger, limit limiter.Limiter) service.SmsService {
	res := &AsyncSmsService{
		svc:   svc,
		repo:  repo,
		l:     l,
		limit: limit,
	}
	go func() {
		res.StartAsyncCycle()
	}()
	return res
}

// StartAsyncCycle 开始轮询异步发送短信
func (s *AsyncSmsService) StartAsyncCycle() {
	for {
		s.AsyncSend()
	}
}

// AsyncSend 异步发送短信
func (s *AsyncSmsService) AsyncSend() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	asyncSms, err := s.repo.PreemptWaitingSMS(ctx)
	cancel()
	switch err {
	case nil:
		ctx, cancel = context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		err = s.svc.Send(ctx, asyncSms.TplId, asyncSms.Args, asyncSms.Numbers...)
		if err != nil {
			s.l.Error("Failed to send asyncchronous SMS", logger.Error(err))
		}
		res := err == nil
		err = s.repo.ReportScheduleResult(ctx, asyncSms.Id, res)
		if err != nil {
			s.l.Error("Success to send asyncchronous SMS, but mark DB failure",
				logger.Error(err),
				logger.Bool("success", res),
				logger.Int64("id", asyncSms.Id),
			)
		}
	case repository.ErrWaitingSMSNotFound:
		time.Sleep(time.Second)
	default:
		s.l.Error("Failed to preempt the asyncronous SMS sending task", logger.Error(err))
		time.Sleep(time.Second)
	}
}

// Send 发送短信
func (s *AsyncSmsService) Send(ctx context.Context, tplId string, args []string, numbers ...string) error {
	if s.needAsync() {
		return s.repo.Add(ctx, domain.AsyncSms{
			TplId:    tplId,
			Args:     args,
			Numbers:  numbers,
			RetryMax: 3,
		})
	}
	err := s.svc.Send(ctx, tplId, args, numbers...)
	if err != nil {
		key := fmt.Sprintf("%s:send", prefix)
		limited, er := s.limit.Limit(ctx, key)
		if er != nil {
			s.l.Error("Limiter error", logger.Error(er))
			return er
		}
		if limited {
			s.l.Warn("The SMS is transferred to the asynchronous sending process", logger.Any("last_error", err))
			s.lock.Lock()
			s.shouldAsync = true
			s.lock.Unlock()

			go func() {
				// 定时一分钟后退出异步发送短信
				timer := time.NewTimer(time.Minute)
				<-timer.C
				s.lock.Lock()
				s.shouldAsync = false
				s.lock.Unlock()
				s.l.Info("Exit the asynchronous sending process")
			}()
		}
	}
	return err
}

func (s *AsyncSmsService) needAsync() bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.shouldAsync
}
