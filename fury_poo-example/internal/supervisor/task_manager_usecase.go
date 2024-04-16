package supervisor

import (
	"fmt"
	"time"
	"users_example/internal/supervisor/developer"
	"users_example/internal/supervisor/developer/task"
)

type TaskManagementUseCase interface {
	ScheduleTask(id string, taskName string) (*developer.Developer, error)
	CompleteTask(id string) (*developer.Developer, error)
}

type useCaseTaskManagement struct {
	taskPublisher task.Publisher
	devRepo       developer.Repository
}

func NewTaskManagementUseCase(taskPublisher task.Publisher, devRepo developer.Repository) *useCaseTaskManagement {
	return &useCaseTaskManagement{taskPublisher: taskPublisher, devRepo: devRepo}
}

func (u useCaseTaskManagement) ScheduleTask(id string, taskName string) (*developer.Developer, error) {
	// chequeo si el usuario existe
	dev, err := u.devRepo.Get(id)
	if err != nil {
		return nil, err
	}

	if dev == nil {
		return nil, fmt.Errorf("user_id:%s not found, task cannot be scheduled", id)
	}

	newTask := task.Task{
		Name:      taskName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Status:    task.StatusPending,
	}

	// le asigno la tarea y lo guardo
	dev.Task = newTask
	if err := u.devRepo.Save(dev); err != nil {
		return nil, err
	}

	return dev, nil
}

func (u useCaseTaskManagement) CompleteTask(id string) (*developer.Developer, error) {
	dev, err := u.devRepo.Get(id)
	if err != nil {
		return nil, err
	}

	if dev == nil {
		return nil, fmt.Errorf("user_id:%s not found, task cannot be completed", id)
	}

	if dev.Task.Status == task.StatusNone {
		return nil, fmt.Errorf("user_id:%s doesn't have any task assigned", id)
	}

	dev.Task.UpdatedAt = time.Now()
	dev.Task.Status = task.StatusComplete
	if err := u.devRepo.Save(dev); err != nil {
		return nil, err
	}

	if err := u.taskPublisher.Publish(id, dev.Task); err != nil {
		return nil, err
	}

	return dev, nil
}
