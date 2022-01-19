package registration_number

func RegistrationNumber(nationalityType string, regNum string) bool {
	if nationalityType == "NATIVE" {
		return ResidentRegistrationNumber(regNum)
	} else if nationalityType == "FOREIGN" {
		return ForeignerRegistrationNumber(regNum)
	} else {
		return false
	}
}
