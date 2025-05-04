package orchestrator

import (
	"github.com/Aman5681/notify/internal/orchestrator/handlers"
	"github.com/Aman5681/notify/internal/payload"
)

type Service struct{}

func NewService() *Service {
	InitActionMap()
	return &Service{}
}

type ActionHandler func(payload payload.Payload) (string, error)

var actionHandlerMap map[string]ActionHandler

func InitActionMap() {
	actionHandlerMap = map[string]ActionHandler{
		"generate": handlers.HandleGenerate,
	}
}

func (service *Service) GetHandler(action string) (ActionHandler, bool) {
	h, ok := actionHandlerMap[action]
	return h, ok
}
