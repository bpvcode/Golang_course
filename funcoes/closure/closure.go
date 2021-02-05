package main

import "fmt"

/*
	Função closure (todas as funções em GO):
	Quando uma função "envolve" outra função
	Faz com que a função interna (que foi envolvida pela função externa) tenha memória do local onde foi declarada.
	Quando declarada uma função dentro de outra função -> A função interna tem acesso ao scope da função externa
*/

func closure() func() {
	x := 10
	var funcao = func() {
		fmt.Println(x)
	}
	return funcao
}

func main() {
	x := 20
	fmt.Println(x)

	imprimeX := closure() //Como closure() retorna uma função, imprimeX é uma variável que recebe uma função, pelo que podemos invocar a função através da variável (linha abaixo)
	imprimeX()
	//Neste caso, a variável x que ele vai ler refere-se ao scope da função externa closure(), por isso o valor é 10
}
