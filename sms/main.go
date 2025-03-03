package main

import "tiktok_electric_business/pkg/viperwatch"

func main() {
	viperwatch.InitViperWatch()
	app := Init()
	if err := app.GrpcServer.Serve(); err != nil {
		panic(err)
	}
}
