package database

type License struct {
	Id                    string   `json:"licenseId"`
	Name                  string   `json:"licenseId"`
	IsDeprecatedLicenseId bool     `json:"isDeprecatedLicenseId"`
	ReferenceNumber       int      `json:"referenceNumber"`
	Reference             string   `json:"reference"`
	DetailsUrl            string   `json:"detailsUrl"`
	IsFsfLibre            string   `json:"isFsfLibre"`
	IsOsiApproved         string   `json:"isOsiApproved"`
	SeeAlso               []string `json:"seeAlso"`
}

type LicenseConditions struct {
	LicenseId                     string `json:"licenseId`
	CommercialUse                 bool   `json:"commercialUse`
	Distribution                  bool   `json:"distribution`
	Modification                  bool   `json:"modification`
	PatentUse                     bool   `json:"patentUse`
	PrivateUse                    bool   `json:"privateUse`
	DiscloseSource                bool   `json:"discloseSource`
	LicenseAndCopyrightNotice     bool   `json:"licenseAndCopyrightNotice`
	LicenseAndCopyrightNoBinaries bool   `json:"licenseAndCopyrightNoBinaries`
	NetworkUseIsDistribution      bool   `json:"networkUseIsDistribution`
	SameLicense                   bool   `json:"sameLicense`
	StateChanges                  bool   `json:"stateChanges`
	Liability                     bool   `json:"liability`
	TrademarkUse                  bool   `json:"trademarkUse`
	Warranty                      bool   `json:"warranty`
}
