package main

import "fmt"

/*
	Funções Variaticas
	São funções que recebem quantidade de parâmetros variáveis (aprovados ...)
*/
func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de aprovados")
	for i, valor := range aprovados {
		fmt.Printf("%d) %s \n", i+1, valor)
	}
}

func main() {

	aprovadosFinal := []string{"Bruno", "Rui", "João"} // Criar e guardar um slice de aprovados (Lista de aprovados)

	imprimirAprovados(aprovadosFinal...) // Passar slice como argumento de função variática

	//NÃO É POSSÍVEL PASSAR ARRAYS PARA ARGUMENTO DA FUNÇÃO VARIÁTICA

}
