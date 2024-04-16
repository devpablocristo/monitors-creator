package presenter

import (
	"users_example/internal/supervisor/developer"
	"users_example/internal/supervisor/developer/task"
)

type jsonDeveloper struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Team      string    `json:"team"`
	Seniority string    `json:"seniority"`
	Task      *jsonTask `json:"assigned_task"`
}

func Developer(d *developer.Developer) *jsonDeveloper {
	toReturn := &jsonDeveloper{
		ID:        d.ID,
		Name:      d.Name,
		Team:      d.Team,
		Seniority: d.Seniority.String(),
		Task:      nil,
	}

	if d.Task.Status != task.StatusNone {
		toReturn.Task = &jsonTask{
			Name:      d.Task.Name,
			CreatedAt: d.Task.CreatedAt.String(),
			UpdatedAt: d.Task.UpdatedAt.String(),
			Status:    d.Task.Status.String(),
		}
	}

	return toReturn
}
