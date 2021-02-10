package main

import (
	"fmt"
	"time"

	"github.com/bpvcode/reusablePackage"
)

func faster(url1, url2, url3 string) string {
	//Estamos a reutilizar o package que temos em '/home/brunovilar/go/src/github.com/bpvcode/reusablePackage'
	//Neste caso a função faster() é uma generator
	channel1 := reusablePackage.Title(url1)
	channel2 := reusablePackage.Title(url2)
	channel3 := reusablePackage.Title(url3)

	//Estrutura de controlo especifica para concorrência - SELECT
	select {
	case t1 := <-channel1:
		return t1
	case t2 := <-channel2:
		return t2
	case t3 := <-channel3:
		return t3
	case <-time.After(time.Second): //Se demorar mais do que 1 segundo
		return "Todos demoraram mais do que 1 segundo"
		/*default:
		return "Ainda sem resposta!"*/
	}

}

func main() {

	faster := faster(
		"http://www.google.com",
		"http://www.facebook.com",
		"http://www.udemy.com",
	)

	fmt.Println(faster)
	//Se descomentar o default: Aparece sempre o default
	//O primeiro dado a chegar entra num dos cases e retorna
}
