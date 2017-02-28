package main

import "fmt"

type person struct {
	name string
	age  int
}

func struct_test() {
	fmt.Println(person{"AceT", 17})
	th := person{"hahaha", 200}
	fmt.Println(th.age)
	fmt.Println(person{name: "AceT2", age: 27})
	fmt.Println(person{name: "Mae"})
	fmt.Println(&person{name: "Sun", age: 40})

	s := person{name: "Moon", age: 50}
	fmt.Println(s.name)

	/*  result :
	{AceT 17}
	200
	{AceT2 27}
	{Mae 0}
	&{Sun 40}
	Moon
	*/

	sp := &s            // 구조체 포인터 &s 로 주소값을 주게 되면 동일하게 sp의 값을 변경하면 s도 변경이 된다. 주소로 참조 되니깐~
	fmt.Println(sp.age) // 50

	sp.age = 100
	fmt.Println(sp.age) // 100
	fmt.Println(s.age)  // 100

}
