package greeter

import (
	"context"
	"log"

	pb "github.com/kodaikumatani/grpc-authentication-go/pkg/pb/greeter"
)

// Handler is used to implement helloworld.GreeterServer.
type Handler struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *Handler) SayHello(_ context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}
