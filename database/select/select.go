package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // Import forçado - Não vamos explicitamente usar este import, dai começar com '_'
)

type user struct {
	id   int
	name string
}

func main() {
	db, err := sql.Open("mysql", "root:password@/udemy_Golang_1")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer db.Close()

	// ================================ SELECT =================================
	rows, _ := db.Query("select id, name from users where id > ?", 3)
	defer rows.Close()

	for rows.Next() {
		var user user                   //Cria um objeto a cada iteração
		rows.Scan(&user.id, &user.name) //Recebe o valor do objeto a cada iteração
		fmt.Println(user)

	}
}
