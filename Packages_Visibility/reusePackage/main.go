package main

import (
	"fmt"

	"github.com/bpvcode/area" //Estamos a importar o package
)

func main() {
	fmt.Println(area.AreaCirculo(6.0))    //Acessamos a função através do nome do package 'area'
	fmt.Println(area.AreaRetangulo(2, 3)) //Acessamos a função através do nome do package 'area'
	//Não conseguimos acessar a área do triângulo pois está PRIVATE
}
