package main

import (
	"fmt"
	"time"
)

func routine(channel chan int) {
	channel <- 1
	channel <- 2
	//fmt.Println("Executou, vou enviar valor para o canal")
	channel <- 3
	channel <- 4
	channel <- 5
	fmt.Println("Executou, vou enviar valor para o canal")
	channel <- 6
}

func main() {
	channel := make(chan int, 3) //Criou um canal de inteiros com um buffer de 3 posições (Consegue armazenar 3 dados sem bloquear)
	go routine(channel)

	time.Sleep(time.Second)
	fmt.Println(<-channel)
}

/*
	Quando o buffer fica cheio, o canal fica bloqueado ate que seja consumido parte do buffer para colocar mais informação dentro do buffer
	Para testar isso basta descomentar a linha 11 e perceber que imprime na consola, contudo se o canal tiver cheio, a linha 15 já não imprime
*/
