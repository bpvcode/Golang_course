package main

import "fmt"

/*
	Funções Variaticas
	São funções que recebem quantidade de parâmetros variáveis (aprovados ...)
*/

func media(numeros ...float64) float64 {
	//Função que pode receber parâmetros variáveis '...'
	//Os parâmetros são tratados como um array
	total := 0.0

	for _, valor := range numeros { // '_' -> Indice || 'valor' -> elemento
		total += valor
	}

	return total / float64(len(numeros)) //cast para float64 porque len() retorna int

}

func main() {
	fmt.Printf("Média: %.2f \n", media(2.2, 3.4, 5.5))
	fmt.Printf("Média: %.2f \n", media(7.7, 8.1, 5.9))
	fmt.Printf("Média: %.2f \n", media(7.7, 8.1, 5.9, 9.9))
}
