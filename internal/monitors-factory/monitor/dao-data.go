package monitor

import (
	"time"

	"github.com/google/uuid"

	"monitors-creator/internal/monitors-factory/monitor/brand"
)

type MonID struct {
	ID string
}

// MonitorDaoData es la representación de la estructura de datos para la base de datos.
type MonitorDaoData struct {
	ID        MonID       `db:"ID"`
	Brand     brand.Brand `db:"Brand"`
	Metric    Metric      `db:"Metric"`
	Threshold Threshold   `db:"Threshold"`
	CreatedAt time.Time   `db:"created_at"`
	UpdatedAt time.Time   `db:"updated_at"`
	Deleted   bool        `db:"deleted"`
}

func ToDao(dom *Monitor) *MonitorDaoData {
	UUID := uuid.New()
	ID := MonID{
		ID: UUID.String(),
	}
	return &MonitorDaoData{
		ID:        ID,
		Brand:     dom.Brand,
		Metric:    dom.Metric,
		Threshold: dom.Threshold,
		CreatedAt: dom.CreatedAt,
		UpdatedAt: time.Now(),
		Deleted:   false,
	}
}

func ToDomain(dao *MonitorDaoData) *Monitor {
	return &Monitor{
		Brand:     dao.Brand,
		Metric:    dao.Metric,
		Threshold: dao.Threshold,
		CreatedAt: dao.CreatedAt,
	}
}
