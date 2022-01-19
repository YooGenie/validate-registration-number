package registration_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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

func TestValidateResidentRegistrationNumber_Ok(t *testing.T) {
	//given
	a := ""
	//when
	actual := ResidentRegistrationNumber(a)
	//then
	assert.Equal(t, true, actual)
}
