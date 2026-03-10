package internal

import (
	"github.com/google/wire"
	"github.com/kodaikumatani/grpc-authentication-go/internal/app"
	"github.com/kodaikumatani/grpc-authentication-go/internal/authn/firebase"
)

var Set = wire.NewSet(
	app.Set,
	firebase.Set,
)
