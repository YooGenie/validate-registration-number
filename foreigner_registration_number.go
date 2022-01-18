package registration_number

import (
	"fmt"
	"strconv"
)

func ForeignerRegistrationNumber(regNum string) bool {

	var arrCheckNum = []int{2, 3, 4, 5, 6, 7, 8, 9, 2, 3, 4, 5}

	sum := 0
	if len(regNum) == 13 {
		for i := 0; i < 12; i++ {
			r, _ := strconv.Atoi(string(regNum[i]))
			sum += r * arrCheckNum[i]
		}

		lastValue, _ := strconv.Atoi(string(regNum[12]))
		fmt.Println(lastValue)
		seventh, _ := strconv.Atoi(string(regNum[6]))
		verificationCode := 13 - (sum % 11)

		if verificationCode >= 10 && verificationCode <= 13 {
			verificationCode = verificationCode - 10
		}

		fmt.Println(verificationCode)
		if lastValue == verificationCode && (seventh == 5 || seventh == 6 || seventh == 7 || seventh == 8) {
			return true
		} else {
			fmt.Println("외국인등록번호를 제대로 입력해주세요")
			return false
		}
	} else {
		fmt.Println("외국인등록번호를 제대로 입력해주세요")
		return false
	}

}
