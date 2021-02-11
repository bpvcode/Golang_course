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

	// ================================ UPDATE =================================
	stmt, _ := db.Prepare("update users set name = ? where id = ?") // '?' usar sempre assim é mais seguro - Não fazer concatenação de strings por caus de SQL injection
	stmt.Exec("Ermelinda da Silva Anacoreta", 2)
	stmt.Exec("Antonieta da Silva Anacoreta", 3)
}
