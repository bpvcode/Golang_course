package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {

	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	//sem sinal (só positivos)... TIPOS - uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b)) // O byte é um inteiro sem sinal com 8 bits (uint8)

	//com sinal... TIPOS - int8 int16 int32 int64
	i1 := math.MaxInt64 // Representa o maior valor possível de um inteiro com 64 bits
	fmt.Println("O valo máximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa o valor da letra 'a' na tabela Unicode (int32)
	fmt.Println("O tipo de rune é", reflect.TypeOf(i2))
	fmt.Println("O valor de i2 é", i2)

	/*
		Nota: por default o tipo de um float é float 34

		var x = 49.99 -> Tipo float34
	*/

	//** BOOLEAN **
	bo := true
	fmt.Println("O tipo da variável bo é", reflect.TypeOf(bo))
	fmt.Println("Coo é boolean o contrário de bo é", !bo)

	//** STRING **
	s1 := "OLA"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é:", len(s1)) // len() == length()

	//** STRING com multiplas linhas**
	s2 := `OLA
	SOU
	O
	BRUNO`
	fmt.Println(s2 + "!")
	fmt.Println("O tamanho desta string é:", len(s2))

	//** char == int32 **
	a := 'a' //Charcteres sempre com '' aspas unicas não "" aspas duplas
	fmt.Println("O tipo de a é:", reflect.TypeOf(a))

	test1, test2 := -2, byte(1)

	fmt.Println("O tipo de test1 é:", reflect.TypeOf(test1))
	fmt.Println("O tipo de test2 é:", reflect.TypeOf(test2))

}
