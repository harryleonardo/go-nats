package context

import (
	"github.com/fgrosse/goldi"
	"github.com/labstack/echo"
	nats "github.com/nats-io/nats.go"
)

// ApplicationContext ...
type ApplicationContext struct {
	echo.Context
	Container   *goldi.Container
	NatsSession *nats.EncodedConn
}
