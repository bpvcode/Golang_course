package main

import (
	"fmt"
	"math"
)

func main() {

	// 2 Formas de criar, no caso do var, o tipo da variável é inferido pelo valor que lhe é passado (neste caso float64)
	const PI float64 = 3.1415
	var raio = 3.2

	//Forma reduzida - MELHOR OPÇÃO - declara e inicializa variável
	area := PI * math.Pow(raio, 2)

	fmt.Println("A área da circunferência é", area)

	//Declarar variáveis em grupo/bloco
	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)
	fmt.Println(a, b, c, d)

	// Declarar variáveis na mesma linha -> var e = true && var f = false
	var e, f bool = true, false
	fmt.Println(e, f)

	//Declarar e inicializar as variáveis
	g, h, i := 2, true, "hey"
	fmt.Println(g, h, i)
}
