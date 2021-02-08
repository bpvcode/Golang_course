package main

import (
	"fmt"
	"time"
)

func routine(channel chan int) {
	time.Sleep(time.Second)
	channel <- 1 //Operação bloqueante (Uma vez que estamos a enviar um dado para um canal, ele só vai prosseguir quando alguém ler esse dado)
	fmt.Println("Só depois de ser lido")
}

func main() {
	channel := make(chan int) //Criar channel sem buffer (mal recebe dado, fica bloqueado, dado tem de ser lido, para podermos continuar a passar dados para o canal)
	go routine(channel)
	fmt.Println(<-channel)

	fmt.Println("Foi lido")
	fmt.Println(<-channel) //Vai gerar deadlock (pois não tem mais dados para ler do canal)
	fmt.Println("FIM")     //Não chega a ser impresso porque antes gera uma deadlock

	/*
		Quando chamamos 'go routine(channel)' automaticamente o programa passar para aproxima linha (imprimi o dado que veio do canal), contudo terá de esperar 1 segundo para receber o dado
	*/
}
