package main

import "fmt"

func maps_test(){
	m := make(map[string]int)  // map
	t := make([]string,4)     // slice

	m["t1"] = 7
	m["t2"] = 14
	fmt.Println("map:", m)
	fmt.Println("slice:", t)

	v1 := m["t1"]
	fmt.Println("v1 :", v1)
/*
map: map[t1:7 t2:14]
slice: [   ]
v1 : 7
*/

	fmt.Println("len:", len(m))
	delete(m, "t2")     // 삭제!
	fmt.Println("map:", m)
/*
len: 2
map: map[t1:7]
*/

	_, prs := m["t2"]          // 값 자체가 필요없는 경우엔, 공백 식별자(blank identifier) _를 사용해 무시할 수 있습니다.
	fmt.Println("prs:", prs)  
// prs: false

	n := map[string]int{"foo": 1, "bar":2}   // 한줄로 맵선업 및 초기화 
	fmt.Println("map:", n)
// map: map[foo:1 bar:2]	

}
