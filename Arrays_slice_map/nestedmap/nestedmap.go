package main

import "fmt"

func main() {

	/*
		Map que recebe uma key:int e um value:map[string]float64
		Map dentro de map
	*/
	nestedMap := map[string]map[string]float64{
		"A": {"Ana": 1250.5,
			"Andreia":   6000.0,
			"Antonieta": 2500.1,
		},
		"B": {"Bruno": 2500.6,
			"Bernardo": 750.65,
			"Bebiana":  987.60,
		},
	}
	fmt.Println(nestedMap)

	//Apagar todos os funcionarios com letra "A"
	delete(nestedMap, "A")
	fmt.Println(nestedMap)

	//Iterar sobre nested maps (for dentro de for)
	count := 0

	for letra, funcionario := range nestedMap {
		fmt.Println("Letra:", letra)

		for nome, salario := range funcionario {
			count++
			fmt.Printf("=== Funcionário %d === \n* Nome: %s \n* Salário: %f \n", count, nome, salario)
		}
	}
}
