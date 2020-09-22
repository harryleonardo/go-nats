package http

import (
	"github.com/github-profile/go-nats/domain/pubsub"
	"github.com/labstack/echo"
)

type handlerPubSubScenario struct {
	usecase pubsub.Usecase
}

// PubSubHandler ...
func PubSubHandler(e *echo.Echo, usecase pubsub.Usecase) {
	handler := handlerPubSubScenario{
		usecase: usecase,
	}

	e.POST("/api/publish-subscribe", handler.PublishSubscribeScenario)
}

func (h handlerPubSubScenario) PublishSubscribeScenario(c echo.Context) error {
	err := h.usecase.Publisher(c)
	if err != nil {
		return err
	}

	return nil
}
