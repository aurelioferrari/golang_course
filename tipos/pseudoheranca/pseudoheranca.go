package main

import (
	"fmt"
)

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	carro
	turboLigado bool
}

func (f *ferrari) mudaNome(novo_nome string) {
	f.nome = novo_nome
}

func main() {
	f1 := ferrari{}
	f1.nome = "F40"
	f1.turboLigado = true
	fmt.Printf("A Ferrari %s está com o turbo ligado? %v\n", f1.nome, f1.turboLigado)
	fmt.Printf("Velocidade atual: %vkm/h\n", f1.velocidadeAtual)

	f1.mudaNome("F15")
	fmt.Println(f1.nome)

}
