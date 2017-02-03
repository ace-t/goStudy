package main

import "fmt"

func main() {
	s := make([]string, 4)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "c"
	s[2] = "e"
	s[3] = "t"

	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))
	/*

	   emp: [   ]
	   set: [a c e t]
	   get: e
	   len: 4

	*/

	s = append(s, "b")
	s = append(s, "l", "o", "g")
	fmt.Println("apd:", s)
	// apd: [a c e t b l o g]

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)
	// copy: [a c e t b l o g]
	//        0 1 2 3 4 5 6 7

	l := s[2:5]
	fmt.Println("sl1:", l)
	// sl1: [e t b]

	l = s[:5]
	fmt.Println("sl2:", l)
	// sl2: [a c e t b]

	l = s[3:]
	fmt.Println("sl3:", l)
	// sl3: [t b l o g]

}
