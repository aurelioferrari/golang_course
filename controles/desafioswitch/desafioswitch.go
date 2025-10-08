package main

import "fmt"

func notaParaConceito(n float64) {
	switch {
	case n < 0:
		fmt.Println("Nota inválida")
	case n < 3:
		fmt.Println("Aluno tipo E")
	case n < 5:
		fmt.Println("Aluno tipo D")
	case n < 8:
		fmt.Println("Aluno tipo C")
	case n < 9:
		fmt.Println("Aluno tipo B")
	case n <= 10:
		fmt.Println("Aluno tipo A")
	default:
		fmt.Println("Nota inválida")
	}
}

func main() {
	notaParaConceito(1)
	notaParaConceito(22)
	notaParaConceito(6)
	notaParaConceito(8.9)
	notaParaConceito(-1)
}
