package main

import "fmt"

func multiplicação(a int, b int) int {
	return a * b
}

func exec(função func(int, int) int, p1 int, p2 int) int {
	/*
		Estou à espera de receber como argumento, uma função que recebe 2 parâmetros inteiros e retorna 1 inteiro (função multiplicação)
		Em seguida passo mais 2 parâmetros do tipo int para a função (p1,p2)
		Esta função também retorna um int
		Então, vou pegar nos dois parâmetros que recebo (p1, p2) e passar para dentro da função multiplicação
	*/

	return função(p1, p2)
}

func main() {
	resultado := exec(multiplicação, 2, 2) //Passar a função como parâmetro e em seguida os valores que a mesma vai receber de parâmetros
	fmt.Println(resultado)
}
