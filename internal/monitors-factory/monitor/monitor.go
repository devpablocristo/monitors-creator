package monitor

import (
	"monitors-creator/internal/monitors-factory/monitor/brand"
	"time"
)

// Metric representa una métrica específica que se está monitoreando.
type Metric struct {
	Type string
}

// Threshold representa el umbral específico que no debe ser cruzado por la métrica.
type Threshold struct {
	Limit string
}

// Monitor es la estructura que representa un monitor genérico.
type Monitor struct {
	Brand     brand.Brand
	Metric    Metric
	Threshold Threshold
	CreatedAt time.Time
}
