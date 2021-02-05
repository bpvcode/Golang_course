package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y)) // converter int para float64 (tipo cast no java)

	nota := 6.9
	notaFinal := int(nota) //converter float64 para int (tipo cast no java)
	fmt.Println(notaFinal)

	//==== CUIDADO ====
	//Assim representa o caractere correspondente ao numero 97 da tabela Ascii neste caso a letra 'a'
	fmt.Println("CUIDADO: Converter inteiro para string e aplicar concatenação:" + string(97))

	//==== Converter int para string ====
	a := "Teste " + strconv.Itoa(97)
	//Agora o numero 97 nao e mais um int mas sim uma string
	fmt.Println(a)
	fmt.Println("O tipo de a é:", reflect.TypeOf(a))

	//==== Converter string para int ====
	b, err := strconv.Atoi("97")
	//strconv.Atoi() retorna 2 valores, ou seja podemos lidar com o erro como estamos a fazer agora, ou utilizar '_' para ignorar o erro
	if err != nil {
		panic("PANICO STRING TO INT")
	}
	fmt.Println(b)
	fmt.Println(98 - b)
	fmt.Println("O tipo de b é:", reflect.TypeOf(b))

	//==== Converter string para boolean ====
	c, err := strconv.ParseBool("0")
	//strconv.ParseBool() retorna 2 valores, ou seja podemos lidar com o erro como estamos a fazer agora, ou utilizar '_' para ignorar o erro.
	//strconv.ParseBool() aceita: "true","false","t","f", "1", 0
	if err != nil {
		panic("PANICO STRING TO BOOLEAN")
	}
	fmt.Println(c)
	fmt.Println(!c)
	fmt.Println("O tipo de c é:", reflect.TypeOf(c))
}
