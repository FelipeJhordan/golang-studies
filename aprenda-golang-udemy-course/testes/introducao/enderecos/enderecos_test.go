package enderecos_test

import (
	. "introducao-testes/enderecos"
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

// TestXxxxxxx
func TestTipoEndereco(t *testing.T) {

	t.Parallel()
	cenarioDeTeste := []cenarioDeTeste{
		{
			enderecoInserido: "Rua ABC",
			retornoEsperado:  "Rua",
		},
		{
			enderecoInserido: "Avenida Brasil",
			retornoEsperado:  "Avenida",
		},
		{
			enderecoInserido: "Praça Das Rosas",
			retornoEsperado:  "Tipo não especificado",
		},
		{
			enderecoInserido: "Rodovia 150",
			retornoEsperado:  "Rodovia",
		},
		{
			enderecoInserido: "Estrada da Vila",
			retornoEsperado:  "Estrada",
		},
		{
			enderecoInserido: "RUA DOS BOBOS",
			retornoEsperado:  "Rua",
		},
		{
			enderecoInserido: "",
			retornoEsperado:  "Tipo não especificado",
		},
	}

	for _, cenario := range cenarioDeTeste {
		TipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if TipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
				TipoDeEnderecoRecebido, cenario.retornoEsperado)
		}
	}
}

func TestQualquer(t *testing.T) {

	t.Parallel()
	if 1 > 2 {
		t.Error("Teste quebrou")
	}
}
