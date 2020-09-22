package usecase

import (
	"log"
	"sync"

	"github.com/github-profile/go-nats/domain/pubsub"
	"github.com/github-profile/go-nats/models"
	SharedContext "github.com/github-profile/go-nats/shared/context"
	SharedNats "github.com/github-profile/go-nats/shared/nats"
	"github.com/labstack/echo"
	nats "github.com/nats-io/nats.go"
)

const (
	Topic = "PUBLISH_SUBSCRIBE_SCENARIO"
)

type usecase struct {
	natsSession SharedNats.NatsInterface
}

// NewPubSubUsecase ...
func NewPubSubUsecase(natsSession SharedNats.NatsInterface) pubsub.Usecase {
	return &usecase{natsSession: natsSession}
}

func (u usecase) Publisher(c echo.Context) error {
	ctx := c.(*SharedContext.ApplicationContext)
	natsSession := ctx.NatsSession

	payload := models.IncomingMessage{}
	if err := ctx.Bind(&payload); err != nil {
		// - failed to bind data
		log.Println("Err Bind : ", err)
		return err
	}

	// - publish data
	err := u.natsSession.PublishMessage(natsSession, Topic, payload)
	if err != nil {
		// - failed to push
		log.Println("Err Publish : ", err)
		return err
	}

	log.Println("Publish Data : ", payload)
	return nil
}

func (u usecase) SubsciberListener(natsSession *nats.EncodedConn) error {
	log.Println("[SubsciberListener] is running")
	// - use a WaitGroup to wait for a message to arrive
	wg := sync.WaitGroup{}
	wg.Add(1)

	if _, err := natsSession.Subscribe(Topic, func(msg *nats.Msg) {
		log.Printf("[SubsciberListener] Incoming Data : %s", msg.Data)
		wg.Done()
	}); err != nil {
		log.Fatal(err)
	}

	// Wait for a message to come in
	wg.Wait()

	return nil
}

// QueueGroupSubsciberListenerOne is a simple implementation of QueueGroupSubsciber on Nats;
func (u usecase) QueueGroupSubsciberListener(natsSession *nats.EncodedConn) error {
	log.Println("[QueueGroupSubsciberListener] is running")

	// Subscribe
	natsSession.QueueSubscribe(Topic, "worker", func(m *nats.Msg) {
		log.Printf("[QueueGroupSubsciberListener] Incoming Data : %s", m.Data)
	})
	return nil
}
