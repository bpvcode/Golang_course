package main

import "fmt"

//NOTA: Go não tem aritmética de ponteiros

func main() {

	i := 1
	/*
		Por inferência o compilador define como uma variável inteira de 64 bits (8 bytes), espaço que ocupa em memória
		É possível ter acesso ao endereço de memória desta variável e guardar esse endereço num ponteiro
		É possível compartilhar o endereço de memória para ser referenciado por várias variáveis
		Ao ter partilhado o espaço de memória com outras referências, sempre que existir uma alteração no endereço de memória, todas as referências terão acesso à alteração
	*/

	var p *int = nil
	/*
		A variável p é um ponteiro por causa da anotação *
		O tipo de retorno do ponteiro será um int
		Ao receber nil, o ponteiro de momento nao aponta para lugar nenhum
	*/

	p = &i //Pegar no endereço de memória da variável i (declarada e inicializada a cima) e atribuir esse valor ao ponteiro p

	fmt.Println("Valor associado ao ponteiro p:", *p) //Processo de desreferenciar (acessar o valor associado ao ponteiro)
	*p++
	fmt.Println("Incrementei anteriormente valor associado ao ponteiro p:", *p)
	fmt.Println("Tanto p como i têm o mesmo valor agora:", *p, i)

	fmt.Println("Endereço de memória guardado no ponteiro p:", p) //valor de p (endereço de memória)
	fmt.Println("Endereço de memória associado à variável i:", &i)
}
