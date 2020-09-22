package nats

import (

	// nats "github.com/nats-io/go-nats"
	SharedConfig "github.com/github-profile/go-nats/shared/config"
	nats "github.com/nats-io/nats.go"
)

type (
	// NatsInterface ...
	NatsInterface interface {
		GetNatsSession() (ec *nats.EncodedConn, err error)
		PublishMessage(ec *nats.EncodedConn, NatsSubject string, data interface{}) (err error)
	}

	natsQueue struct {
		config SharedConfig.ImmutableConfigInterface
	}
)

// - Connect to  Config
func (q *natsQueue) GetNatsSession() (ec *nats.EncodedConn, err error) {
	natsEndpoint := q.config.GetNATSHost()
	if natsEndpoint == "" {
		natsEndpoint = nats.DefaultURL
	}

	nc, err := nats.Connect(natsEndpoint)
	if err != nil {
		return nil, err
	}

	// - encoded
	ec, err = nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		return nil, err
	}

	return ec, nil
}

func (q *natsQueue) PublishMessage(ec *nats.EncodedConn, QueueSubject string, data interface{}) (err error) {
	err = ec.Publish(QueueSubject, data)
	if err != nil {
		return err
	}

	return nil
}

// NewService ...
func NewService(config SharedConfig.ImmutableConfigInterface) NatsInterface {
	if config == nil {
		panic("shared immutable config is required")
	}

	return &natsQueue{config}
}
