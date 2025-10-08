package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"José João":     11325.45,
		"Gabriel Silva": 15456.78,
		"Pedro Junioe":  1200.0,
	}

	funcsESalarios["Rafael Filho"] = 1350.0 // adiciona novo valor ao map
	delete(funcsESalarios, "Inexistente")   //não gera erro ao tentar excluir um elemento que não existe

	for nome, salario := range funcsESalarios {
		fmt.Printf("Nome: %s\nSalário: R$%.2f\n", nome, salario)
	}
}
