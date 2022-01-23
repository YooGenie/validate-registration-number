package registration_number

import (
	"fmt"
	"strconv"
	"strings"
)

//주민등록번호
func ResidentRegistrationNumber(regNum string) bool {
	regNum = strings.Replace(regNum, "-", "", 1)

	var arrCheckNum = []int{2, 3, 4, 5, 6, 7, 8, 9, 2, 3, 4, 5}

	sum := 0
	if len(regNum) == 13 {
		for i := 0; i < 12; i++ {
			r, _ := strconv.Atoi(string(regNum[i]))
			sum += r * arrCheckNum[i]
		}

		lastValue, _ := strconv.Atoi(string(regNum[12]))
		seventh, _ := strconv.Atoi(string(regNum[6]))
		verificationCode := 11 - (sum % 11)

		if verificationCode >= 10 && verificationCode <= 11 {
			verificationCode = verificationCode - 10
		}

		if lastValue == verificationCode && (seventh == 1 || seventh == 2 || seventh == 3 || seventh == 4) {
			return true
		} else {
			fmt.Println("주민등록번호를 제대로 입력해주세요")
			return false
		}
	} else {
		fmt.Println("주민등록번호를 13자리를 입력하세요")
		return false
	}

}
