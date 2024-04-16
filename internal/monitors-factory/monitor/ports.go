package monitor

type MonitorRepositoryPort interface {
	Save(*MonitorDaoData)
	FindByID(*MonID)
	Update(*MonitorDaoData)
	Delete(*MonID)
	SoftDelete(*MonID)
}
