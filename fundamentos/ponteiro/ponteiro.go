package main

import "fmt"

func main() {
	i := 1
	// Go não tem aritmética de ponteiros
	// notação de ponteiro é *
	var p *int = nil

	p = &i // atribui o endereço de memória que aponta pro mesmo local que a var i
	// para ter agora acesso à informação desse endereço de memória:
	*p++

	fmt.Println(p, *p, i)
}
