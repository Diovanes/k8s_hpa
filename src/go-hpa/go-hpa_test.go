package main

import "testing"

func TestSoma(t *testing.T) {
	resultado := executaCalculo()
	if resultado != "Code.education Rocks!" {
		t.Errorf("Resultado esperado: <b>code</b>, obtido: %s", resultado)
	}
}
