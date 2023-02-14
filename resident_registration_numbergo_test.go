package registration_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 테스트할 때 유효한 등록번호는 넣어서 테스트하기
func TestValidateResidentRegistrationNumber_Ok(t *testing.T) {
	//given
	a := ""
	//when
	actual := ResidentRegistrationNumber(a)
	//then
	assert.Equal(t, true, actual)
}

func TestValidateResidentRegistrationNumber_하이픈있는경우(t *testing.T) {
	//given
	a := ""
	//when
	actual := ResidentRegistrationNumber(a)
	//then
	assert.Equal(t, true, actual)
}

func TestValidateResidentRegistrationNumber_13자리아닌경우(t *testing.T) {
	//given
	a := "123456789011"
	//when
	actual := ResidentRegistrationNumber(a)
	//then
	assert.Equal(t, false, actual)
	assert.Equal(t, len(a), 12)

}

func TestValidateResidentRegistrationNumber_유효한번호아닌경우(t *testing.T) {
	//given
	a := "0001014000000"
	//when
	actual := ResidentRegistrationNumber(a)
	//then
	assert.Equal(t, false, actual)
}

func TestValidateResidentRegistrationNumber_2020년10월이후_OK(t *testing.T) {
	//given
	a := "2010014000000"
	//when
	actual := ResidentRegistrationNumber(a)
	//then
	assert.Equal(t, true, actual)
}

func TestValidateResidentRegistrationNumber_2020년10월이후_NO(t *testing.T) {
	//given
	a := "2010012000000"
	//when
	actual := ResidentRegistrationNumber(a)
	//then
	assert.Equal(t, false, actual)
}

func TestValidateResidentRegistrationNumber_내국인_5로_시작(t *testing.T) {
	//given
	a := "9101015020063"
	//when
	actual := ResidentRegistrationNumber(a)
	//then
	assert.Equal(t, false, actual)
}

func TestValidateResidentRegistrationNumber_내국인_1930년_7번째자리_3으로_했을때(t *testing.T) {
	//given
	a := "3007233000000"
	//when
	actual := ResidentRegistrationNumber(a)
	//then
	assert.Equal(t, false, actual)
}
