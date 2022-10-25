package main

import "fmt"

func main() {
	s := struct1{}
	fmt.Println(s.a)
	s.inc()
	fmt.Println(s.a)
	s.a = dec(s.a)
	fmt.Println(s.a)
}
