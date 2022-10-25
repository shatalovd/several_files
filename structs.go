package main

type struct1 struct {
	a int
}

func (s *struct1) inc() {
	s.a++
}
