package main

import (
	"fmt"
)

type Service interface {
	mywork(value int)
}

func (s *service) mywork(value int) int {
	return value
}

type service struct {
	value  int
	second float64
}

func main() {

	s := service{22, 22.0}
	value := s.mywork(20000)

	fmt.Println(value)
	//fmt.Println(&s)
}
