package strings //Estamos a testar um package built in no go (package strings)

import (
	"strings"
	"testing"
)

const msgPadrão = "%s (parte: %s) - indices: esperado (%d) <> encontrado (%d)"

func TestIndex(t *testing.T) {
	/*
		Utilizando uma '[] struct' conseguimos criar uma tabela dataset, para testar vários cenários de testes diferentes
		Um método testa vários cenários diferentes
	*/

	testes := []struct {
		texto         string
		parte         string
		valorEsperado int
	}{
		{"Bruno Vilar", "Bruno", 0},
		{"", "", 0},
		{"Bruno Vilar", "r", 1},
		{"Bruno Vilar", "Joao", -1},
	}

	for i, teste := range testes {
		t.Logf("Teste %d: %v", i, teste)

		atual := strings.Index(teste.texto, teste.parte)
		//func Index -> Index returns the index of the first instance of 'teste.parte' in 'teste.texto', or -1 if 'teste.parte' is not present in 'teste.texto'.

		if teste.valorEsperado != atual {
			t.Errorf(msgPadrão, teste.texto, teste.parte, teste.valorEsperado, atual)
		}
	}
}
