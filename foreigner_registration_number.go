package registration_number

import (
	"strconv"
	"strings"
)

func ForeignerRegistrationNumber(regNum string) bool {
	regNum = strings.Replace(regNum, "-", "", 1)

	var arrCheckNum = []int{2, 3, 4, 5, 6, 7, 8, 9, 2, 3, 4, 5}

	if len(regNum) != 13 {
		return false
	}

	seventhDigit := regNum[6:7] == "7" || regNum[6:7] == "8"
	if regNum[0:4] >= "2010" && seventhDigit {
		return true
	}

	sum := 0
	if len(regNum) == 13 {
		for i := 0; i < 12; i++ {
			r, _ := strconv.Atoi(string(regNum[i]))
			sum += r * arrCheckNum[i]
		}

		lastValue, _ := strconv.Atoi(string(regNum[12]))
		seventh, _ := strconv.Atoi(string(regNum[6]))
		verificationCode := 13 - (sum % 11)

		if verificationCode >= 10 && verificationCode <= 13 {
			verificationCode = verificationCode - 10
		}

		if lastValue == verificationCode && (seventh == 5 || seventh == 6 || seventh == 7 || seventh == 8) {
			return true
		} else {
			return false
		}
	}
	return false
}
