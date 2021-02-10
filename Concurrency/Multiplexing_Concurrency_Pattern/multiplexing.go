package main

import (
	"fmt"

	"github.com/bpvcode/reusablePackage"
)

/*
	MULTIPLEX:
	* Pegar em dois ou mais channels que geram dados e juntar num único channel com a junção das informações dos canais
*/
func receiveChannelDataAndSendToAnotherChannel(origin <-chan string, destino chan string) {
	for { //For infinito estilo while(true)
		destino <- <-origin
		/*
			EXPLICAÇÃO:
			<-origin (Pegar nos dados do canal de origem. Passamos dados e não o canal inteiro)
			destino <- (Receber dados no canal de destino)
		*/
	}
}

//FUNÇÃO RESPONSÁVEL POR MULTIPLEXAR
func aggregate(channel1, channel2 <-chan string) <-chan string { //Argumentos: Canais somente de leitura
	channel := make(chan string)
	go receiveChannelDataAndSendToAnotherChannel(channel1, channel)
	go receiveChannelDataAndSendToAnotherChannel(channel2, channel)
	return channel
}

func main() {
	chan1 := aggregate(
		//Estamos a reutilizar o package que temos em '/home/brunovilar/go/src/github.com/bpvcode/reusablePackage'
		reusablePackage.Title("http://www.google.com", "http://www.facebook.com"), //Método retorna um channel
		reusablePackage.Title("http://www.udemy.com", "http://www.youtube.com"),   //Método retorna um channel
		//Conclusão juntamos os dois canais dentro de um único canal 'chan1'
	)

	fmt.Println("Primeiro a chegar:", <-chan1)
	fmt.Println("Segundo a chegar:", <-chan1)
	fmt.Println("Terceiro a chegar:", <-chan1)
	fmt.Println("Quarto a chegar:", <-chan1)
}
