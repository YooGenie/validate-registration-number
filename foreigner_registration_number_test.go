package registration_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateForeignerRegistrationNumber_Ok(t *testing.T) {
	//given
	a := "9901015020063"
	//when
	actual :=  ForeignerRegistrationNumber(a)
	//then
	assert.Equal(t, true, actual)
}

