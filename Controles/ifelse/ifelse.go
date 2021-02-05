package main

import "fmt"

func imprimirResultado(nota float64) {
	if nota < 5 {
		fmt.Println("REPROVADO:", nota)
	} else {
		fmt.Println("APROVADO:", nota)
	}
}

func main() {
	imprimirResultado(7.3)
	imprimirResultado(4.9)
}
