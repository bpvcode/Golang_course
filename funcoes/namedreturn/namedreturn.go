package main

import "fmt"

// Podemos ter named return (Atribuir nomes para o retorno)
func troca(p1, p2 int) (segundo int, primeiro int) {
	segundo = p2
	primeiro = p1
	return
	/*
		NOTA:
		* Clean return -> Uma vez que nomeamos as variáveis de retorno na assinatura da função (segundo/primeiro),
		só precisamos de dar 'return' e ele retorna as duas variaiveis que ja receberam valores, na ordem declarada
		na assinatura (segundo, primeiro)
		* Ambas as variáveis so estão disponíveis no scope da função
	*/
}

func main() {
	r, r1 := troca(2, 1)
	fmt.Println("Primeiro:", r)
	fmt.Println("Segundo:", r1)
	/*Ou seja:
	Ao 'r' será atribuído o segundo valor de retorno
	Ao 'r1' o primeiro
	*/
}
