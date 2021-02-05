package main

import "fmt"

//GUARDAR FUNÇÕES EM VARIÁVEIS - conceito fundamental em go permite usar princípios do paradigma funcional (Apesar de nao ser o core da linguagem)
/*
	As funções em GO podem:
 	* ser guardadas em variáveis
	* ser passadas com argumentos
	* retornar outras funções
/*
	A variável recebe o resultado de uma função anónima (função que nao tem nome)
	A partir do nome da variável a função pode ser executada soma(2,2) -> Output:4
*/
var soma = func(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println(soma(2, 2))

	subtração := func(a int, b int) int { //Função anónima
		return a - b
	}

	fmt.Println(subtração(2, 2))
}
