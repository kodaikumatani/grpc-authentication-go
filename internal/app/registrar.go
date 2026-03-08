package app

import (
	pb "github.com/kodaikumatani/grpc-authentication-go/pkg/pb/greeter"
	"google.golang.org/grpc"
)

type Registrar struct {
	handler pb.GreeterServer
}

func NewRegistrar(
	handler pb.GreeterServer,
) *Registrar {
	return &Registrar{
		handler: handler,
	}
}

func (r *Registrar) Register(app *grpc.Server) *grpc.Server {
	pb.RegisterGreeterServer(app, r.handler)

	return app
}
