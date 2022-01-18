package registration_number

import (
"github.com/stretchr/testify/assert"
"testing"
)


func TestValidateRegistrationNumber_13자리아닌경우(t *testing.T) {
	//given
	a := "123456789"
	//when
	actual :=  RegistrationNumber(a)
	//then
	assert.Equal(t, false, actual)
	assert.Equal(t, len(a), 9)
}


func TestValidateRegistrationNumber_유효한번호아닌경우(t *testing.T) {
	//given
	a := "1234567890"
	//when
	actual :=  RegistrationNumber(a)
	//then
	assert.Equal(t, false, actual)
}

func TestValidateRegistrationNumber_Ok(t *testing.T) {
	//given
	a := "9901015020063"
	//when
	actual :=  RegistrationNumber(a)
	//then
	assert.Equal(t, true, actual)
}

