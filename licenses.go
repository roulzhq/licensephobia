package main

import (
	"reflect"

	"github.com/roulzhq/licensephobia/database"
)

func LicenseExists(license string) bool {
	_, err := DB.GetLicenseNameById(license)
	return err == nil
}

func ConditionsToArray(conditions []database.LicenseConditions) map[string][]bool {
	conditionMap := make(map[string][]bool)

	for _, condition := range conditions {
		c := reflect.ValueOf(&condition)

		t := reflect.TypeOf(condition)
		v := reflect.VisibleFields(t)

		for i, field := range v {
			if field.Name != "LicenseId" {
				value := c.Elem().Field(i)
				conditionMap[field.Name] = append(conditionMap[field.Name], value.Bool())
			}
		}
	}

	return conditionMap
}
