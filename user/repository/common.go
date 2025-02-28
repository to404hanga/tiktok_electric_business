package repository

import (
	"database/sql"
	"tiktok_electric_business/user/domain"
	"tiktok_electric_business/user/repository/dao"
	"time"
)

func convertToDomain(user dao.User) domain.User {
	return domain.User{
		Id:       user.Id,
		Email:    user.Email.String,
		Password: user.Password,
		Nickname: user.Nickname.String,
		Phone:    user.Phone.String,
		UserType: domain.UserType(user.UserType),
		WechatInfo: domain.WechatInfo{
			OpenId:  user.WechatOpenId.String,
			UnionId: user.WechatUnionId.String,
		},
		CreatedAt: time.UnixMilli(user.CreatedAt),
	}
}

func convertToEntity(user domain.User) dao.User {
	return dao.User{
		Id: user.Id,
		Email: sql.NullString{
			String: user.Email,
			Valid:  user.Email != "",
		},
		Password: user.Password,
		Phone: sql.NullString{
			String: user.Phone,
			Valid:  user.Phone != "",
		},
		Nickname: sql.NullString{
			String: user.Nickname,
			Valid:  user.Nickname != "",
		},
		UserType: uint8(user.UserType),
		WechatOpenId: sql.NullString{
			String: user.WechatInfo.OpenId,
			Valid:  user.WechatInfo.OpenId != "",
		},
		WechatUnionId: sql.NullString{
			String: user.WechatInfo.UnionId,
			Valid:  user.WechatInfo.UnionId != "",
		},
	}
}
