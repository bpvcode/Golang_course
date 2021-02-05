package main

import "fmt"

func main() {

	//Declarar e atribuir valores ao map
	listaDeFuncionários := map[string]float64{
		"Bruno": 1000.07,
		"Ana":   2000.80,
		"João":  909.77, //Virgula no ultimo !!IMPORTANTE
	}
	fmt.Println(listaDeFuncionários)

	//Alterando value(salario) através da key(nome)
	listaDeFuncionários["Bruno"] = 1200.5
	fmt.Println(listaDeFuncionários)

	//Adicionando elemento através da key(nome) atribuindo um value(salario)
	listaDeFuncionários["Pedro"] = 800
	fmt.Println(listaDeFuncionários)

	//Apagar elemento através da key(nome)
	delete(listaDeFuncionários, "Ana")
	fmt.Println(listaDeFuncionários)
	delete(listaDeFuncionários, "Inexistente") //Não gera erro apagar um elemento que nao existe

	//Percorrer o map()
	for nome, salario := range listaDeFuncionários {
		fmt.Printf("NOME: %s || Salario: %f \n", nome, salario)
	}

}
