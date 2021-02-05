package main

import "fmt"

func main() {

	//Atribuição
	var b byte = 3
	fmt.Println(b)

	//Atribuição simples
	c := 3
	c += 2 // c = c + 2
	c -= 2 // c = c - 2
	c /= 2 // c = c / 2
	c *= 2 // c = c * 2
	c %= 2 // c = c % 2

	fmt.Println(c)

	//Inicialização e Atribuição a mais do que uma variavel
	d, e := 1, 2
	fmt.Println(d)
	fmt.Println(e)

	//Atribuição a mais do que uma variavel (troca dos valores das variáveis entre si)
	d, e = e, d
	fmt.Println(d)
	fmt.Println(e)

}
