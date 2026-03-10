package greeter

import (
	"context"
	"errors"
	"log"

	"github.com/kodaikumatani/grpc-authentication-go/internal/authn"
	pb "github.com/kodaikumatani/grpc-authentication-go/pkg/pb/greeter"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

// handler is used to implement helloworld.GreeterServer.
type handler struct {
	pb.UnimplementedGreeterServer
}

func NewHandler() pb.GreeterServer {
	return &handler{}
}

// SayHello implements helloworld.GreeterServer
func (s *handler) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	uid, ok := ctx.Value(authn.UIDKey{}).(string)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, ErrUserNotFound.Error())
	}

	return &pb.HelloReply{Message: "Hello " + uid}, nil
}
