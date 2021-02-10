package math

import "testing"

const erroPadrao = "Valor esperado %v, mas o resultado encontrado foi %v."

func TestMedia(t *testing.T) {
	valorEsperado := 7.28
	valorReal := Media(7.2, 9.9, 6.1, 5.9)

	if valorReal != valorEsperado {
		t.Errorf(erroPadrao, valorEsperado, valorReal)
	}

}
