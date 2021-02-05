package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatório() int {
	s := rand.NewSource(time.Now().UnixNano()) //rand = random  -> Vai pegar na data atual, no nanosegundo, e passa como argumento para gerar uma nova source
	r := rand.New(s)                           // Gera um novo numero aleatório mediante a source que foi passada
	return r.Intn(10)                          // Gera um numero aleatório de 0 a 10(argumento passado) a partir da source passada
}

func main() {
	if i := numeroAleatório(); i > 5 {
		fmt.Println("Ganhou")
	} else {
		fmt.Print("Perdeu")
	}
}
