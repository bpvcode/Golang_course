package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

//channel <-   * Inserir dados no canal
//<-channel    * Ler dados do canal

/*
	GENERATOR:
	* Função que encapsula a chamada de um go routine (neste caso, função title() )
	* Objetivo quem chamar a função nao tem de se preocupar com channels etc...
*/

func title(urls ...string) <-chan string {
	/*
		Recebe uma lista de URLS (Strings) de forma variática (pode receber quantidade de argumentos variável)
		Retorna um channel com a lista dos titulos de cada URL
	*/
	channel := make(chan string) //Função cria o canal

	for _, url := range urls { // For range que itera sobre a lista de URLS que vêm de argumentos (ignoramos o indice '_' em cada iteração, porque só precisamos de trabalhar com os valores associados ao indice)
		go func(url string) { //Função anónima a partir de uma Go routine
			resp, err := http.Get(url) //Fazer pedido de Get ao URL recebido em cada iteração
			if err != nil {
				panic("http.GET - ERRO")
			}
			html, err := ioutil.ReadAll(resp.Body) //Ler o body do conteudo que veio na resposta do pedido get ao URL
			if err != nil {
				panic("http.GET - ERRO")
			}
			r, err := regexp.Compile("<title.*?>(.*?)<\\/title>") //Regexp em que só pega o titulo da página
			if err != nil {
				panic("http.GET - ERRO")
			}
			channel <- r.FindStringSubmatch(string(html))[1] //Conversão para String e aplicação de regex para retornar só o titulo. Por fim, envia os dados para o channel
		}(url) //Como é uma função anónima temos de invocá-la e passar o argumento 'url'. Cada vez que o for executa, lança-se uma go routine para chamar de forma independente a chamada para os vários sites. Sempre que o resultado chega, aplica expressão regex ficando só com o tiutlo e envia o titulo para o canal.
	}
	return channel //A função retorna um channel mesmo antes de fazer o primeiro http.GET , à medida que vao chegando as respostas, vai inserindo os dados nesse mesmo channel
}

func main() {
	title1 := title("https://www.facebook.com", "https://www.google.com") //Função generator title
	title2 := title("https://www.udemy.com", "https://www.youtube.com")

	fmt.Println("Primeiros: '", <-title1, "':'", <-title2, "'")
	fmt.Println("Segundos: '", <-title1, "':'", <-title2, "'")
}
