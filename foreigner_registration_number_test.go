package registration_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//테스트할 때 유효한 등록번호는 넣어서 테스트하기
func TestValidateForeignerRegistrationNumber_Ok(t *testing.T) {
	//given
	a := "9901015"
	//when
	actual := ForeignerRegistrationNumber(a)
	//then
	assert.Equal(t, true, actual)
}

func TestValidateForeignerRegistrationNumber_하이픈있는경우(t *testing.T) {
	//given
	a := "990101-5"
	//when
	actual := ForeignerRegistrationNumber(a)
	//then
	assert.Equal(t, true, actual)
}

func TestValidateForeignerRegistrationNumber_13자리아닌경우(t *testing.T) {
	//given
	a := "990101502006"
	//when
	actual := ForeignerRegistrationNumber(a)
	//then
	assert.Equal(t, false, actual)
}

func TestValidateForeignerRegistrationNumber_유효한번호아닌경우(t *testing.T) {
	//given
	a := "9901015123654"
	//when
	actual := ForeignerRegistrationNumber(a)
	//then
	assert.Equal(t, false, actual)
}
