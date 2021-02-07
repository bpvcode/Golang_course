package main

import (
	"fmt"

	"github.com/bpvcode/goareatest"
	"github.com/bpvcode/goareatest/anotherPackage"
)

/*
	Para ir buscar esta biblioteca, que estava num repositório do github (na minha conta), e para a poder usar utilizamos o seguinte comando:
	go get -u github.com/bpvcode/goareatest
	Com este comando, foi criado na pasta ~/go/src um directory goareatest (~/go/src/goareatest) com o conteudo do repositório do github
	Também foi criado um ficheiro executavel do conteudo do repositório na pasta pkg no caminho -> /home/brunovilar/go/pkg/linux_amd64/github.com/bpvcode
*/

func main() {
	fmt.Println(goareatest.AreaCirculo(2))

	anotherPackage.Hola()
	//Este package é outro package que foi criado dentro de goareatest

}
