package dao

import (
	"context"
	"database/sql"
)

//go:generate mockgen -source=./types.go -package=daomocks -destination=./mocks/user.mock.go UserDAO
type UserDAO interface {
	Insert(ctx context.Context, user User) error
	FindByEmail(ctx context.Context, email string) (User, error)
	UpdateNonZeroFields(ctx context.Context, user User) error
	FindById(ctx context.Context, id int64) (User, error)
	FindByPhone(ctx context.Context, phone string) (User, error)
	FindByWechat(ctx context.Context, openId string) (User, error)
	DeleteById(ctx context.Context, id int64) error
}

type User struct {
	Id            int64          `gorm:"primaryKey,autoIncrement"`
	Email         sql.NullString `gorm:"unique"`
	Password      string
	Phone         sql.NullString `gorm:"unique"`
	Nickname      sql.NullString `gorm:"type=varchar(50)"`
	UserType      uint8          `gorm:"type=tinyint;index"` // 用户类型，1=顾客，2=商家
	WechatOpenId  sql.NullString `gorm:"type=varchar(256);unique"`
	WechatUnionId sql.NullString `gorm:"type=varchar(255)"`
	CreatedAt     int64          `gorm:"autoCreateTime:milli"`
	UpdatedAt     int64          `gorm:"autoUpdateTime:milli"`
	DeletedAt     int64          `gorm:"index"`
}
