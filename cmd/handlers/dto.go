package handlers

import (
	"monitors-creator/internal/monitors-creator/monitor"
)

func MonitorResponse(m monitor.Monitor) *MonitorPayload {
	//TODO FIX THIS SHIT
	return &MonitorPayload{
		ID: "123",
	}
}

func mapPayloadToDomain(payload MonitorPayload) (monitor.Monitor, error) {
	// TODO: CHECAR SE PRECISAMOS CHECAR ERRO EM TIME.PARSE. SE TIME = NIL DÁ ERRO.
	return monitor.Monitor{
		ID: monitor.ID(payload.ID),
		MonthlyTPV: monitor.MonthlyTPV{
			USD: monitor.Currency(payload.TPV)},
	}, nil
}

// func mapPayloadToDomain(payload MonitorPayload) (monitor.Monitor, error) {
// 	// TODO: CHECAR SE PRECISAMOS CHECAR ERRO EM TIME.PARSE. SE TIME = NIL DÁ ERRO.
// 	layout := "15:04"
// 	startTime, err := time.Parse(layout, payload.Brand.CustomerHours.WorkingHours.Start)
// 	if err != nil {
// 		return monitor.Monitor{}, err
// 	}

// 	endTime, err := time.Parse(layout, payload.Brand.CustomerHours.WorkingHours.Start)
// 	if err != nil {
// 		return monitor.Monitor{}, err
// 	}
// 	return monitor.Monitor{
// 		ID: monitor.ID(payload.ID),
// 		Brand: monitor.Brand{Name: monitor.Name(payload.Brand.Name),
// 			CustomerHours: monitor.CustomerHours{
// 				SupportHours: timeinterval.Interval{
// 					StartsAt: startTime,
// 					EndsAt:   endTime},
// 			},
// 		},
// 		BusinessUnit: monitor.BusinessUnit{Name: monitor.Name(payload.BusinessUnit)},
// 		Product:      monitor.Product{Name: monitor.Name(payload.Product)},
// 	}, nil
// }

type MonitorPayload struct {
	ID string `json:"id" validate:"required"`
	// Brand        brand   `json:"brandName" validate:"required"`
	// BusinessUnit string  `json:"businessUnit" validate:"required"`
	// Product      string  `json:"product" validate:"required"`
	// Flow         string  `json:"flow" validate:"oneof=B2B B2C"`
	// Platform     bool    `json:"platform" validate:"required"`
	TPV float64 `json:"tpv" validate:"required"`
}

type brand struct {
	Name             string        `json:"name" validate:"required"`
	Site             string        `json:"site" validate:"required,url"`
	CustomerHours    customerHours `json:"customerHours" validate:"required"`
	SellerContact    contact       `json:"sellerContact" validate:"required"`
	TechnicalContact contact       `json:"technicalContact" validate:"required"`
}

type customerHours struct {
	SupportHours hours `json:"supportHours" validate:"required"`
	WorkingHours hours `json:"workingHours" validate:"required"`
}

type hours struct {
	Start string `json:"start" validate:"required, datetime=12:00"`
	End   string `json:"end" validate:"required, datetime=12:00"`
}

type contact struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Phone string `json:"phone" validate:"required,phone"`
}
