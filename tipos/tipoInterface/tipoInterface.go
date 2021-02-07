package main

import "fmt"

type course struct {
	name string
}

func main() {
	var something interface{}
	fmt.Println(something) // Something é do tipo da interface (neste caso vazia - <nil>)

	something = 3 //Reatribuir à váriavel something
	fmt.Println(something)

	type dynamic interface{}
	var something1 dynamic = "Hola"
	fmt.Println(something1)

	something1 = true
	fmt.Println(something1)

	something1 = course{"Golang udemy course"}
	fmt.Println(something1)

	/*
		CONCLUSÃO:
		* Interface é do tipo genérico, pelo que pode receber diferentes tipos, como vimos nos exemplos a cima:
		  vazio , int, string, boolean, struct
		* dynamic, como é uma struct do tipo da interface, também pode receber vários tipo
	*/
}
