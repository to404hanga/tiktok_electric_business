//go:build wireinject

package main

import (
	"tiktok_electric_business/pkg/wego"
	"tiktok_electric_business/sms/grpc"
	"tiktok_electric_business/sms/ioc"

	"github.com/google/wire"
)

func Init() *wego.App {
	wire.Build(
		ioc.InitLogger,
		ioc.InitEtcdClient,
		ioc.InitSmsMemoryService,
		grpc.NewSmsServiceServer,
		ioc.InitGrpcxServer,
		wire.Struct(new(wego.App), "GrpcServer"),
	)
	return new(wego.App)
}
