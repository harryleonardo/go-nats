package mock

import (
	nats "github.com/nats-io/nats.go"
	"github.com/stretchr/testify/mock"
)

//QoalaQueue ...
type QoalaQueue struct {
	mock.Mock
}

//GetNatsConfig ...
func (q *QoalaQueue) GetNatsConfig() (ec *nats.EncodedConn, err error) {
	args := q.Called()
	return args.Get(0).(*nats.EncodedConn), args.Error(1)
}

//GetNatsOfflineConfig ...
func (q *QoalaQueue) GetNatsOfflineConfig() (ec *nats.EncodedConn, err error) {
	args := q.Called()
	return args.Get(0).(*nats.EncodedConn), args.Error(1)
}

//PublishMessage ...
func (q *QoalaQueue) PublishMessage(ec *nats.EncodedConn, QueueSubject string, data interface{}) (err error) {
	args := q.Called(ec, QueueSubject, data)
	return args.Error(0)
}
