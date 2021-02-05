package main

import (
	"fmt"
	"time"
)

func main() {

	//NUMERO DETERMINADO DE REPETIÇÕES - Estilo while() do java
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	//NUMERO INDETERMINADO DE REPETIÇÕES - for loop normal do java
	for i := 0; i < 5; i++ {
		fmt.Print(i)
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("PAR")
		} else {
			fmt.Println("IMPAR")
		}
	}

	for i := 0; i <= 10; i += 2 {
		fmt.Print(i)
	}

	//LOOP INFINITO
	for {
		fmt.Println("loop infinito")
		time.Sleep(time.Second) //Colocar o programa em stand-by durante 1 segundo (time.Second * 3 = 3 segundos) - estilo thread.sleep()
	}
	// ctrl+alt+m => stop running the code (code runner)

	/*
		Vamos ver o foreach no capitulo do array, slice e map
	*/
}
