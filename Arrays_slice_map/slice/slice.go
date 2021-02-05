package main

import (
	"fmt"
	"reflect"
)

func main() {
	arr := [3]int{1, 2, 3} // == array == -> Estrutura de tamanho fixo (neste caso 3 elementos)
	fmt.Println("Isto é um array:", arr)
	fmt.Println("Do tipo:", reflect.TypeOf(arr))

	slc := []int{1, 2, 3} // == slice == -> Estrutura de tamanho variável
	fmt.Println("Isto é um slice:", slc)
	fmt.Println("Do tipo:", reflect.TypeOf(slc))

	/*
		MUITO IMPORTANTE:
		* Podemos ter um array interno e vários slices apontando para partes especificas desse mesmo array (Ex: Um slice aponta para o primeiro e segundo elemento ; Outro slice aponta para o terceiro e quarto elemento)
		* Slice ("Uma parte de algo; Uma fatia de um bolo" - O slice "so aponta para uma fatia do bolo inteiro")
		* Exemplo:
			O slice tem uma estrutura com um tamanho de 3 elementos
			Temos uma referência dentro da estrutura do slice que aponta para um elemento que está a ser referenciado e que está dentro de um array
			Então slice define um pedaço de um array
		* Slice não cria um novo array, ele cria uma estrutura que têm um ponteiro para um elemento de um array e um tamanho que segue a partir desse ponteiro
	*/

	arr2 := [5]int{1, 2, 3, 4, 5}

	slc2 := arr2[1:3]
	// Slice que referêncía o array2 e aponta para o elemento de indice 1 INCLUSIVE e indice 3 NÃO INCLUSIVE
	// Output: [2,3]
	fmt.Printf("Array de 5 elementos: %v || Slice que referêncía array: %v \n", arr2, slc2)

	slc3 := arr2[:2]
	// Este slice vai do inicio do array 2 (do indice 0), até ao indice 2 NÂO INCLUSIVE (Ou seja, indice 1)
	// Novo slice que aponta para o mesmo array arr[2]
	fmt.Println("Slice 3 ->", slc3)

	slc4 := slc2[:1]
	// Um slice de um slice. Este apon
	// Vai apontar para o slc2 e guarda o indice 0 (vai do 0 ao 1 nao inclusive)
	fmt.Println("Slice 4 (slice de um slice) ->", slc4)

	//NOTA FINAL: Todos estes slices apontam para o mesmo array

}
