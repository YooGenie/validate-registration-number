package registration_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateRegistrationNumber_Ok(t *testing.T) {
	//given
	regNum := "9901015020063"
	nationalityType := "FOREIGN"
	//when
	actual := RegistrationNumber(nationalityType, regNum)
	//then
	assert.Equal(t, true, actual)
}


func TestValidateRegistrationNumber_하이픈있는경우(t *testing.T) {
	//given
	regNum := "990101-5020063"
	nationalityType := "FOREIGN"
	//when
	actual := RegistrationNumber(nationalityType, regNum)
	//then
	assert.Equal(t, true, actual)
}
