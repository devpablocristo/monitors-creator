package developer

import (
	"fmt"
	"users_example/internal/supervisor/developer/task"
)

type Repository interface {
	Save(d *Developer) error
	Get(id string) (*Developer, error)
	Delete(id string) error
	SearchByStatus(status task.Status) ([]Developer, error)
}

type Developer struct {
	ID        string
	Name      string
	Team      string
	Seniority Seniority
	Task      task.Task
}

func (d Developer) IsBusy() bool {
	return d.Task.Status == task.StatusPending
}

type Seniority int

func SeniorityFromString(value string) (Seniority, error) {
	switch value {
	case "senior":
		return Senior, nil
	case "semi_senior":
		return SemiSenior, nil
	case "analyst":
		return Analyst, nil
	case "junior":
		return Junior, nil
	default:
		return 0, fmt.Errorf("%s is not a valid seniority", value)
	}
}

const (
	Senior Seniority = iota
	SemiSenior
	Analyst
	Junior
)

func (s Seniority) String() string {
	return [...]string{"senior", "semi_senior", "analyst", "junior"}[s]
}
