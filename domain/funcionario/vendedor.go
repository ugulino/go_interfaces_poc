package funcionario

import (
	"varejo.org.br/domain/comissao"
	"varejo.org.br/domain/venda"
)

type Vendedor struct {
	Pessoa   Pessoa
	Gerente  Gerente
	Comissao comissao.Comissao
	Venda    venda.Venda
}

func (v *Vendedor) CalculaComissao() (float64, error) {
	var venda = v.Venda
	if venda.Status == "Faturada" {
		v.Comissao.Valor =
			(v.Comissao.Percents / 100) * venda.Valor
		return v.Comissao.Valor, nil
	}
	return 0.00, nil
}
