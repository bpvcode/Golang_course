package main

//Getters and Setters in GO
import (
	"fmt"
	"strings"
)

type pessoa struct {
	firstname string
	lastname  string
}

/*
	=== Getter ===
	Referenciado à estrutura (Devolve uma cópia do valor que está atribuido)
*/
func (p pessoa) getCompleteName() string {
	return p.firstname + " " + p.lastname
}

/*
	=== Setter ===
	Para setar o valor de uma variável/Estrutura, temos de ter acesso à referência de memória e alterar lá o valor, por isso usamos ponteiros
*/
func (p *pessoa) setCompleteName(completeName string) {
	partes := strings.Split(completeName, " ") // Estamos a dizer que vai separar pelo espaço " ", ou seja o que o cliente escrever separa por espaços
	p.firstname = partes[0]                    //Atribuição do nome passado em argumento, atualizando o valor da propriedade 'firstname' da estrutura pessoa
	p.lastname = partes[1]                     //Atribuição do nome passado em argumento, atualizando o valor da propriedade 'lastname' da estrutura pessoa
}

func main() {
	pessoa1 := pessoa{"Bruno", "Vilar"}

	fmt.Println(pessoa1.getCompleteName())

	pessoa1.setCompleteName("Rui Acácio")
	fmt.Println(pessoa1.getCompleteName())
}
