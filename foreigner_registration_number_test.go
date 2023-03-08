package registration_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 테스트할 때 유효한 등록번호는 넣어서 테스트하기
func TestValidateForeignerRegistrationNumber_Ok(t *testing.T) {
	//given
	a := "990101"
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
	a := "9901015"
	//when
	actual := ForeignerRegistrationNumber(a)
	//then
	assert.Equal(t, false, actual)
}

func TestValidateForeignerRegistrationNumber_유효한번호아닌경우(t *testing.T) {
	//given
	a := "9901015"
	//when
	actual := ForeignerRegistrationNumber(a)
	//then
	assert.Equal(t, false, actual)
}

func TestValidateForeignerRegistrationNumber_2020년10월이후_OK(t *testing.T) {
	//given
	a := "2010017000000"
	//when
	actual := ForeignerRegistrationNumber(a)
	//then
	assert.Equal(t, true, actual)
}

func TestValidateForeignerRegistrationNumber_2020년10월이후_NO(t *testing.T) {
	//given
	a := "2010015000000"
	//when
	actual := ForeignerRegistrationNumber(a)
	//then
	assert.Equal(t, false, actual)
}

func TestValidateBusinessNumber_외국인_2시작(t *testing.T) {
	//given
	a := "8807242000000"

	//when
	actual := ForeignerRegistrationNumber(a)
	//then
	assert.Equal(t, false, actual)
	assert.Equal(t, len(a), 13)
}

func TestValidateBusinessNumber_1930_일곱번째_7로_시작(t *testing.T) {
	//given
	a := "3007247000000"

	//when
	actual := ForeignerRegistrationNumber(a)
	//then
	assert.Equal(t, false, actual)
	assert.Equal(t, len(a), 13)
}

func TestValidateBusinessNumber_앞자리오류(t *testing.T) {
	//given
	a := "2323213213231"

	//when
	actual := ForeignerRegistrationNumber(a)
	//then
	assert.Equal(t, false, actual)
	assert.Equal(t, len(a), 13)
}

func TestValidateBusinessNumber_한글오류(t *testing.T) {
	//given
	a := "9008013sksksk"

	//when
	actual := ForeignerRegistrationNumber(a)
	//then
	assert.Equal(t, false, actual)
	assert.Equal(t, len(a), 13)
}
