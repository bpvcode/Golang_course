package main

import (
	"fmt"
	"time"
)

func main() {
	//Operadores relacionais (Retornam true OU false)
	fmt.Println("Strings:", "Banana" == "Banana")
	fmt.Println("!=", 3 != 2)
	fmt.Println(">=", 4 >= 2)
	fmt.Println("<=", 4 <= 2)
	fmt.Println("== em ints:", 4 == 4)
	fmt.Println(">", 3 > 3)
	fmt.Println("<", 4 < 4)

	d1 := time.Unix(0, 0) //Dado do tipo data. Retorna a data do dia de hoje. (0,0) de argumento representa o dia inicial
	d2 := time.Unix(0, 0)
	fmt.Println(d1)
	fmt.Println(d2)

	fmt.Println("ATENÇÃO:", d1 == d2)
	/*
		CONCLUSÃO:
			== refere-se ao valor do dado e não à referência de memória do mesmo
	*/
	fmt.Println("Equal() method:", d1.Equal(d2))

	//SCRUCT

	type Pessoa struct {
		Nome  string
		Idade int
	}

	p1 := Pessoa{"Bruno", 27}
	p2 := Pessoa{"Bruno", 27}
	fmt.Println("Pessoas iguais:", p1 == p2)
	/*
		CONCLUSÃO:
			== refere-se ao valor do dado e não à referência de memória do mesmo
			p1 e p2 têm referências de memória diferentes, mas têm os mesmos valores para todas as propriedades(Nome , Idade)
	*/
}
