package main
import "fmt"
func main() {
	var a string = "ace-t"
	fmt.Println(a)
	
	var b, c int = 1, 2
   	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int      // 초기화 없이 선언된 변수는 default는 0 
	fmt.Println(e)

	f := "short"
 	fmt.Println(f) 
/*
	:= 문법은 변수를 선언하는 동시에 초기화하기 위한 단축 표현식입니다. 
	예를 들면 이 경우는 var f string = "short"를 뜻합니다.
*/

}



