package monitor

import "context"

type MonitorRepositoryPort interface {
	Create(*Monitor) error
	Update(*Monitor) error
	Get(string) (*Monitor, error)
	GetAll() ([]*Monitor, error)
	Delete(string) error
}

type MonitorUsecasePort interface {
	CreateMonitor(context.Context, *Monitor) error
	GetMonitor(context.Context, string) (*Monitor, error)
	UpdateMonitor(context.Context, *Monitor) error
	DeleteMonitor(context.Context, string) error
	GetAllMonitors(context.Context) ([]*Monitor, error)
}
