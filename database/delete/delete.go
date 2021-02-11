package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" // Import forçado - Não vamos explicitamente usar este import, dai começar com '_'
)

func main() {

	db, err := sql.Open("mysql", "root:password@/udemy_Golang_1")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer db.Close()

	// ================================ DELETE =================================
	stmt, _ := db.Prepare("delete from users where id=?")
	stmt.Exec(78878)

}
