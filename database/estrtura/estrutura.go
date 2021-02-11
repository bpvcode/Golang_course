package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Import forçado - Não vamos explicitamente usar este import, dai começar com '_'
)

//Função que facilita execuções (porque vai ser chamada várias vezes)
func exec(db *sql.DB, sql string) sql.Result {
	/*
		Argumentos:
		* db *sql.DB  -> Recebe uma referência('*') para a conexão estabelecida com o sistema de gestão da base de dados
		* sql string  -> Recebe uma query sql
	*/
	result, err := db.Exec(sql) //O resultado deste método pode ser favorável, executando o comando sql que foi passado pelo argumento do tipo string, ou então pode gerar um erro.
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	db, err := sql.Open("mysql", "root:password@/") //Inicia e guarda uma referência para a conexão estabelecida com o sistema de gestão da base de dados
	/*
		Abrir conexão à database
		1º argumento - string 'mysql' argumento da função Open() -> serve para indiretamente usar a dependência do import forçado a cima, que é o driver de aceso ao banco de dados
		2º argumento - entra diretamenta na "home" do mysql
	*/
	if err != nil {
		panic(err)
	}

	defer db.Close()
	/*
		defer - será a ultima ação a executar dentro do método (neste caso o main) antes dele acabar
		Ou seja, antes do programa terminar queremos terminar a ligação à base de dados
	*/

	exec(db, "create database if not exists udemy_Golang_1") //Cria uma database
	exec(db, "use udemy_Golang_1")                           //Usa a databes
	exec(db, "drop table if exists users")                   //Se existir uma tabela de users, apaga
	exec(db, `create table users(
		id integer auto_increment,
		name varchar(80),
		PRIMARY KEY(id)
	)`) //Cria uma tabela de users com varias colunas

}
