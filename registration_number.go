package registration_number

import "strings"

func RegistrationNumber(nationalityType string, regNum string) bool {
	regNum = strings.Replace(regNum, "-", "", 1)

	if nationalityType == "NATIVE" {
		return ResidentRegistrationNumber(regNum)
	} else if nationalityType == "FOREIGN" {
		return ForeignerRegistrationNumber(regNum)
	} else {
		return false
	}
}
