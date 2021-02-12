package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/user/", UserHandler) //Estmos a importar a func UserHandler() que está no ficheiro client.go (Sempre que for chamada o path .../user/ -> ele chama a função UserHandler() )
	log.Println("Executando")
	log.Fatal(http.ListenAndServe(":3000", nil)) //Port
}
