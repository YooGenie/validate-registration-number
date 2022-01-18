package registration_number

import (
	"fmt"
	"strconv"
)

//주민등록번호
func RegistrationNumber(regNum string) bool {

	var arrCheckNum = []int{2, 3, 4, 5, 6, 7, 8, 9, 2, 3, 4, 5}

	sum := 0
	if len(regNum) == 13 {
		for i := 0; i < 12; i++ {
			r, _ := strconv.Atoi(string(regNum[i]))
			sum += r * arrCheckNum[i]
		}

		lastValue, _ := strconv.Atoi(string(regNum[12]))
		verificationCode := 11 - (sum % 11)

		if verificationCode >= 10 && verificationCode <= 11 {
			verificationCode = verificationCode - 10
		}

		if lastValue == verificationCode {
			return true
		} else {
			fmt.Println("주민등록번호를 제대로 입력해주세요")
			return false
		}
	} else {
		fmt.Println("주민등록번호를 제대로 입력해주세요")
		return false
	}

}
