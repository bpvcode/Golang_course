package main

import (
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("public")) //Identifica a pasta 'public' a partir do qual este webserver vai ler
	http.Handle("/", fileServer)                      //No home (localhost:3000) - aparece o ficheiro index.html da pasta 'public'

	log.Println("Executando file server...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
