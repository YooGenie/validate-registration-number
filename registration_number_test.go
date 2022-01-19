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
