package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql" // Import forçado - Não vamos explicitamente usar este import, dai começar com '_'
)

//User - model
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// UserHandler - analisa o request e delega para a função adequada
func UserHandler(w http.ResponseWriter, r *http.Request) {
	idString := strings.TrimPrefix(r.URL.Path, "/user/")
	/*
		Exemplo:
		https://localhost:3000/user/1 - Ou seja estariamos a aceder aos dados do user com id=1
		idString - vai ao url a cima, retira tudo o que está antes de '/user/', e também retira '/user/', pelo que irá guardar neste caso só o numero 1 (id)
		Nota -retorna 1 em string
	*/

	idInt, _ := strconv.Atoi(idString) // Converter string to int
	fmt.Println("idInt:", idInt)

	switch {
	case r.Method == "GET" && idInt > 0: // Tipo de método recebido == "GET"
		userPorID(w, r, idInt)
	case r.Method == "GET":
		userTodos(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "SORRY, NOT FOUND :(")
	}
	/*
		Switch conclusão:
		se o path for :'.../user/id'  (Ou seja tem id)-> entra no primeiro case
		se o path for :'.../user/' (Ou seja não tem id)-> entra no segundo case
	*/
}

func userPorID(w http.ResponseWriter, r *http.Request, id int) {
	db, err := sql.Open("mysql", "root:password@/udemy_Golang_1") //URL para database ao deveria ser colocado hard coded como esta
	if err != nil {
		log.Fatal(err)
		//Nao enviar informações de erro para user final para nao partilhar estrutura da aplicação
	}
	defer db.Close()

	var user User
	db.QueryRow("select id, name from users where id = ?", id).Scan(&user.ID, &user.Name)
	//QueryRow - vamos pegar só uma linha, neste caso a linha em que o id for igual ao qe for passado no argumento da função. Retorna um struct do tipo User

	json, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
	/*
		Convertemos o result set que veio do QueryRow() method, ou seja, ele retornou uma struct User, convertemos a struct User para json
		Informamos que o conteudo do Header() da resposta será no tipo json
		Passamos o json como resposta da requisição
	*/
}

func userTodos(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password@/udemy_Golang_1")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, _ := db.Query("select id, name from users")
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		rows.Scan(&user.ID, &user.Name)
		users = append(users, user)
	}

	json, _ := json.Marshal(users)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}
