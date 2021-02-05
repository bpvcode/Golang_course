package main

import "fmt"

/*
	PROBLEMA:
	* Tenho 2 trabalhos para fazer esta semana.
	* Se fizer o trabalho 1, compro uma tv de 50 polegadas
	* Se fizer o trabalho 2, compro uma tv de 32 polegadas
	* Se fizer qualquer um dos trabalhos, vou comer um gelado
	* Se não fizer nenhum trabalho, fico em casa, e fico mais saudável por nao comer gelados.

*/

func compras(trabalho1, trabalho2 bool) (bool, bool, bool) {
	comprarTv50 := trabalho1 && trabalho2
	comprarTv32 := trabalho1 != trabalho2
	comprarGelado := trabalho1 || trabalho2

	return comprarTv50, comprarTv32, comprarGelado

}

func main() {
	tv50, tv32, gelado := compras(true, true)
	fmt.Printf("TV50: %v \nTV32: %v \nGelado: %v \nSaudável: %v", tv50, tv32, gelado, !gelado)
}
