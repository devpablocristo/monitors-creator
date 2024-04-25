package monitor

import (
	"net/mail"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type Monitor struct {
	ID                    ID
	BusinessUnit          BusinessUnit
	Product               Product
	BrandName             Name
	Flow                  Flow
	Site                  Site
	Descripion            Description
	BusinessHours         BusinessHours
	Integration           Integration
	MonthlyTPV            Currency
	KAM                   Email
	SellerContact         Email
	TechnicalContact      Email
	MarketplaceID         ID
	PlatformID            ID
	SponsorID             ID
	CollectorID           []ID
	ApplicationOrClientID ID
	BrandID               ID
}

type BusinessHours struct {
	Sunday    []Hours
	Monday    []Hours
	Tuesday   []Hours
	Wednesday []Hours
	Thursday  []Hours
	Friday    []Hours
	Saturday  []Hours
}

type Hours struct {
	Open  time.Time
	Close time.Time
}

type BusinessUnit string

func (bu BusinessUnit) String() string {
	return string(bu)
}

type Integration string

func (i Integration) String() string {
	return string(i)
}

type Site string

func (s Site) String() string {
	return string(s)
}

type Product string

func (p Product) String() string {
	return string(p)
}

type Name string

func (n Name) String() string {
	return string(n)
}

type Description string

func (d Description) String() string {
	return string(d)
}

func (d Description) Validate() bool {
	return len(d) < 200 &&
		len(d) > 10
}

type ID string

func (i ID) String() string {
	return string(i)
}

func (i *ID) Create() ID {
	return ID(uuid.New().String())
}

type Enabled bool

func (e Enabled) String() string {
	return strconv.FormatBool(bool(e))
}

type Email string

func (e Email) String() string {
	return string(e)
}

func (e Email) Validate() bool {
	_, err := mail.ParseAddress(e.String())
	return err == nil
}

type Flow string

func (f Flow) String() string {
	return string(f)
}

type Currency float64

func (c Currency) String() string {
	return strconv.FormatFloat(float64(c), 'f', 2, 64)
}
