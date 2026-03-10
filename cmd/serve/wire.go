//go:build wireinject
// +build wireinject

package main

import (
	"context"

	"github.com/google/wire"
	"github.com/kodaikumatani/grpc-authentication-go/internal"
	"github.com/kodaikumatani/grpc-authentication-go/internal/app"
	"github.com/kodaikumatani/grpc-authentication-go/internal/authn"
)

type services struct {
	*app.Registrar
	Verifier authn.Verifier
}

var set = wire.NewSet(
	internal.Set,
	wire.Struct(new(services), "*"),
)

func initializeServices(ctx context.Context) (*services, error) {
	panic(wire.Build(set))
}
