package main

import "fmt"

func main() {

	// ======= append() =======
	arr := [3]int{1, 2, 3} // Declarar e instanciar/inicializar uma variável do tipo array de inteiros com 3 elementos
	fmt.Println("'arr' :", arr)

	var slc []int // Declarar uma variável slice sem nenhuma referência atribuída (Vazia)
	fmt.Println("'slc' vazio:", slc)
	slc = append(slc, 4, 5, 6) // append() - Permite aplicar a um slice acrescentado ao mesmo os valores que lhe forem passados
	fmt.Println("'slc' depois de append():", slc)

	slc1 := arr[:3]
	fmt.Println("Slice1:", slc1)
	slc1 = append(slc1, 7, 8, 9)
	fmt.Println("Slice1 after append():", slc1)

	// ======= copy() =======

	slc2 := make([]int, 2) //Declarar e inicializar um slice com 2 posições ainda sem elementos
	fmt.Println("Slice2:", slc2)
	copy(slc2, slc1) // Copiar para o 'slc2' os elementos do 'slc1'
	fmt.Println("Slice2 after copy():", slc2)
	//Como o 'slc2' tem um tamanho fixo de 2, copiou os primeiros 2 elementos de 'slc1'

	/*
		append(slc, elemento, elemento, ...):
			Adiciona elementos dentro de um slice
			Se o slice esta no tamanho maximo do array interno, substitui por outro com o dobro da capacidade, sem termos de nos preocupar
			Faz com que o slice cresça

		copy(slc, slc1):
			Não faz com que o slice cresça, so copia a quantida de elementos possíveis para a quantidade de elementos do slice de destino
			Os argumentos tem de ser do tipo slice
	*/
}
