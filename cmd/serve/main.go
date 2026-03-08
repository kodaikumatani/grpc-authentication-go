package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/kodaikumatani/grpc-authentication-go/internal/app"
	"github.com/kodaikumatani/grpc-authentication-go/internal/app/greeter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	handler := new(greeter.Handler)
	s = app.NewRegistrar(handler).Register(s)
	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
