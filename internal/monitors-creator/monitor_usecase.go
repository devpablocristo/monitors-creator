package usecase

import (
	"context"
	"monitors-creator/internal/monitors-creator/monitor"
)

type MonitorUsecase struct {
	repository monitor.MonitorRepositoryPort
}

func NewMonitorUsecase(repo monitor.MonitorRepositoryPort) monitor.MonitorUsecasePort {
	return &MonitorUsecase{repository: repo}
}

func (u MonitorUsecase) CreateMonitor(ctx context.Context, monitor *monitor.Monitor) error {
	if err := u.repository.Create(monitor); err != nil {
		return err
	}
	monitor.ID = "999"
	return nil
}

func (u MonitorUsecase) GetMonitor(ctx context.Context, id string) (*monitor.Monitor, error) {
	return u.repository.Get(id)
}

func (u MonitorUsecase) UpdateMonitor(ctx context.Context, monitor *monitor.Monitor) error {
	return u.repository.Update(monitor)
}

func (u MonitorUsecase) DeleteMonitor(ctx context.Context, id string) error {
	return u.repository.Delete(id)
}

func (u MonitorUsecase) GetAllMonitors(ctx context.Context) ([]*monitor.Monitor, error) {
	return u.repository.GetAll()
}
