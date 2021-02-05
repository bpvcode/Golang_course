package main

import "fmt"

func main() {
	/*
		Array recebe estes valores, logo compilador assume as posições totais do array (neste caso 5 posições de 0 a 4)
		Se não tiver os '...' dentro de [] é um slice (A ver mais a frente)
	*/
	arr := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	//=== FOR RANGE com indice ===
	for i, numero := range arr {
		fmt.Printf("Indice %d => Elemento %d \n", i, numero)
		/*
			i -> Representa o indice
			numero -> Representa o valor referente ao indice i
		*/
	}

	//=== FOR RANGE sem indice ===
	for _, numero := range arr {
		fmt.Printf("Agora só me interessa os elementos dentro do array: %d \n", numero)
	}

	/* === ATENÇÃO ===
	NOTA MT IMPORTANTE:
	* Se quisermos só pegar no valor de cada indice tem de ser como o for escrito a cima
	* Se so passar um elemento, esse elemento representa sempre o indice, ver exemplo a baixo
	*/

	for numero := range arr {
		fmt.Printf("Numero agora faz referência ao indice: %d \n", numero)
	}
}
