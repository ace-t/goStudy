package main

import "fmt"

// fact 함수는 베이스 케이스인 fact(0)에 도달할 때까지 자기자신을 호출합니다.
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
func recursive_test() {
	fmt.Println(fact(7))
}
