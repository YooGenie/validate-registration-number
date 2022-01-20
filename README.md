# validate-registration-number
주민등록번호와 외국인번호 유효성 검사하는 모듈

# Installation
```
go get -u github.com/YooGenie/validate-registration-number
```

# Description

* 주민등록번호 입력만 해서 유효성 검사(regNum 숫자 13자리 입력)

```
func ResidentRegistrationNumber(regNum string) bool {
    return bool
	}
```

* 외국인등록번호 입력만 해서 유효성 검사(regNum 숫자 13자리 입력)
```
func ForeignerRegistrationNumber(regNum string) bool {
    return bool
}
```

* 국적타입과 등록번호 입력하면 외국인등록번호와 주민등록번호 유효성 검사(regNum 숫자 13자리 입력)
```
func RegistrationNumber(nationalityType string, regNum string) bool {
    return bool
}
```
