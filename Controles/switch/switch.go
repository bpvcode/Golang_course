package main

import "fmt"

func nota(n float64) string {
	var notaFinal = int(n)

	switch notaFinal {
	case 10:
		return "hello: 10"
		// fallthrough (ver explicação abaixo)
	case 15:
		return "hey: 15"
	case 20, 25:
		return "hey: 20 ou 25"
	default:
		return "Nota invalida"
	}
	//tem que ter default!!
}

func main() {
	fmt.Println(nota(10))
	fmt.Println(nota(15))
	fmt.Println(nota(20))
	fmt.Println(nota(25))
	fmt.Println(nota(30))
	fmt.Println(nota(-1))
}
