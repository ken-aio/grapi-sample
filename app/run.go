package app

import (
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"github.com/ken-aio/grapi-sample/app/server"
)

// Run starts the grapiserver.
func Run() error {
	s := grapiserver.New(
		grapiserver.WithDefaultLogger(),
		grapiserver.WithServers(
			server.NewBookServiceServer(),
		),
	)
	return s.Serve()
}
