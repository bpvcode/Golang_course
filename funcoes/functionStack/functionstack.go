package main

import (
	"fmt"
)

func f1() {
	f2()
	fmt.Println("SOU O TERCEIRO A APARECER")
}

func f2() {
	f3()
	fmt.Println("SOU O SEGUNDO A APARECER")
}

func f3() {
	// debug.PrintStack() // Vai imprimir o stack no momento em que a função for chamada (Verificar que a primeira coisa a ser chamada é este método)
	fmt.Println("SOU O PRIMEIRO A APARECER")
}

func main() {
	f1()

	/*
		NOTA MUITO MUITO IMPORTANTE:
		* A primeira ação a ser executada num caso em que uma função chama outra, será a primeira ação do corpo da função, que deixa de chamar outra.
	*/
}
