# goStudy
## go lang을 해보자!

# go cmd
## go build : 빌드 
## go run   : 실행 
## go vet   : vet 도구가 자동으로 에러 검출
### printf 스타일의 함수 호출 시 잘못된 매개변수 지정.
### 메서드 정의 시 시그너처(signature) 관련 에러 
### 잘못 구성된 태그(tag)
### 조합 리터럴(composite literal) 사용 시 누락된 키(key)
## go fmt : 소스코드에 자동으로 format을 맞춰준다. 

# 슬라이스의 내부 구조 및 원리 
## 슬라이스는 데이터의 컬렉션을 처리할 수 있는 방법을 제공하는 데이터 구조
## 동적 배열의 개념으로 구현. -> 컬렉션의 크기를 늘리거나 줄일 수 있음. 
## 인덱싱, 반복 및 가비지 컬렉션 최적화 등의 장점 제공. 
# 슬라이스 생성 
## make 내장 함수를 사용하는 방법.make 함수를 호출할 때 슬라이스의 길이를 명시. 
## ex) slice := make([]string, 5) 
## 슬라이스 리터럴을 이용한 슬라이스 생성 
## slice := []string{"RED", "BLUE"}
## slice := []int{10,20,30}
## m := make(map[string]int)  // map
## t := make([]string,4)     // slice

# 아직은 어색한 
## func plusPlus(a, b, c int) int 
### func 함수명(변수1, 변수2, 변수3 타입) 리턴타입 
