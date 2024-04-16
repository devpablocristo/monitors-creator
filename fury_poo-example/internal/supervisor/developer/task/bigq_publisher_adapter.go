package task

import (
	"github.com/mercadolibre/go-meli-toolkit/gobigqueue"
)

type bigQPublisher struct {
	publisher gobigqueue.Publisher
}

func NewBigQPublisher(client gobigqueue.Publisher) *bigQPublisher {
	return &bigQPublisher{
		client,
	}
}

func (q bigQPublisher) Publish(ownerId string, t Task) error {
	msg := map[string]interface{}{
		"owner_id":    ownerId,
		"task_name":   t.Name,
		"task_status": t.Status,
		"updated_at":  t.UpdatedAt,
		"created_at":  t.CreatedAt,
	}

	return q.publisher.Send(&gobigqueue.Payload{Msg: msg})
}
