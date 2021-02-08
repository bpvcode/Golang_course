package main

import (
	"fmt"
	"time"
)

func multiply(base int, channel chan int) {
	time.Sleep(time.Second)
	channel <- 2 * base //Enviar dados para o canal

	time.Sleep(time.Second)
	channel <- 3 * base

	time.Sleep(time.Second * 3) //Espera 3 segundos
	channel <- 4 * base
}

func main() {
	channel := make(chan int)
	go multiply(2, channel)

	a, b := <-channel, <-channel //Variáveis a e b recebem dados do canal - 2*base ; 3*base
	fmt.Println(a, b)

	fmt.Println(<-channel) //Imprime o ultimo valor a ser recebido no canal - 4*base
	//Se tentar receber mais um valor no channel, gera um deadlock pois o channel so receberia 3 valores que ja foram "pedidos" (<-channel) a cima.

	/*
		Ao correr o código, reparar que o channel fica à espera de receber os valores de 'a' e 'b' e só depois imprime o ultimo valor recebido pelo channel (ultima linha)
	*/
}
