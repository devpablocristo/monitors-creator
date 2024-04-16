package task

import "time"

type Task struct {
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Status    Status
}

type Publisher interface {
	Publish(ownerId string, t Task) error
}

type Status int

const (
	StatusNone Status = iota
	StatusPending
	StatusCancelled
	StatusComplete
)

func (s Status) String() string {
	return [...]string{"none", "pending", "cancelled", "completed"}[s]
}
