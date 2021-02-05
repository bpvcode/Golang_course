package main

import (
	"fmt"
	"math"
)

func main() {
	a := 1
	b := 2

	fmt.Println("Soma =>", a+b)
	fmt.Println("Subtração =>", a-b)
	fmt.Println("Multiplicação =>", a*b)
	fmt.Println("Divisão =>", a/b)
	fmt.Println("Módulo =>", a&b)

	//NOTA: Operações têm de ser efetuadas com variáveis/argumentos do mesmo tipo. (int + int)

	//bitwise - Operadores bit a bit ???
	fmt.Println("E =>", a&b)
	fmt.Println("OU =>", a|b)
	fmt.Println("Xor =>", a^b)

	//outras operações usando math...
	c := 3.0
	d := 2.0

	fmt.Println(math.Round(c))
	fmt.Println(math.RoundToEven(d))

	/*
		NOTA IMPORTANTE:
			A maioria dos métodos math recebem argumentos do tipo float64
		Exemplo nas linha abaixo:
			Ao usar as variáveis 'a' e 'b', do tipo int, tivemos de fazer um cast para float64().
			Ao usar as variáveis 'c' e 'd', do tipo float64, não temos de fazer nada
	*/
	fmt.Println("Maior entre os dois numeros =>", math.Max(float64(a), float64(b)))

	fmt.Println("Maior entre os dois numeros =>", math.Max(c, d))
	fmt.Println("Menor entre os dois numeros =>", math.Min(c, d))
	fmt.Println("c elevado a d =>", math.Pow(c, d))

}
