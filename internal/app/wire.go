package app

import (
	"github.com/google/wire"
	"github.com/kodaikumatani/grpc-authentication-go/internal/app/greeter"
)

var Set = wire.NewSet(
	NewRegistrar,
	greeter.NewHandler,
)
