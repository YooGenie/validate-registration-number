package registration_number

import (
	"strconv"
	"strings"
	"time"
)

// 주민등록번호
func ResidentRegistrationNumber(regNum string) bool {
	regNum = strings.Replace(regNum, "-", "", 1)

	if len(regNum) != 13 {
		return false
	}

	var arrCheckNum = []int{2, 3, 4, 5, 6, 7, 8, 9, 2, 3, 4, 5}

	nowDate := time.Now().Format("0601")

	sum := 0
	for i := 0; i < 12; i++ {
		r, _ := strconv.Atoi(string(regNum[i]))
		sum += r * arrCheckNum[i]
	}

	lastValue, _ := strconv.Atoi(string(regNum[12]))
	verificationCode := 11 - (sum % 11)

	if regNum[6:7] == "1" || regNum[6:7] == "2" || regNum[6:7] == "3" || regNum[6:7] == "4" {
		seventhDigit := regNum[6:7] == "3" || regNum[6:7] == "4"
		if regNum[0:4] >= "2010" && regNum[0:4] <= nowDate && (seventhDigit) {
			return true
		}
		if verificationCode >= 10 && verificationCode <= 11 {
			verificationCode = verificationCode - 10
		}
		if lastValue == verificationCode {
			return true
		}
	}
	return false
}
