package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

// interfaces são implementadas implicitamente

func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (pr produto) toString() string {
	return fmt.Sprintf("%s - R$ %f", pr.nome, pr.preco)
}

func main() {
	prod1 := produto{"Sabonete", 1.23}
	fmt.Println(prod1.toString())

}
