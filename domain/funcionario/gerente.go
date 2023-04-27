package funcionario

import (
	"varejo.org.br/domain/comissao"
	"varejo.org.br/domain/venda"
)

const Faturada = "Faturada"

type Gerente struct {
	Pessoa   Pessoa
	Venda    venda.Venda
	Comissao comissao.Comissao
}

func (g *Gerente) CalculaComissao() (float64, error) {
	var venda = g.Venda
	var comissao = 0.00
	if venda.Status == Faturada {
		comissao = (g.Comissao.Percents / 100) * venda.Valor * 1.50
		g.Comissao.Valor = comissao
	}
	return comissao, nil
}
