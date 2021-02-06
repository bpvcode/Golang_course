package main

import (
	"fmt"
)

type printable interface {
	toString() string
}

/*
Declaramos a interface printable. Uma estrutura para ser printable tem de implementar este método toString()
*/

type person struct {
	name     string
	lastname string
}

type product struct {
	name  string
	price float64
}

//Interfaces são implementadas implicitamente

func (p person) toString() string {
	return p.name + " " + p.lastname
}

func (pr product) toString() string {
	return fmt.Sprintf("%s - %.2f€", pr.name, pr.price)
}

func printer(x printable) { //Função que recebe como parâmetro qualquer estrutura que respeite a interface printable
	fmt.Println(x.toString())
}

func main() {
	var something printable = person{ //something é do tipo person, mas do tipo da interface - programar para a interface
		name:     "Bruno",
		lastname: "Vilar",
	}
	fmt.Println(something.toString())
	fmt.Println(something)

	another := product{
		name:  "Cleaner",
		price: 1.89,
	}
	fmt.Println(another.toString())
	fmt.Println(another)

	//something tanto pode receber product como person (como esta no 1º exemplo) isto porque ambos são do mesmo tipo, do tipo da interface
	something = product{
		name:  "Shampoo",
		price: 1.5,
	}
	fmt.Println(something.toString())
	fmt.Println(something)

}
