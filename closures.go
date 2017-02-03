package main

import "fmt"

// intSeq 함수는 내부에 익명으로 정의한 또 다른 함수를 반환합니다.
// 반환된 함수는 클로저를 만들기 위해 i 변수를 가둬둡니다(closes over).
func intSeq() func() int { // 형태가 매우 이상하네.. 지금까지는 func intSeq(a,b int) int {} 등의 형태가 되겠다.
	i := 0
	return func() int {
		i += 1
		return i
	}
}
func main() {
	// intSeq를 호출하여, 결괏값(함수)을 nextInt에 할당합니다. 이 함숫값은 nextInt를 호출할 때마다 업데이트 되는 i 값을 캡쳐합니다.
	nextInt := intSeq()
	// netxtInt를 몇 번 호출하면서 클로저가 어떻게 동작하는지 봅시다.
	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3
	// 특정 함수에서 상태값이 유일한지 확인하기 위해, 하나를 새롭게 생성하고 테스트 해봅니다.
	newInts := intSeq()
	fmt.Println(newInts()) // 1
}
