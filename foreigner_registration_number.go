package registration_number

import (
	"strconv"
	"strings"
	"time"
)

func ForeignerRegistrationNumber(regNum string) bool {
	regNum = strings.Replace(regNum, "-", "", 1)

	var arrCheckNum = []int{2, 3, 4, 5, 6, 7, 8, 9, 2, 3, 4, 5}

	if len(regNum) != 13 {
		return false
	}

	nowDate := time.Now().Format("0601")

	sum := 0
	for i := 0; i < 12; i++ {
		r, _ := strconv.Atoi(string(regNum[i]))
		sum += r * arrCheckNum[i]
	}

	lastValue, _ := strconv.Atoi(string(regNum[12]))
	verificationCode := 13 - (sum % 11)

	if regNum[6:7] == "5" || regNum[6:7] == "6" || regNum[6:7] == "7" || regNum[6:7] == "8" {
		seventhDigit := regNum[6:7] == "7" || regNum[6:7] == "8"
		if regNum[0:4] >= "2010" && regNum[0:4] <= nowDate && (seventhDigit) {
			return true
		}
		if verificationCode >= 10 && verificationCode <= 13 {
			verificationCode = verificationCode - 10
		}
		if lastValue == verificationCode {
			return true
		}
	}
	return false

}
