package main

import "fmt"

func imprimirPontoDaCarne(valor float64) {
	if valor > 40 && valor < 50 {
		fmt.Println("Esta no ponto")
	} else if valor > 35 && valor <= 40 {
		fmt.Println("Mal passado")
	} else {
		fmt.Println("Desperdicio")
	}
}

func main() {
	imprimirPontoDaCarne(45)
	imprimirPontoDaCarne(40)
	imprimirPontoDaCarne(36)
	imprimirPontoDaCarne(10)
	imprimirPontoDaCarne(60)
}
