package main

import "fmt"

func media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	if len(numeros) == 0 {
		return 0
	} else {
		return total / float64(len(numeros))
	}
}

func main() {
	fmt.Printf("Média: %.2f", media())
	fmt.Scanln()
}
