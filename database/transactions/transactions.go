package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // Import forçado - Não vamos explicitamente usar este import, dai começar com '_'
)

func main() {
	db, err := sql.Open("mysql", "root:password@/udemy_Golang_1")
	if err != nil {
		log.Fatal(err) //Faz um log do erro
		panic(err)
	}
	defer db.Close() //defer - será a ultima ação a executar dentro do método

	// ================================ TRANSACTION =================================
	//OU ACONTECE A AÇÃO ATÉ AO FIM OU ENTÃO NADA ACONTECE

	transaction, _ := db.Begin()
	//Begin() - Inicializa uma transaction e retorna um objeto (neste caso o 'transaction') e é nesse objeto onde serão aplicados os comandos mysql
	stmt, _ := transaction.Prepare("insert into users(id,name) values(?,?)")
	/*
		Prepare() - creates a prepared statement for later queries or executions
		Prepare está a ser aplicado ao objeto 'transaction' e não ao db
	*/

	stmt.Exec(1000, "Edu")
	stmt.Exec(1001, "Diana")

	resp, _ := stmt.Exec(1002, "Ze")
	id, _ := resp.LastInsertId()
	rows, _ := resp.RowsAffected()
	fmt.Println("ID do last insert:", id)
	fmt.Println("Quantidade de linhas afetadas:", rows)

	// INSERT ID DUPLICADO (ERRO) - se método não estiver comentado gera erro e faz Rollback(), se tiver comentado faz Commit()
	/*
		Exemplo:
		Queremos inserir um grupo inteiro numa plataforma de chat, ou cria o grupo com todos ou nao cria com nenhum
		Estamos a tentar inserir o Edu a Diana o Zé e a Monica
		A mónica dá erro porque tem o id duplicado (id 1 já existe na base de dados)
	*/
	/*
			_, err = stmt.Exec(1, "Monica")

			if err != nil {
			transaction.Rollback() //Se acontecer um erro, faz roll back
			log.Fatal(err)         // Faz um log com o erro
		}
	*/

	transaction.Commit() //Se nada de errado acontecer, dá commit de todas as alterações
}
