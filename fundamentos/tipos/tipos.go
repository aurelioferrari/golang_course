package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é tipo", reflect.TypeOf(32000))

	// sem sinal (só positivos) - uint8 uint16...
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	// com sinal int8 int16
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo da variável i1 é:", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tabela unicode (int32)
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	// números reais
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99))

	// boolean
	bo := true
	fmt.Println("O Booleano é do tipo:", reflect.TypeOf(bo))

	// ! nega o booleano
	fmt.Println(!bo)

	// string - Golang não permite aspas simples
	s1 := "Olá. Meu nome é Aurélio"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))

	// string com múltiplas linhas - se usa crase
	s2 := `Olá
	Essa é
	minha string
	quebrada`
	fmt.Println(s2)
	fmt.Println("O tamanho dela é", len(s2))

	// char??
	char := 'a'

	fmt.Println("O tipo de char é:", reflect.TypeOf(char))

}
