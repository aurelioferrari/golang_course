package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	// float e int dá erro - float64(variável) transforma uma variável para o tipo float64
	// int(variável) transforma para int
	fmt.Println(x / float64(y))

	nota := 6.9
	nota_final := int(nota)
	fmt.Println("O tipo da nota é", reflect.TypeOf(nota_final))

	// string(número) - não converte um número para string, mas sim ele identifica o código de uma letra

	// jeito correto de int para str
	fmt.Println("Teste " + strconv.Itoa(123))

	// de string para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("true")
	// Se b for true (ou 1), ene vai entrar no bloco do if
	if b {
		fmt.Println("Verdadeiro")
	}
}
