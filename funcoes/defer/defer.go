package main

import "fmt"

/*
	Defer -
	* colocamos defer quando queremos atrasar um bocado de código para a execução do mesmo só ocorrer na linha antes do return da função
	* Imaginando que temos vários if e else, e mesmo que entre em qualquer uma dessas condições, temos sempre de executar um bocado de código antes do return de cada um
	* Para isso colocamos no defer esse bocado de código, que só irá ser executado antes do return da função
	* Exemplo prático: Terminar a conexão existente entre cliente e server
*/
func deferExemplo(numero int) int {
	defer fmt.Println("Hello")
	if numero > 5000 {
		fmt.Println("MAIOR:", numero)
		//"Hello" - vindo do defer
		return numero
	} else if numero < 5000 {
		fmt.Println("MENOR:", numero)
		//"Hello" - vindo do defer
		return numero
	} else {
		fmt.Println("IGUAL:", numero)
		//"Hello" - vindo do defer
		return numero
	}

}

func main() {
	deferExemplo(5000)
	deferExemplo(5001)
	deferExemplo(4999)
}
