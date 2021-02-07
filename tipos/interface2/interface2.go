package main

import "fmt"

type sportMode interface {
	turnTurboOn()
}

type ferrari struct {
	name  string
	speed float64
	turbo bool
}

func (f *ferrari) turnTurboOn() { //O * é necessário para podermos alterar o valor da instância de ferrari. Se não, o que faz é criar uma cópia
	f.turbo = true
}

func main() {

	ferrari1 := ferrari{
		name:  "F40",
		speed: 0,
		//turbo -> valor default: false
	}
	fmt.Println("Turbo ON:", ferrari1.turbo)

	ferrari1.turnTurboOn()
	fmt.Println("Turbo ON:", ferrari1.turbo)

	//OUTRA FORMA - IMPORTANTE (Melhor fazer como fizemos em ferrari1)
	var ferrari2 sportMode = &ferrari{ //Tem de ter & para podermos aceder ao valor
		name:  "F50",
		speed: 0,
	}
	fmt.Println(ferrari2)

	ferrari2.turnTurboOn()
	fmt.Println(ferrari2)

}
