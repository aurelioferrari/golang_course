package main

import "fmt"

func notaParaConceito(n float64) string {
	var nota = int(n)
	switch nota {
	case 10:
		fallthrough // faz com que vá para o próximo case
	case 9:
		return "A"
	case 8, 7: // se a nota for 8 OU 7
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota inválida" // caso não entre em nenhum dos casos
	}
}

func main() {
	fmt.Println(notaParaConceito(1))
	fmt.Println(notaParaConceito(1.1))
	fmt.Println(notaParaConceito(8))
	fmt.Println(notaParaConceito(12))
}
