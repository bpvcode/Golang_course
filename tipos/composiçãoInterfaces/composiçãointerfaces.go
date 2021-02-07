package main

import "fmt"

type sportMode interface {
	turnTurboOn()
}

type luxury interface {
	visualImprove()
}

//Interface que é a composição das duas (Desportivo e Luxuoso)
type luxurySportMode interface {
	sportMode
	luxury
	/*
		Podemos colocar mais métodos
		Ou seja, um desportivo luxuoso tem de implementar tanto os métodos da interface sportMode como os métodos da interface luxury, como os outros métodos declarados nesta interface luxurySportMode.
	*/
}

type bmw struct {
}

func (b bmw) turnTurboOn() {
	fmt.Println("Sou um carro desportivo porque implemento o método turnTurboOn()")
}

func (b bmw) visualImprove() {
	fmt.Println("Sou um carro luxuoso porque implemento o método visualImprove()")
}

/*
	CONCLUSÃO:
	Neste caso construimos uma estrutura que pode ser usada como sendo um desportivo, um luxuoso ou um desportivo luxuoso.
	Isto acontece porque bmw implementa todos os métodos das 3 interfaces, pelo que "implementa" as 3 interfaces, é do tipo das 3 interfaces
	Se a interface luxurySportMode tivesse mais 1 método ao qual bmw não dava implementação. Então ela só respeitaria as 2 primeiras interfaces e não a terceira
*/

func main() {
	bmw1 := bmw{}
	bmw1.turnTurboOn()
	bmw1.visualImprove()

	//OUTRA FORMA
	var bmw2 luxurySportMode = bmw{}
	bmw2.turnTurboOn()
	bmw2.visualImprove()

}
