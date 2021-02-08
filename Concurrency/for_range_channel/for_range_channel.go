package main

import (
	"fmt"
	"time"
)

//Numero primo - valor que é divisível por 1 ou por ele mesmo (não pode ser divisível por mais nenhum valor e o modulo da divisão ser igual a 0)
func isPrimo(numero int) bool {
	for i := 2; i < numero; i++ {
		if numero%i == 0 {
			return false
		}
	}
	return true
}

func primos(n int, channel chan int) {
	//n - quantidade de números primos a retornar pela função
	//Os valores encontrados serão enviados para o canal

	inicio := 2 //Variável que vai sendo alterada sempre que um numero primo for encontrado (para nao começar sempre a procurar a partir do 2)

	for i := 0; i < n; i++ {
		for primo := inicio; ; primo++ {
			//FOR SEM NENHUM TIPO DE CONDIÇÃO ASSOCIADA
			if isPrimo(primo) {
				channel <- primo
				inicio = primo + 1 //Procurar a partir do primo achado mais 1
				time.Sleep(time.Second)
				break //Se encontrar um numero primo sai deste loop interno e continua para o loop até atingir o numero máximo de primos a retornar (n)
			}
		}
	}
	close(channel) //Dentro do canal não vai ser enviado mais nenhum dado
}

func main() {
	channel := make(chan int, 30)
	go primos(cap(channel), channel) //cap(channel) - vai averiguar a capacidade do channel (neste caso 30), pelo que queremos descobrir os 30 primeiros numeros primos
	for primo := range channel {
		fmt.Printf("%d \n", primo)
	} //For vai iterando sobre o channel à procura de valores, sempre que chega um novo, ele imprime. Assim que o channel for fechado, acaba este for. (Se nao tiver o close() vai gerar um deadlock)
	fmt.Println("FIM!")
}
