package main

import "fmt"

func main() {
	a1 := [3]int{1, 2, 3} //array
	s1 := []int{1, 2, 3}  //slice
	fmt.Println(a1, s1)

	a2 := [5]int{1, 2, 3, 4, 5}
	// gerando slices a partir desse array
	// slice define um pedaço CONTÍNUO de um array
	s2 := a2[1:3] // inclui o 1 e nao inclui o 3
	fmt.Println(a2, s2)
	s3 := a2[:2]
	fmt.Println(s3)
	fmt.Println(s3[0])

	// também é possível fazer slice do slice
	s4 := s2[:1]
	fmt.Println(s4)
}
