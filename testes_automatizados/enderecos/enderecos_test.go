// TESTE DE UNIDADE
package enderecos_test

import (
	. "introducao-testes/enderecos"
	"testing"
)

type cenarioDeTeste struct {
	endercoInserido string
	retornoEsperado string
}

func TestTipoDeEndereco(t *testing.T) {

	t.Parallel()
	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		//{"Avenida", "Avinida"},
		{"Rodovia dos imigrantes", "Rodovia"},
		//{"Praça das Rosas", "Tipo invalido"},
		{"Estrada travessia", "Estrada"},
		{"RUA DOS BOBOS", "Rua"},
		{"AVENIDA REBOLÇAS", "Avenida"},
		//{"", "Tipo invalido"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.endercoInserido)

		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
				tipoDeEnderecoRecebido,
				cenario.retornoEsperado,
			)
		}
	}

	// enderecoParaTeste := "Avenida Pedra"
	// tipoDeEnderecoEsperado := "Avenida"
	// tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

	//if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado{
	//	t.Errorf("O tipo recebido é diferente do esperado! esperava %s e redebeu %s",
	//tipoDeEnderecoEsperado, tipoDeEnderecoRecebido)

	//}
}

func TestQualquer(t *testing.T) {
	t.Parallel()
	if 1 > 2 {
		t.Errorf("teste querou!")
	}
}
