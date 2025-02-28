package domain

import "time"

type User struct {
	Id         int64
	Email      string
	Password   string
	Nickname   string
	Phone      string
	UserType   UserType
	WechatInfo WechatInfo
	CreatedAt  time.Time
}

type WechatInfo struct {
	OpenId  string
	UnionId string
}

type UserType uint8

const (
	USER_TYPE_UNKNOWN UserType = iota
	USER_TYPE_CUSTOMER
	USER_TYPE_MERCHANT
)
