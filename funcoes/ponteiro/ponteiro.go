package main

import "fmt"

func exemplo(a int) {
	a++
	fmt.Println("a)", a)
}

//Ponteiro é representado por *
func ponteiroComoParametro(a *int) { //recebe o endereço de memória para o qual 'a' aponta
	// Ter acesso ao valor de a -> *a (desreferenciar)
	*a++
}

func main() {
	b := 1
	fmt.Println("b)", b)

	exemplo(b)
	//Passa cópia de valor de 'a' para a função
	//Imprime valor de 'a' como 2. Contudo 'a' só esta disponível dentro da função exemplo

	fmt.Println("b)", b)
	//b mantém-se inalterado, o que alteramos foi o valor de 'a' que iniciou como sendo uma cópia do valor de 'b'

	//Alterar o real valor de 'b', uma vez que estamos a passar o endereço de memória
	ponteiroComoParametro(&b)
	fmt.Println("b)", b)
}

/*
	CONCLUSÃO:
		* Um ponteiro é representado por *  -> a *int (Ou seja, faz referência ao endereço de memória de 'a')
		* Desreferenciar é acessar o valor do ponteiro -> *a (valor de 'a')
		* Obter endereço de memória de uma variável é representado por &  -> &a (endereço de memória de 'a')

		* Ao utilizarmos a função ponteiroComoParametro() estamos a mexer em valores que estão fora da função (Deixa de ser uma função pura - NÃO É BOA PRÁTICA DE FAZER EM GO)
		* Em GO tentamos sempre trabalhar com cópias de valores e não referências
*/
