package main
import "fmt"
import "math"

const s string = "acet-t"

/*
숫자 상수는 명시적 캐스팅등으로 타입이 주어지기 전까진 타입을 가지지 않습니다.
*/
func constant_test() {
	fmt.Println(s)
	
	const n = 500000
	const d = 3e20 / n

	fmt.Println(d)    	// 6e+14
	fmt.Println(int64(d)) 	// 600000000000000

	fmt.Println(math.Sin(n))

}


