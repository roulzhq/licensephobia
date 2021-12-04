package main

type License struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Conditions LicenseCondition
}

type LicenseCondition map[string][]bool

type LicenseInfo struct {
	Found   bool   `json:"found"`
	Known   bool   `json:"known"`
	License string `json:"type"`
}

type PackageManger string

const (
	Npm   PackageManger = "npm"
	Pip   PackageManger = "pip"
	Cargo PackageManger = "cargo"
)

type Package struct {
	Id            string      `json:"id"`
	Name          string      `json:"name"`
	Description   string      `json:"description"`
	LatestVersion string      `json:"latestVersion"`
	License       LicenseInfo `json:"license"`
	Author        string      `json:"author"`
	Homepage      string      `json:"homepage"`
}

type Summary struct {
	Conditions SummaryConditions `json:"conditions"`
}

type SummaryConditions struct {
	Conditions  []string `json:"conditions"`
	Permissions []string `json:"permissions"`
	Limitations []string `json:"limitations"`
}
