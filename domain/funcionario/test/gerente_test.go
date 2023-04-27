package gerente

import (
	"testing"

	"varejo.org.br/domain/comissao"
	"varejo.org.br/domain/funcionario"
	"varejo.org.br/domain/venda"
)

func Test_cacula_comissao_venda_(t *testing.T) {
	var venda venda.Venda = venda.Venda{
		Valor:  100,
		Status: "Faturada",
	}

	var john = funcionario.Gerente{
		Pessoa: funcionario.Pessoa{
			Nome: "John",
		},
		Comissao: comissao.Comissao{
			Percents: 10,
		},
		Venda: venda,
	}

	var valor_comissao_expected = 0.00
	var valor_comissao_calculated = venda.Execute(&john)

	if valor_comissao_expected != valor_comissao_calculated {
		t.Error("Expected valor_comissao_calculated to be 0.00")
	}
}

func Test_cacula_comissao_venda_cancelada(t *testing.T) {
	var venda venda.Venda = venda.Venda{
		Valor:  100,
		Status: "Cancelada",
	}

	var john = funcionario.Gerente{
		Pessoa: funcionario.Pessoa{
			Nome: "John",
		},
		Comissao: comissao.Comissao{
			Percents: 10,
		},
		Venda: venda,
	}

	var valor_comissao_expected = 15.00
	var valor_comissao_calculated = venda.Execute(&john)

	if valor_comissao_expected != valor_comissao_calculated {
		t.Error("Expected valor_comissao_calculated to be 15.00")
	}
}
