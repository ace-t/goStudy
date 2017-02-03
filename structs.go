package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
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

	sp := &s
	fmt.Println(sp.age) // 50

	sp.age = 100
	fmt.Println(sp.age) // 100

}
