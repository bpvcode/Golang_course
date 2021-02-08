package main

import (
	"fmt"
	"time"
)

func speak(person, text string, qtd int) {
	for i := 0; i < qtd; i++ {
		time.Sleep(time.Second) //Espera 1 segundo
		fmt.Printf("%s: %s (iteração %d)\n", person, text, i+1)
	}
}
func main() {
	/*
		speak("Bruno", "OLAAAAA", 3)
		speak("João", "Falaaa", 1)
	*/
	//NOTA: O João só "fala" depois do Bruno acabar de dizer as 3 frases

	//go speak("Bruno", "OLAAAAA", 500)
	//go speak("João", "Falaaa", 500)

	/*
		AO PASSAR A KEYWORD 'go' - CRIA UMA GO ROUTINE (vai executar código de forma independente - como se fosse uma thread)
		AO INICIAR UM PROGRAMA, A FUNÇÃO MAIN CRIA UMA LINHA DE EXECUÇÃO, DURANTE O DECORRER DO PROGRAMA, OUTRAS LINHAS (threads) PODEM SER CRIADAS
		NO CASO DO go NÃO CRIAMOS THREADS, MAS SIM GO ROUTINES - Funções que são executadas de forma independente
		Go Routine só vai continuar a executar se a thread principal do programa estiver a funcionar
	*/

	go speak("Bruno", "OLAAAAA", 10)
	speak("João", "Falaaa", 5)
	/*
		Reparar que o Bruno só falou algumas vezes e não 10 como estava estipulado
		Isto acontece porque a função go speak fica a executar numa go routine independente, e assim que o fio principal de execução acabar (ou seja que o joao fale 5 vezes e nada mais haja no main para correr) acaba o programa
	*/
}
