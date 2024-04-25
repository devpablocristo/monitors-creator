package usecase

import (
	"context"
	"fmt"
	"monitors-creator/internal/monitors-creator/monitor"
	"monitors-creator/internal/platform/restclient"
)

type MonitorUsecase struct {
	repository monitor.MonitorRepositoryPort
	datadog    *restclient.EndpointType
}

func NewMonitorUsecase(repo monitor.MonitorRepositoryPort, datadog *restclient.EndpointType) monitor.MonitorUsecasePort {
	return &MonitorUsecase{
		repository: repo,
		datadog:    datadog,
	}
}

func (u MonitorUsecase) CreateMonitor(ctx context.Context, monitor *monitor.Monitor) error {
	monitor.ID = monitor.ID.Create()
	if err := u.repository.Create(monitor); err != nil {
		return err
	}
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

func (u MonitorUsecase) GetAllMonitors(ctx context.Context) ([]monitor.Monitor, error) {
	var r monitor.DatadogResponse
	if err := u.datadog.Get(ctx, &r); err != nil {
		return nil, err
	}
	fmt.Println(r)
	return u.repository.GetAll()
}
