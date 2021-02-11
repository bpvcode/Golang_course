package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // Import forçado - Não vamos explicitamente usar este import, dai começar com '_'
)

func main() {
	db, err := sql.Open("mysql", "root:password@/udemy_Golang_1")
	/*
		1º argumento - informar qual o tipo de sistema de gestão de base de dados a usar
		2º argumento - Neste caso vamos diretamente para a base de dados (udemy_Golang_1) criada no exercicio anterior (estrutura.go)
	*/
	if err != nil {
		panic(err)
	}
	defer db.Close()
	/*
		defer - será a ultima ação a executar dentro do método
		Ou seja, antes do programa terminar vamos encerrar a ligação à base de dados
	*/

	// ================================ INSERT =================================

	statement, err := db.Prepare("insert into users(name) values(?)")
	/*
		Prepare() - creates a prepared statement for later queries or executions
		O '?' significa que quando executarmos o statement, podemos passar a string em argumento que preenche aquele valor
	*/
	if err != nil {
		panic(err)
	}
	statement.Exec("Bruno") //Executar o statement
	statement.Exec("Miguel")
	/*
		NOTA - Função Exec() retorna:
		* Id da pessoa que foi inserida
		* Quantidade de linhas afetadas depois de executar o comando
	*/
	res, _ := statement.Exec("Pedro")

	id, _ := res.LastInsertId()
	fmt.Println("ID do last insert:", id)

	rows, _ := res.RowsAffected()
	fmt.Println("Quantidade de linhas afetadas:", rows)
}
