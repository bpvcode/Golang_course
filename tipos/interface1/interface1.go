package main

import (
	"fmt"
)

type printable interface {
	toString() string
}

/*
Declaramos a interface printable. Uma estrutura para ser printable tem de dar implementação a este método toString()
O compilador automaticamente identifica como "implements interface"- em java
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

func printer(x printable) { //Função que recebe como parâmetro qualquer estrutura que respeite a interface printable (Neste caso, que implemente o método toString() )
	fmt.Println(x.toString())
}

func main() {
	var something printable = person{ //something é do tipo person, mas do tipo da interface - programar para a interface
		name:     "Bruno",
		lastname: "Vilar",
	}
	fmt.Println(something.toString())
	fmt.Println(something)
	printer(something)

	another := product{
		name:  "Cleaner",
		price: 1.89,
	}
	fmt.Println(another.toString())
	fmt.Println(another)
	printer(another)

	//Reatribuindo o valor à variável something
	//something tanto pode receber product como person (como esta no 1º exemplo função main) isto porque ambos são do mesmo tipo, do tipo da interface. -> POLIMORFISMO (something tanto tem a forma de product como person - multiplas formas)
	something = product{
		name:  "Shampoo",
		price: 1.5,
	}
	fmt.Println(something.toString())
	fmt.Println(something)
	printer(something)

	printer(product{"Drink", 1.5}) //Outra forma de fazer

}
