package main

import "fmt"

func main() {

	// var aprovados map[int]string - declaração de um map de key:int , value:string (MAPAS DEVEM SER INICIALIZADOS)
	listaDeAprovados := make(map[int]string)

	// Adicionar valores a um map
	listaDeAprovados[123456] = "Bruno"
	listaDeAprovados[1547] = "Rui"
	fmt.Println(listaDeAprovados)

	// Percorrer a listaDeAprovados, um map(), com for range sendo que key:id , value:nome
	for id, nome := range listaDeAprovados {
		fmt.Printf("ID: %d | NOME: %s\n", id, nome)
	}

	//Como acessar o value que esta associado a uma certa key=1547
	println("Antes de delete():", listaDeAprovados[1547])
	//Como eliminar eliminar um elemento do mapa (key e value do mesmo) através da sua key=1547
	delete(listaDeAprovados, 1547)
	println("Depois de delete():", listaDeAprovados[1547])

	/*
		NOTA: Podemos eliminar ou key ou value utilizando '_'
	*/

}
