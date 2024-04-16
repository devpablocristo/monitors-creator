package brandid

// CustID representa los identificadores de cliente para una marca.
type CustID struct {
	IDs []string
}

// AppID representa el identificador único de una aplicación específica asociada a una marca.
type AppID struct {
	ID string
}

// MarketplaceID representa el identificador único de una marca dentro de un marketplace.
type MarketplaceID struct {
	ID string
}

// PlatformID representa el identificador único de la plataforma de servicios donde opera la marca.
type PlatformID struct {
	ID string
}

// BrandID representa el identificador único de una marca.
type BrandID struct {
	ID string
}

// SponsorID representa el identificador único de un patrocinador asociado a la marca.
type SponsorID struct {
	ID string
}

// BrandIdentifiers agrupa todos los posibles identificadores asociados a una marca.
type BrandIdentifiers struct {
	CustIDs       CustID
	AppID         AppID
	MarketplaceID MarketplaceID
	PlatformID    PlatformID
	BrandID       BrandID
	SponsorID     SponsorID
}
