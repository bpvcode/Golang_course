package main

import (
	"fmt"
	"time"
)

//speak() funciona como generator (encapsula a a criação do chanel e a criação de go routines e retorna dados do chanel como resposta)
//Quem utilizar esta função nao vai conseguir escrever no canal, mas vai conseguir ter acesso aos dados que estão dentro do mesmo
func speak(person string) <-chan string {
	channel := make(chan string)

	go func() { //Criação de uma go routine
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second) //A cada segundo é enviada uma string para o canal (String a baixo)
			channel <- fmt.Sprintf("%s speaks: %d", person, i)
		}
	}() //Função anónima, tem de ser invocada assim que é declarada

	return channel
}

func aggregate(channel1, channel2 <-chan string) <-chan string {
	channel := make(chan string) //Canal a retornar

	go func() {
		for {
			select {
			case resp1 := <-channel1:
				channel <- resp1
			case resp2 := <-channel2:
				channel <- resp2
			}
		}
	}()
	return channel
}

func main() {
	person1 := speak("Bruno")
	person2 := speak("Rui")

	finalChannel := aggregate(person1, person2)

	fmt.Println(<-finalChannel, <-finalChannel)
	fmt.Println(<-finalChannel, <-finalChannel)
	fmt.Println(<-finalChannel, <-finalChannel)

}
