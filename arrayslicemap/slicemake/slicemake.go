package main

import "fmt"

func main() {
	s := make([]int, 10) // o slice de tamanho 10 aponta para um array de 10 elementos
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20)        // aqui ele diz que o slice aponta para um array interno de 20 elementos
	fmt.Println(s, len(s), cap(s)) // cap(s) pega o tamanho máximo do slice (array interno)

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s)
}
