package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprartv50 := trab1 && trab2
	comprartv32 := trab1 != trab2
	comprarsorv := trab1 || trab2

	return comprartv50, comprartv32, comprarsorv
}

func main() {
	tv50, tv32, sorvete := compras(false, false)
	fmt.Printf("Tv50: %t\nTv32 %t\nSorvete %t\nSaudável? %t", tv50, tv32, sorvete, !sorvete)
}
