package main

import "fmt"

func main() {

	//## QUEBRAS DE LINHA ##
	fmt.Print("Mesma ")
	fmt.Print("Linha ")

	fmt.Println("*Nova linha")
	fmt.Println("Anterior começou na mesma linha por causa do Print(). Este ja começa na linha abaixo porque o anterior é Println()")

	//## TIPOS ##
	x := 3.14156 // tipo -> float64

	//fmt.Println("O valor de x é " + x) -> nao funciona a concatenação pois não é uma string mas sim um float64...

	xs := fmt.Sprint(x)                 // Sprint() -> Retorna x passado em argumento mas em formato string
	fmt.Println("O valor de x é " + xs) // -> Ja funciona

	//Exemplo para concatenar valores que nao seja em String:
	fmt.Println("O valor de x é", x)
	//Exemplo para concatenar valores formatados:
	fmt.Printf("O valor de x é %f \n", x)

	// ## PLACEHOLDERS ## -> %...

	/*
		Imprimir valor no formato ( \n -> quebra de linha ):
		%f - float
		%s - string
		%d - inteiro
		%t - boolean
		%v - Serve para vários tipos diferentes - TIPO GENÉRICO
		%T - imprime o tipo do valor que for passado
	*/

	fmt.Printf("O valor de x é %.2f \n", x) // Imprimir formatado a duas casa decimais .2
	a := 1
	b := 1.999 // Arredonda pela ordem natural maior ou igual a 0.5 arredonda para cima
	c := true
	d := "hey"
	fmt.Printf("%d %f %.1f %t %s \n", a, b, b, c, d)

	fmt.Printf("%v %v %.1v %v %v", a, b, b, c, d)

}
