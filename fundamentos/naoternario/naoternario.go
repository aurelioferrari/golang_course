package main

import "fmt"

// não tem operador ternário em go

// função que recebe uma nota e retorna dizendo se o aluno está aprovado ou não
func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"

}

func main() {
	fmt.Println(obterResultado(7.1))
}
