package monitor

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type MonitorRepository struct {
	db *sql.DB
}

func NewMonitorRepository(db *sql.DB) *MonitorRepository {
	return &MonitorRepository{
		db: db,
	}
}

func (repo *MonitorRepository) Save(monitor *MonitorDaoData) error {
	query := `INSERT INTO monitors (ID, Brand, Metric, Threshold, CreatedAt, UpdatedAt, Deleted) VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := repo.db.Exec(query, monitor.ID.ID, monitor.Brand.Name, monitor.Metric.Type, monitor.Threshold.Limit, monitor.CreatedAt, monitor.UpdatedAt, monitor.Deleted)
	return err
}

func (repo *MonitorRepository) FindByID(id MonID) (*MonitorDaoData, error) {
	return nil, fmt.Errorf("not implemented")
}

func (repo *MonitorRepository) Update(monitor *MonitorDaoData) error {
	return fmt.Errorf("not implemented")
}

func (repo *MonitorRepository) Delete(id *MonID) error {
	return fmt.Errorf("not implemented")
}

func (repo *MonitorRepository) SoftDelete(id *MonID) error {
	return fmt.Errorf("not implemented")
}
