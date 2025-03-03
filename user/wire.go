//go:build wireinject

package main

import (
	"tiktok_electric_business/pkg/wego"
	"tiktok_electric_business/user/grpc"
	"tiktok_electric_business/user/ioc"
	"tiktok_electric_business/user/repository"
	"tiktok_electric_business/user/repository/cache"
	"tiktok_electric_business/user/repository/dao"
	"tiktok_electric_business/user/service"

	"github.com/google/wire"
)

var thirdProvider = wire.NewSet(
	ioc.InitDB,
	ioc.InitEtcdClient,
	ioc.InitLimiter,
	ioc.InitLogger,
	ioc.InitRedis,
)

func Init() *wego.App {
	wire.Build(
		thirdProvider,
		cache.NewRedisUserCache,
		dao.NewGormUserDAO,
		repository.NewCachedUserRepository,
		service.NewLimitedUserService,
		grpc.NewUserServiceServer,
		ioc.InitGrpcxServer,
		wire.Struct(new(wego.App), "GrpcServer"),
	)
	return new(wego.App)
}
