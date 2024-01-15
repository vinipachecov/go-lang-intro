package tabela

import "testing"

const msgIndex = "%s (parte %s) = indices esperado (%d) <> encontrado (%d)"

func TestIndex(t *testing.T) {
	teses := struct {
		texto string
		parte string
		esperado int
	}{
		{"Cod34 e show", "Cod3r", 0},
		{"", "", 0},
		{"Opa", "opa", -1},
		{"vinicius", "v", 2}
	}

	for _, teste := range testes {
		t.Logf("Massa: %v", teste)
		atual := strings.Index(teste.texto, teste.parte)

		if atual != teste.esperado {
			t.Errorf(msgIndex, teste.texto, test.parte, teste.esperado, atual)
		}
	}
}
