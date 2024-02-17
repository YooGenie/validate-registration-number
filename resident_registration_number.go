package registration_number

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

// 주민등록번호
func ResidentRegistrationNumber(regNum string) bool {
	regNum = strings.Replace(regNum, "-", "", 1)
	r, _ := regexp.Compile("[0-9]{2}[0-1][0-9][0-3][0-9][1-4][0-9]{6}")

	regexpCheck := r.MatchString(regNum)

	if !regexpCheck {
		return false
	}

	if len(regNum) != 13 {
		return false
	}

	var arrCheckNum = []int{2, 3, 4, 5, 6, 7, 8, 9, 2, 3, 4, 5}

	nowDate := time.Now().Format("0601")
	nowYear := time.Now().Format("2006")

	sum := 0
	for i := 0; i < 12; i++ {
		r, _ := strconv.Atoi(string(regNum[i]))
		sum += r * arrCheckNum[i]
	}

	lastValue, _ := strconv.Atoi(string(regNum[12]))
	verificationCode := 11 - (sum % 11)

	if regNum[6:7] == "1" || regNum[6:7] == "2" || regNum[6:7] == "3" || regNum[6:7] == "4" {
		seventhDigit := regNum[6:7] == "3" || regNum[6:7] == "4"
		if regNum[0:4] >= "2010" && regNum[0:4] <= nowDate && "20"+regNum[0:2] <= nowYear && (seventhDigit) {
			return true
		}
		if regNum[0:2] >= nowYear && (seventhDigit) {
			return false
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
