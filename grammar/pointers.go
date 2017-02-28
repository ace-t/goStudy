package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}
func zeroptr(iptr *int) {
	fmt.Println("i:=1의 iptr의 포인터 값 : ", *iptr)
	*iptr = 0
	fmt.Println("*iptr=0 이후 의 포인터 값 : ", *iptr)

}
func pointer_test() {
	i := 1
	fmt.Println("initial:", i)
	zeroval(i)
	fmt.Println("zeroval:", i)
	// &i는 i의 메모리 주소를 반환합니다.즉, i의 포인터입니다.
	zeroptr(&i) //
	fmt.Println("zeroptr:", i)
	// 포인터도 출력할 수 있습니다.
	fmt.Println("pointer:", &i)
}

/*
initial: 1
zeroval: 1
i:=1의 iptr의 포인터 값 :  1
*iptr=0 이후 의 포인터 값 :  0
zeroptr: 0
pointer: 0xc42000a118
*/
