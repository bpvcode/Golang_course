package main

import "fmt"

func main() {

	//PROVA DE QUE OS ARRAYS INTERNOS SÃO OS MESMOS EM AMBOS OS SLICES:
	slc := make([]int, 10, 20)
	slc1 := append(slc, 1, 2, 3)
	fmt.Printf("Slice: %v || Tamanho slice: %v || Tamanho array interno: %v\n", slc, len(slc), cap(slc))
	fmt.Printf("Slice: %v || Tamanho slice: %v || Tamanho array interno: %v\n", slc1, len(slc1), cap(slc1))

	//Alterando só um slice
	slc[0] = 1

	//Ambos alteraram apesar de a alteração ter sido feita só num
	fmt.Println("ALTERADO no primeiro elemento:")
	fmt.Printf("Slice: %v || Tamanho slice: %v || Tamanho array interno: %v\n", slc, len(slc), cap(slc))
	fmt.Printf("Slice: %v || Tamanho slice: %v || Tamanho array interno: %v\n", slc1, len(slc1), cap(slc1))

}
