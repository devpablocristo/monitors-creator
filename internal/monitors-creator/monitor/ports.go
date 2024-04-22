package monitor

type MonitorRepositoryPort interface {
	Create(*Monitor) error
	// Save(*MonitorDaoData)
	// FindByID(*MonID)
	// Update(*MonitorDaoData)
	// Delete(*MonID)
	// SoftDelete(*MonID)
}

type MonitorUsecasePort interface {
	CreateMonitor(*Monitor) error
}
