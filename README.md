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
nationalityType 종류 : NATIVE, FOREIGN


# 사용법
```
	registration_number.ForeignerRegistrationNumber(regNum)
	registration_number.ResidentRegistrationNumber(regNum)
	registration_number.RegistrationNumber(nationalityType,regNum)
```

# UPDATE
2022-01-23 v0.0.1 #4 주민번호과 외국인번호에 하이픈있으면 제거하기
