package monitor

import (
	"net/mail"

	"github.com/corthmann/go-time-intervals/timeinterval"
	"github.com/google/uuid"

	"monitors-creator/internal/monitors-creator/monitor/brandid"
)

type Monitor struct {
	ID               ID
	Brand            Brand
	BusinessUnit     BusinessUnit
	Product          Product
	Flow             Flow
	Platform         Platform
	BrandIdentifiers brandid.BrandIdentifiers
	MonthlyTPV       MonthlyTPV
	KAM              KAM
}

type Name string

func (n Name) String() string {
	return string(n)
}

func (n Name) Validate() bool {
	return len(n) < 50 &&
		len(n) > 3
}

func (n Name) isEmpty() bool {
	return n.String() == ""
}

type ID string

func (i ID) String() string {
	return string(i)
}

func (i *ID) Create() string {
	return uuid.New().String()
}

type Enabled bool

type Email string

func (e Email) String() string {
	return string(e)
}

func (e Email) Validate() bool {
	_, err := mail.ParseAddress(e.String())
	return err == nil
}

type Phone string

type URL string

type TransactionFlow string

type Platform string

type Currency float64

type Brand struct {
	Name             Name
	Site             URL
	CustomerHours    CustomerHours
	SellerContact    Contact
	TechnicalContact Contact
}

type BusinessUnit struct {
	Name Name
}

type Product struct {
	Name Name
}

type Flow struct {
	TransactionFlow TransactionFlow
}

type CustomerHours struct {
	SupportHours timeinterval.Interval
	WorkingHours timeinterval.Interval
}

type MonthlyTPV struct {
	USD Currency
}

type KAM struct {
	AccountManager Contact
}

type Contact struct {
	Name  Name
	Email Email
	Phone Phone
}
