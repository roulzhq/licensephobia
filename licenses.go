package main

func LicenseExists(license string) bool {
	_, err := DB.GetLicenseNameById(license)

	return err == nil
}
