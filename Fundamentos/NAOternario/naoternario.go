package main

import "fmt"

//Não existe operador ternário em Go

func obterResultado(nota float64) string {
	// return nota>=6 ? "Passou" : "NAO PASSOU" -> NÃO EXISTE OPERADOR TERNÁRIO
	if nota < 6 {
		return "NAO PASSOU"
	}
	return "Passou"
}

func main() {
	fmt.Println(obterResultado(5))
}
