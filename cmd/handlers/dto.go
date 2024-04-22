package handlers

import (
	"monitors-creator/internal/monitors-creator/monitor"
	"time"

	"github.com/corthmann/go-time-intervals/timeinterval"
)

func MonitorResponse(m monitor.Monitor) *MonitorPayload {
	//TODO FIX THIS SHIT
	return &MonitorPayload{
		ID: "123",
	}
}

func mapPayloadToDomain(payload MonitorPayload) (monitor.Monitor, error) {
	// TODO: CHECAR SE PRECISAMOS CHECAR ERRO EM TIME.PARSE. SE TIME = NIL DÁ ERRO.
	layout := "15:04"
	startTime, err := time.Parse(layout, payload.CustomerHoursStart)
	if err != nil {
		return monitor.Monitor{}, err
	}
	endTime, err := time.Parse(layout, payload.CustomerHoursEnd)
	if err != nil {
		return monitor.Monitor{}, err
	}
	return monitor.Monitor{
		ID: monitor.ID(payload.ID),
		Brand: monitor.Brand{Name: monitor.Name(payload.Brand),
			CustomerHours: monitor.CustomerHours{
				SupportHours: timeinterval.Interval{
					StartsAt: startTime,
					EndsAt:   endTime},
			},
		},
		BusinessUnit: monitor.BusinessUnit{Name: monitor.Name(payload.BusinessUnit)},
		Product:      monitor.Product{Name: monitor.Name(payload.Product)},
	}, nil
}

type MonitorPayload struct {
	ID                 string `json:"id" validate:"required"`
	Brand              string `json:"brand" validate:"required"`
	BusinessUnit       string `json:"businessUnit"`
	Product            string `json:"product"`
	CustomerHoursStart string `json:"customerHoursStart"`
	CustomerHoursEnd   string `json:"customerHoursEnd"`
}
