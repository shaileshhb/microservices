package eventbus

import (
	"github.com/shaileshhb/microservices/query/app/entity"
	"gorm.io/gorm"
)

type EventBusService struct {
	postDB    *gorm.DB
	commentDB *gorm.DB
}

func NewEventBusService(postDB, commentDB *gorm.DB,
) *EventBusService {
	return &EventBusService{
		postDB:    postDB,
		commentDB: commentDB,
	}
}

func (service *EventBusService) ListenEvent(event *entity.Event) error {

	return nil
}
