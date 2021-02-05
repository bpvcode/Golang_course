package main

import "fmt"

func factorial(valor uint) uint {
	//Colocar uint - faz com que nao possa receber valores negativos. Neste caso a função deixa de ter de retornar um erro, pois qualquer valor positivo pode ter factorial
	switch {
	case valor == 0:
		return 1
	default:
		factorialAnterior := factorial(valor - 1)
		return valor * factorialAnterior
	}
}

func main() {

	resultado := factorial(5)
	fmt.Println(resultado)

	fmt.Println(factorial(2))

}

// ================================================================
/*
var counter int = 0
var total int = 0

func potenciaResultado(base int, potencia int) {
	counter++
	fmt.Println("VOLTA Nº:", counter)
	if potencia == 0 {
		fmt.Println(base)
	}
	if base == 0 {
		fmt.Println(0)
	}
	if counter == 1 {
		total += base * base
		fmt.Printf("TOTAL: %d (Volta nº: %d) \n", total, counter)
		potenciaResultado(base, potencia)
	}
	if counter < potencia {
		total += total * base
		fmt.Printf("TOTAL: %d (Volta nº: %d) \n", total, counter)
		potenciaResultado(base, potencia)
	} else {
		fmt.Println("TOTAL FINAL:", total)
		counter = 0
		total = 0
	}
}

func main() {
	potenciaResultado(2, 2)
	fmt.Println("===========")
	potenciaResultado(3, 3)
}
*/
