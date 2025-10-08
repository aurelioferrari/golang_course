package main

import "fmt"

//sem separar por vírgula
type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// método: função com receiver

func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

// aqui eu declaro a variável que terá atributos de produto
func main() {
	// aqui precisa de vírgula
	produto1 := produto{
		nome:     "lápis",
		preco:    1.79,
		desconto: 0.05,
	}
	fmt.Println(produto1, produto.precoComDesconto(produto1))

	produto2 := produto{"Notebook", 1989.90, 0.10}
	fmt.Println(produto2.precoComDesconto())
}
