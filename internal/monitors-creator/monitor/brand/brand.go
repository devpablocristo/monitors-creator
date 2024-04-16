package brand

import (
	brandid "monitors-creator/internal/monitors-factory/monitor/brand/brand-id"
)

type Name struct {
	BrandName string
}

type BusinessUnit struct {
	UnitName string
}

type Product struct {
	ProductName string
}

type Flow struct {
	TransactionFlow string
}

type Site struct {
	WebsiteURL string
}

type CustomerHours struct {
	SupportHours string
}

type Platform struct {
	IsPlatform bool
}

type MonthlyTPV struct {
	TransactionVolume float64
}

type KAM struct {
	AccountManager string
}

type Contact struct {
	EmailAddress string
}

type Brand struct {
	Name          Name
	BusinessUnit  BusinessUnit
	Product       Product
	Flow          Flow
	Site          Site
	CustomerHours CustomerHours
	Platform      Platform
	Identifiers   brandid.BrandIdentifiers
	MonthlyTPV    MonthlyTPV
	KAM           KAM
	SellerContact Contact
	TechContact   Contact
}
