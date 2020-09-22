package pubsub

import (
	"github.com/labstack/echo"
	nats "github.com/nats-io/nats.go"
)

type Usecase interface {
	Publisher(c echo.Context) error
	SubsciberListener(natsSession *nats.EncodedConn) error
	QueueGroupSubsciberListener(natsSession *nats.EncodedConn) error
}
