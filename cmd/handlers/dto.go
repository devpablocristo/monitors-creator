package handlers

import (
	"monitors-creator/internal/monitors-creator/monitor"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterValidation("productvalidation", validateProduct)
}

type BusinessUnit string

const (
	OnlinePayments BusinessUnit = "Online Payments"
	InStore        BusinessUnit = "In Store"
)

type Product string

const (
	CheckoutPRO    Product = "Checkout PRO"
	CheckoutAPI    Product = "Checkout API"
	CheckoutBricks Product = "Checkout Bricks"
	WalletConnect  Product = "Wallet Connect"
	QR             Product = "QR"
	Point          Product = "Point"
	Delivery       Product = "Delivery"
)

type MonitorPayload struct {
	ID                    string        `json:"id,omitempty"`
	BusinessUnit          BusinessUnit  `json:"business_unit" validate:"required,oneof='Online Payments' 'In Store'"`
	Product               Product       `json:"product" validate:"required,productvalidation"`
	BrandName             string        `json:"brand_name" validate:"required"`
	Flow                  string        `json:"flow"`
	Site                  string        `json:"site" validate:"required,oneof=MLA MLB MLC MLM MLU MCO"`
	Description           string        `json:"description"`
	BusinessHours         businessHours `json:"business_hours"`
	Integration           string        `json:"integration"`
	MonthlyTPV            float64       `json:"monthly_tpv"`
	KAM                   string        `json:"kam"`
	SellerContact         string        `json:"seller_contact"`
	TechnicalContact      string        `json:"technical_contact"`
	MarketplaceID         string        `json:"marketplace_id"`
	PlatformID            string        `json:"platform_id"`
	SponsorID             string        `json:"sponsor_id"`
	CollectorID           []string      `json:"collector_id"`
	ApplicationOrClientID string        `json:"application_or_client_id"`
	BrandID               string        `json:"brand_id"`
}

func (m MonitorPayload) Validate() error {
	return validate.Struct(m)
}

// validateProduct validates Product values based on BusinessUnit value
// Business Unit = Online Payments -> Checkout PRO, Checkout API, Checkout Bricks, Wallet Connect
// Business Unit = In Store -> QR, Point, Delivery
func validateProduct(fl validator.FieldLevel) bool {
	data := fl.Parent().Interface().(MonitorPayload)
	switch data.BusinessUnit {
	case OnlinePayments:
		return data.Product == CheckoutPRO || data.Product == CheckoutAPI || data.Product == CheckoutBricks || data.Product == WalletConnect
	case InStore:
		return data.Product == QR || data.Product == Point || data.Product == Delivery
	default:
		return false
	}
}

type businessHours struct {
	Sunday    []hours `json:"sunday"`
	Monday    []hours `json:"monday"`
	Tuesday   []hours `json:"tuesday"`
	Wednesday []hours `json:"wednesday"`
	Thursday  []hours `json:"thursday"`
	Friday    []hours `json:"friday"`
	Saturday  []hours `json:"saturday"`
}

type hours struct {
	Open  string `json:"open"`
	Close string `json:"close"`
}

func MonitorResponse(m monitor.Monitor) *MonitorPayload {
	//TODO FIX THIS SHIT
	return &MonitorPayload{
		ID: "123",
	}
}

func mapPayloadToDomain(payload MonitorPayload) (monitor.Monitor, error) {
	// TODO: CHECAR SE PRECISAMOS CHECAR ERRO EM TIME.PARSE. SE TIME = NIL DÁ ERRO.
	return monitor.Monitor{
		// ID: monitor.ID(payload.ID),
		// MonthlyTPV: monitor.MonthlyTPV{
		// 	USD: monitor.Currency(payload.TPV)},
	}, nil
}
