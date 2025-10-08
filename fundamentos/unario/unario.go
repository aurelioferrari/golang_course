package main

import "fmt"

func main() {
	x := 1
	y := 2

	// apenas postfix
	x++ // ou x += 1 ou x = x + 1
	fmt.Println(x)

	y-- // ou y-= 1
	fmt.Println(y)

	// nao podemos fazer (x == y--)

}
