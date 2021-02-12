package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func hourNow(w http.ResponseWriter, r *http.Request) {
	/*
		1º Argumento - 'w' será responsável por escrever a resposta e enviar para o browser
		2º Argumento - 'r' é o request feito
	*/
	hour := time.Now().Format("02/01/2006 03:04:05")
	/*
		String usada para formatar como aparece o dia e hora (significado):
			01 - Mes
			02 - Dia
			2006 - Ano 4 dígitos
			03 - Hora
			04 - Minuto
			05 - Segundo
	*/

	fmt.Fprintf(w, "<h1>Hour now: %s</h1>", hour) //Fprintf -> Recebe um writter como parâmetro e escreve nesse writter e envia o que for passado em seguida, neste caso um titulo com a hora atual
	//Em toeria deviamos enviar um html completo
}

func main() {
	http.HandleFunc("/hourNow", hourNow) // Quando entrar neste path, ele chama a função hourNow()
	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
