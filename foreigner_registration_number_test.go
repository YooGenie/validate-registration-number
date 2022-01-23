package registration_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateForeignerRegistrationNumber_Ok(t *testing.T) {
	//given
	a := "9901015020063"
	//when
	actual := ForeignerRegistrationNumber(a)
	//then
	assert.Equal(t, true, actual)
}

func TestValidateForeignerRegistrationNumber_하이픈있는경우(t *testing.T) {
	//given
	a := "990101-5020063"
	//when
	actual := ForeignerRegistrationNumber(a)
	//then
	assert.Equal(t, true, actual)
}