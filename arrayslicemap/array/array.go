package main

import "fmt"

func main() {
	// são homogêneas e estática - parecido com tupla em python
	// se tiver inteiros, só terá esse tipo de variável nele
	var notas [3]float64 // nome [tamanho]tipo - por padrão é criado tudo com zero
	fmt.Println(notas)
	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.1
	fmt.Println(notas)
	var total float64 = 0
	for i := 0; i < len(notas); i++ {

		total += notas[i]
	}
	fmt.Printf("Média: %.2f", total/float64(len(notas)))
}
