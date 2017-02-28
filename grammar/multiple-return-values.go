package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func multiple_return_test() {
	a, b := vals()
	fmt.Println(a) // 3
	fmt.Println(b) // 7

	// 반환값들중 일부만 사용하고 싶을땐, 공백 식별자 _을 사용합니다.
	_, c := vals()
	fmt.Println(c) // 7

}
