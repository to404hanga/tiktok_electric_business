package repository

import "tiktok_electric_business/user/repository/dao"

var (
	ErrUserDuplicate = dao.ErrUserDuplicate
	ErrUserNotFound  = dao.ErrDataNotFound
)
