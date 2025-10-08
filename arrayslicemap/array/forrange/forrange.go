package main

import "fmt"

func main() {
	numeros := [...]float64{1, 2, 3, 4, 5}

	for i, numero := range numeros {
		fmt.Printf("%.1f) %.2f\n", float64(i), numero)
	}

	// para ignorar o índice
	for _, num := range numeros {
		fmt.Println(num)
	}

	for i := 0; i < len(numeros); i++ {
		numeros[i] = 3.0
	}
	fmt.Println(numeros)
}
