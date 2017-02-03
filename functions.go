package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int { // 형태를 주시해보자!  변수명 나오고 뒤에 타입! 그리고 리턴타입!
	return a + b + c
}

func myFunc(a, b int) string {
	returnVar := fmt.Sprintf("%s %d %d", "ace-t", a, b)
	return returnVar
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	res2 := myFunc(1, 7)
	fmt.Println("myFunc =", res2)
}
