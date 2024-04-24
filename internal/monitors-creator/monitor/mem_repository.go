package monitor

import (
	"errors"
	"monitors-creator/internal/platform/memdb"
)

func NewMemoryRepo(db *memdb.MemDB) MonitorRepositoryPort {
	return &MemRepo{
		db: *db,
	}
}

type MemRepo struct {
	db memdb.MemDB
}

func (m *MemRepo) Create(monitor *Monitor) error {
	m.db[monitor.ID.String()] = monitor
	return nil
}

func (m *MemRepo) Update(monitor *Monitor) error {
	m.db[monitor.ID.String()] = monitor
	return nil
}

func (m *MemRepo) Get(id string) (*Monitor, error) {
	value, ok := m.db[id]
	if !ok {
		return nil, errors.New("id not found")
	}

	monitor, ok := value.(*Monitor)
	if !ok {
		return nil, errors.New("value is not a monitor")
	}

	return monitor, nil
}

func (m *MemRepo) GetAll() ([]Monitor, error) {
	var monitors []Monitor
	for _, value := range m.db {
		monitor, ok := value.(*Monitor)
		if !ok {
			return nil, errors.New("value is not a monitor")
		}

		monitors = append(monitors, monitor)
	}

	return monitors, nil
}

func (m *MemRepo) Delete(id string) error {
	delete(m.db, id)
	return nil
}
