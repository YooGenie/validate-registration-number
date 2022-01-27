package registration_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//테스트할 때 유효한 등록번호는 넣어서 테스트하기
func TestValidateRegistrationNumber_외국인인경우Ok(t *testing.T) {
	//given
	regNum := "9901015"
	nationalityType := "FOREIGN"
	//when
	actual := RegistrationNumber(nationalityType, regNum)
	//then
	assert.Equal(t, true, actual)
}

func TestValidateRegistrationNumber_내국인인경우Ok(t *testing.T) {
	//given
	regNum := ""
	nationalityType := "NATIVE"
	//when
	actual := RegistrationNumber(nationalityType, regNum)
	//then
	assert.Equal(t, true, actual)
}

func TestValidateRegistrationNumber_하이픈있는경우(t *testing.T) {
	//given
	regNum := "990101-5"
	nationalityType := "FOREIGN"
	//when
	actual := RegistrationNumber(nationalityType, regNum)
	//then
	assert.Equal(t, true, actual)
}

func TestValidateRegistrationNumber_13자리수아닌경우(t *testing.T) {
	//given
	regNum := "990101512345"
	nationalityType := "FOREIGN"
	//when
	actual := RegistrationNumber(nationalityType, regNum)
	//then
	assert.Equal(t, false, actual)
}

func TestValidateRegistrationNumber_외국인인경우유효하지않는경우(t *testing.T) {
	//given
	regNum := "9901015000000"
	nationalityType := "FOREIGN"
	//when
	actual := RegistrationNumber(nationalityType, regNum)
	//then
	assert.Equal(t, false, actual)
}

func TestValidateRegistrationNumber_내국인인경우유효하지않는경우(t *testing.T) {
	//given
	regNum := ""
	nationalityType := "NATIVE"
	//when
	actual := RegistrationNumber(nationalityType, regNum)
	//then
	assert.Equal(t, false, actual)
}

func TestValidateRegistrationNumber_국적이유효하지않는경우(t *testing.T) {
	//given
	regNum := "9901015"
	nationalityType := "KOREA"
	//when
	actual := RegistrationNumber(nationalityType, regNum)
	//then
	assert.Equal(t, false, actual)
}
