package main

import "fmt"

/*
Um array tem de ser do mesmo tipo (como java). Todos os elementos têm de ser do mesmo tipo
Arrays têm uma estrutura fixa (Se tem 10 posições, tem 10, nem mais nem menos)
*/
func main() {
	var notas [3]float64 // => Declaração de array do tipo float64 com 3 posições
	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 7.2, 8.4, 9.7
	fmt.Println(notas)

	totalNotas := 0.0

	for i := 0; i < len(notas); i++ {
		totalNotas += notas[i]
	}

	media := totalNotas / float64(len(notas))
	fmt.Printf("Media: %.2f\n", media)
}
