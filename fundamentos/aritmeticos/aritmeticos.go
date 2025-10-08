package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2
	fmt.Println("Soma =>", a+b)
	fmt.Println("Subtração =>", a-b)
	fmt.Println("Divisão =>", float64(a)/float64(b))
	fmt.Println("Multiplicação =>", a*b)
	// resto da divisão
	fmt.Println("Resto da divisão =>", a%b)
	//bitwise
	fmt.Println("E =>", a&b)
	fmt.Println("Ou =>", a|b)
	fmt.Println("Xor =>", a^b)

	c := 3.0
	d := 2.0

	// outras operações usando math
	// math.Max ou math.Min precisa ter números float como entrada
	fmt.Println("Maior =>", math.Max(c, d))
	fmt.Println("Menor =>", math.Min(c, d))
	fmt.Println("Exponencial =>", math.Pow(c, d))

}
