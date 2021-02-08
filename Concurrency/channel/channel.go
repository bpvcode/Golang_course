package main

import "fmt"

//Channel é um tipo da linguagem (tal como int, string, float64, etc..)

func main() {

	ch := make(chan int, 1) //Criar channel
	/*
		1º argumento - Tipo do channel - Criou um channel de inteiros (Envia valores inteiros)
		2º argumento - Buffer (tamanho de 1)
	*/

	ch <- 1 //Enviar dados para o canal com o valor de 1 (Visto que ele recebe valores inteiros)
	<-ch    //Receber dados do canal

	ch <- 2
	fmt.Println(<-ch) //Imprimir resultado que foi recebido pelo canal (podiamos também guardar numa variável)

	/*
		Canal é a forma de comunicação entre as Go Routines
		Canal é um ponto de sincronismo. Canal fica à espera de receber o valor que foi estipulado (no caso anterior inteiro)
	*/
}
