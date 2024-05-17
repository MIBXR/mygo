package main

import "fmt"

func main_slice() {
	s1 := []int{1, 2, 3, 4}
	s2 := s1[2:]
	fmt.Println("s1: %v, s2: %v", s1, s2)

	s2[0] = 99
	fmt.Println("s1: %v, s2: %v", s1, s2)

	s2 = append(s2, 199)
	fmt.Println("s1: %v, s2: %v", s1, s2)

	s2[1] = 1999
	fmt.Println("s1: %v, s2: %v", s1, s2)
}