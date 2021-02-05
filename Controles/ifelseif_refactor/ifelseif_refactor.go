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

func refactor(valor float64) string {
	switch {
	case valor > 40 && valor < 50:
		return "Esta no ponto"
	case valor > 35 && valor <= 40:
		return "Mal passado"
	default:
		return "Desperdício"
	}
}

func main() {
	imprimirPontoDaCarne(45)
	imprimirPontoDaCarne(40)
	imprimirPontoDaCarne(36)
	imprimirPontoDaCarne(10)
	imprimirPontoDaCarne(60)

	fmt.Println("============")

	fmt.Println(refactor(45))
	fmt.Println(refactor(40))
	fmt.Println(refactor(36))
	fmt.Println(refactor(10))
	fmt.Println(refactor(60))
}
