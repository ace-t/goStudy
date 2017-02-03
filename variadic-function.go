package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)
	// 만약 값들이 이미 슬라이스에 들어 있다면, 다음과 같이 func(slice...)를 사용하여 가변 함수의 인자에 적용하세요.
	nums := []int{1, 2, 3, 4}
	sum(nums...)

	numSlice := make([]int, 3)
	numSlice = append(numSlice, 1, 2, 3)
	fmt.Println("numSlice:", numSlice)
	sum(numSlice...)

	/*
		[1 2] 3
		[1 2 3] 6
		[1 2 3 4] 10
		numSlice: [0 0 0 1 2 3]
		[0 0 0 1 2 3] 6
	*/
}
