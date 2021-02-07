package main

import "fmt"

func main() {
	p1 := Ponto{2.0, 2.0}
	p2 := Ponto{2.0, 4.0}

	fmt.Println(catetos(p1, p2)) //Conseguimos usar catetos porque ambos os arquivos são do package main (mesmo package)
	fmt.Println(Distance(p1, p2))
	//PARA EXECUTAR ESTE CÓDIGO TEM DE SER PELO TERMINAL NAO OUTPUT COM CODE RUNNER
}
