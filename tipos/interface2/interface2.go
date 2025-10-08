package main

import "fmt"

// type esportivo interface {
// 	ligarTurbo()
// }

type ferrari struct {
	modelo          string
	turbo           bool
	velocidadeAtual int
}

func (f *ferrari) ligarTurbo() {
	f.turbo = true
}

func main() {
	carro1 := ferrari{"F40", false, 0}
	carro1.ligarTurbo()
	fmt.Println(carro1)
}
