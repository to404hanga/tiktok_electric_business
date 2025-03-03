package wego

import (
	"github.com/to404hanga/pkg404/ginx"
	"github.com/to404hanga/pkg404/grpcx"
	"github.com/to404hanga/pkg404/saramax"
)

type App struct {
	GrpcServer *grpcx.Server
	WebServer  *ginx.Server
	Consumers  []saramax.Consumer
}
