package main

import (
	"fmt"

	SharedConfig "github.com/github-profile/go-nats/shared/config"
	SharedContainer "github.com/github-profile/go-nats/shared/container"
	SharedContext "github.com/github-profile/go-nats/shared/context"
	SharedNats "github.com/github-profile/go-nats/shared/nats"
	"github.com/labstack/echo"
	"github.com/rs/zerolog/log"

	PubSubUsecase "github.com/github-profile/go-nats/domain/pubsub/usecase"

	PubSubHandler "github.com/github-profile/go-nats/domain/pubsub/delivery/http"
)

func main() {
	// - initialize echo
	e := echo.New()

	container := SharedContainer.DefaultContainer()
	config := container.MustGet("shared.config").(SharedConfig.ImmutableConfigInterface)
	nats := container.MustGet("shared.nats").(SharedNats.NatsInterface)

	// - gets nats connection
	natsSession, err := nats.GetNatsSession()
	if err != nil {
		msgError := fmt.Sprintf("Failed to open nats connection: %s", err.Error())
		log.Printf("[NatsSession] %s", msgError)
	}

	// - declaring context that will be used
	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ac := &SharedContext.ApplicationContext{
				Context:     c,
				Container:   container,
				NatsSession: natsSession,
			}
			return h(ac)
		}
	})

	pubSubUsecase := PubSubUsecase.NewPubSubUsecase(nats)

	PubSubHandler.PubSubHandler(e, pubSubUsecase)

	// - listener for Asynchronous Subscriptions
	go pubSubUsecase.SubsciberListener(natsSession)

	// - listener for Queue Subscriptions
	go pubSubUsecase.QueueGroupSubsciberListener(natsSession)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.GetPort())))
}
