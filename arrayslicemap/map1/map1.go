package main

import "fmt"

func main() {
	// var aprovados map[int]string // map[chave]valor
	// maps devem ser inicializados
	aprovados := make(map[int]string)
	aprovados[12345678978] = "Maria"
	aprovados[45678912345] = "Pedro"
	aprovados[78945612312] = "Ana"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("Nome: %s - CPF: %d\n", nome, cpf)
	}

	// deletando um usuario
	delete(aprovados, 12345678978)
	if len(aprovados[12345678978]) > 0 {
		fmt.Println("Usuário ainda cadastrado")
	} else {
		fmt.Println("Usuário removido com sucesso")
	}
}
