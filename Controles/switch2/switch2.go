package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()
	//fmt.Println(t) -> Output: Dia / hora segundos...
	switch { // Igual a ter switch true{} - O primeiro caso do switch a dar verdadeiro ele executa e sai, se nao executa o default
	case t.Hour() < 12:
		fmt.Println("Horas:", t.Hour())
		fmt.Println("Bom dia!")
	case t.Hour() < 18:
		fmt.Println("Horas:", t.Hour())
		fmt.Println("Boa tarde")
	default:
		fmt.Println("Horas:", t.Hour())
		fmt.Println("Boa noite!")
	}

}
