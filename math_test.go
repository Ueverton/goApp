package main

func TestSoma(t *testint.T) {
	total := soma(10, 15)

	if total != 30 {
		t.Errorf("Resultado da soma é inválido: Resultado %d. Esperado %d", total, 30)
	}
}
