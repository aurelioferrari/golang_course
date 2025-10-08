package main

import "fmt"

func main() {
	var a int
	var b float64
	var c bool
	var d string
	var e *int
	fmt.Printf("%v %v %v %q %v", a, b, c, d, e)

	// valores-padrão: 0 para int e float, bool é false, string é vazia e ponteiro é <nil>
}
