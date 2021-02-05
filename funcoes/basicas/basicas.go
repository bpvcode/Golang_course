package main

import "fmt"

//Não recebe parâmetros nem retorna
func f1() {
	fmt.Println("Primeira função")
}

//Recebe 2 parâmetros de tipos distintos não retorna nada
func f2(p1 string, p2 int) {
	fmt.Printf("Segunda função: %s , %d \n", p1, p2)
}

//Não recebe parâmetros retorna um tipo de dado
func f3() string {
	return "Terceira função"
}

//Recebe 2 parâmetros do mesmo tipo e retorna um tipo de dado
func f4(p1, p2 string) string { // -> Os dois parametros são do tipo string
	return fmt.Sprintf("Quarta função: %s , %s \n", p1, p2)
}

//NOTA: Função pura - recebe um conjunto de parâmetros e gera sempre o mesmo resultado para os mesmos parâmetros recebidos. Função pura não influencia nada que esteja fora da função, so trabalha com os parâmetros que recebe

//FUNÇÃO QUE RETORNA 2 VALORES (Podia também receber parâmetros)
func f5() (string, string) {
	return "Primeiro valor", "Segundo Valor"
}

func main() {
	f1()
	f2("Bruno", 1)

	r3, r4 := f3(), f4("Bruno", "Vilar")
	fmt.Printf("Variável r3: %s || Variável r4: %s", r3, r4)
	/*
		SÓ SE PODE FAZER ASSIM PORQUE TANTO f3() COMO f4() RETORNAM UMA STRING
		Forma de guardar o resultado que resulta da invocação das funções, direto nas variáveis
		r3 := f3()
		r4 := f4(.. , ..)
	*/

	r5_1, r5_2 := f5()
	fmt.Printf("Variável r5_1: %s || Variável r5_2: %s", r5_1, r5_2)
	/*
		Como a função f5() retorna 2 valores do tipo string
		Temos de tratar de guardar os dois valores
		Podemos ignorar qualquer um utilizando '_'
		Ex:  _, r5_2 := f(5)  -> Ignora o primeiro retorno
		Ex:  r5_1, _ := f(5)  -> Ignora o segundo retorno
	*/
}
