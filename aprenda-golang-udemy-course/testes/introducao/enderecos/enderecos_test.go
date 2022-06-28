package enderecos

import "testing"

// TestXxxxxxx
func TestTipoEndereco(t *testing.T) {
	enderecoParaTeste := "Avenida Paulista"

	TipoDeEnderecoEsperado := "Avenida"

	tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

	if tipoDeEnderecoRecebido != TipoDeEnderecoEsperado {
		t.Errorf("O tipo recebido é diferente do esperado ! Esperava %s e recebeu %s",
			TipoDeEnderecoEsperado, tipoDeEnderecoRecebido)
	}
}
