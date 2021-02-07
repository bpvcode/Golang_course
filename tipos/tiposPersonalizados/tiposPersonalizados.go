package main

import "fmt"

type nota float64 //Criamos um tipo nosso a partir de um tipo float64

func (n nota) resultadoNota(inicio, fim float64) bool { //Criamos um método que só vai ser usado pelo type que criamos acima, que será o receiver (n nota)
	return float64(n) >= inicio && float64(n) <= fim
}

func notaParaPassar(n nota) string {
	if n.resultadoNota(9.0, 10.0) {
		return "A"
	} else if n.resultadoNota(7.0, 8.99) {
		return "B"
	} else if n.resultadoNota(5.0, 6.99) {
		return "C"
	} else if n.resultadoNota(3.0, 4.99) {
		return "D"
	} else {
		return "E"
	}
}

func notaParaPassarRefactor(n nota) string {
	switch {
	case n.resultadoNota(9.0, 10.0):
		return "A"
	case n.resultadoNota(7, 8.99):
		return "B"
	case n.resultadoNota(5.0, 7.99):
		return "C"
	case n.resultadoNota(3.0, 4.99):
		return "D"
	default:
		return "E"
	}
}

func main() {
	fmt.Println(notaParaPassar(9.8))
	fmt.Println(notaParaPassar(6.9))
	fmt.Println(notaParaPassar(2.1))
	fmt.Println("====================")
	fmt.Println(notaParaPassarRefactor(9.8))
	fmt.Println(notaParaPassarRefactor(6.9))
	fmt.Println(notaParaPassarRefactor(2.1))
}
