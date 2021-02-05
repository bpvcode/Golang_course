package main

import "fmt"

func factorial(valor int) (int, error) {
	switch {
	case valor < 0:
		return -1, fmt.Errorf("Número invalido: %d", valor) //Valores negativos geram erros
	case valor == 0:
		return 1, nil //Valor 0 não gera erro (nil -> erro é null)
	default:
		factorialAnterior, _ := factorial(valor - 1)
		return valor * factorialAnterior, nil
	}
}

func main() {
	//Como ignorei o erro imprime só o resultado e não imprime o nil
	resultado, _ := factorial(5)
	fmt.Println(resultado)
	//Como não ignorei o resultado nem o erro imprime o resultado e o nil
	fmt.Println(factorial(2))
	//Como não ignorei o resultado nem o erro imprime o resultado e o erro
	fmt.Println(factorial(-1))
	//Como ignorei o resultado no primeiro, vou tratar o valor do erro, imprime só o erro
	_, err := factorial(-1)
	if err != nil {
		fmt.Println(err)
	}

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
