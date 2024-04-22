package usecase

import "monitors-creator/internal/monitors-creator/monitor"

type MonitorUsecase struct {
	repository monitor.MonitorRepositoryPort
}

func NewMonitorUsecase(repo monitor.MonitorRepositoryPort) monitor.MonitorUsecasePort {
	return &MonitorUsecase{repository: repo}
}

func (u MonitorUsecase) CreateMonitor(monitor *monitor.Monitor) error {
	if err := u.repository.Create(monitor); err != nil {
		return err
	}
	monitor.ID = "999"
	return nil
}
