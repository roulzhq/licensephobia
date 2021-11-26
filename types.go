package main

type License struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Conditions LicenseCondition
}

type LicenseCondition map[string][]bool
