package arquitetura

import (
	"fmt"
	"runtime"
	"testing"
)

func TestDependente(t *testing.T) {
	t.Parallel()                   // Executa o testes de forma paralela otimizando o tempo final de execução de todos os testes em simultâneo
	if runtime.GOARCH == "amd64" { //arquitetura do processador 64bits (Forma de verificar qual a arquitetura do processador para poder correr ou não correr o teste em cada computador, mediante a sua arquitetura)
		fmt.Printf("ARQUITETURA: %v \n", runtime.GOARCH)
		t.Skip("Não funciona em arquitetura amd64") //Neste caso, deixa passar o teste porque nao funciona neste tipo de arquitetura
	}
	//... Real implementação do teste (por exemplo para amd32) - condição que tem que passar
	t.Fail() //se não corresponder a nenhum dos if's passados em '...' , falha o teste(proxima linha)
}
