package dao

import (
	"context"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GormAsyncSmsDAO struct {
	db *gorm.DB
}

var _ AsyncSmsDAO = (*GormAsyncSmsDAO)(nil)

func NewGormAsyncSmsDAO(db *gorm.DB) AsyncSmsDAO {
	return &GormAsyncSmsDAO{
		db: db,
	}
}

// Insert 往异步短信表中新增一条短信
func (s *GormAsyncSmsDAO) Insert(ctx context.Context, sms AsyncSms) error {
	return s.db.WithContext(ctx).Create(&sms).Error
}

// GetWaitingSMS 获取最近1分钟内的未发送的异步短信
func (s *GormAsyncSmsDAO) GetWaitingSMS(ctx context.Context) (AsyncSms, error) {
	var sms AsyncSms
	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 为了避开一些偶发性失败，只找1分钟前的异步短信发送
		now := time.Now().UnixMilli()
		endTime := now - time.Minute.Milliseconds()
		err := tx.Clauses(clause.Locking{
			Strength: "UPDATE", // 加行级锁
		}).Where("updated_at < ? AND status = ?", endTime, asyncStatusWaiting).First(&sms).Error
		if err != nil {
			return err
		}
		err = tx.Model(&AsyncSms{}).Where("id = ?", sms.Id).Updates(map[string]any{
			"retry_cnt":  gorm.Expr("retry_cnt + 1"),
			"updated_at": now,
		}).Error
		return err
	})
	return sms, err
}

// MarkSuccess 标记异步短信为已成功
func (s *GormAsyncSmsDAO) MarkSuccess(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Model(&AsyncSms{}).Where("id = ?", id).Updates(map[string]any{
		"updated_at": time.Now().UnixMilli(),
		"status":     asyncStatusSuccess,
	}).Error
}

// MarkFailed 标记异步短信为已失败
func (s *GormAsyncSmsDAO) MarkFailed(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Model(&AsyncSms{}).Where("id =?", id).Updates(map[string]any{
		"updated_at": time.Now().UnixMilli(),
		"status":     asyncStatusFailed,
	}).Error
}
