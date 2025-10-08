package main

import "fmt"

// init é chamado antes de main e mesmo sem ser chamada

func init() {
	fmt.Println("Inicializando...")
}

func main() {
	fmt.Println("Main...")
}
