package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type aluno struct {
	Nome   string
	Idade  int
	Bairro string
}

func main() {
	a := aluno{Nome: "Aurélio", Idade: 39, Bairro: "Eugênio de Melo"}
	data, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		fmt.Println("Erro ao converter para json.", err)
		return
	}
	_ = os.WriteFile("data.json", data, 0644)
	fmt.Println("Arquivo criado com sucesso")

}
