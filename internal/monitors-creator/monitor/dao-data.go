package monitor

import (
	"time"
)

type MonitorDaoData struct {
	ID               int64      `db:"id"`
	Brand            BrandDao   `db:"brand"`
	BusinessUnit     string     `db:"business_unit"`
	ProductName      string     `db:"product_name"`
	Flow             string     `db:"flow"`
	Platform         string     `db:"platform"`
	BrandIdentifiers BrandIDDao `db:"brand_ids"`
	MonthlyTPV       float64    `db:"monthly_tpv"`
	KAM              KAMDao     `db:"KAM"`
	CreatedAt        time.Time  `db:"created_at"`
	UpdatedAt        time.Time  `db:"updated_at"`
	Deleted          bool       `db:"deleted"`
	// Metric    Metric      `db:"metric"`
	// Threshold Threshold   `db:"threshold"`
}

type BrandDao struct {
	Name                    string     `db:"name"`
	Site                    string     `db:"site"`
	CustomerHourStart       time.Time  `db:"customer_hour_start"`
	CustomerHourEnd         time.Time  `db:"customer_hour_end"`
	SellerContact           ContactDao `db:"sellet_contact"`
	TechnicalContactContact ContactDao `db:"technical_contact"`
}

type BrandIDDao struct {
	CustIDs       string `db:"customer_id"`
	AppID         string `db:"app_id"`
	MarketplaceID string `db:"marketplace_id"`
	PlatformID    string `db:"platform_id"`
	BrandID       string `db:"brand_id"`
	SponsorID     string `db:"sponsor_id"`
}

type KAMDao struct {
	AccountManager ContactDao `db:"account_manager"`
}

type ContactDao struct {
	Name  string `db:"name"`
	Email string `db:"email"`
	Phone string `db:"phone"`
}

// func ToDao(dom *Monitor) *MonitorDaoData {
// 	UUID := uuid.New()
// 	ID := MonID{
// 		ID: UUID.String(),
// 	}
// 	return &MonitorDaoData{
// 		ID:               dom.ID,
// 		BrandIdentifiers: dom.Identifiers,
// 		// Metric:    dom.Metric,
// 		// Threshold: dom.Threshold,
// 		// CreatedAt:        CreatedAt,
// 		UpdatedAt: time.Now(),
// 		Deleted:   false,
// 	}
// }

// func ToDomain(dao *MonitorDaoData) *Monitor {
// 	return &Monitor{
// 		Brand: dao.Brand,
// 		// Metric:    dao.Metric,
// 		// Threshold: dao.Threshold,
// 		CreatedAt: dao.CreatedAt,
// 	}
// }
