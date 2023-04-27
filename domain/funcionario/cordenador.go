package funcionario

import (
	"varejo.org.br/domain/comissao"
	"varejo.org.br/domain/venda"
)

type Coordenador struct {
	Pessoa   Pessoa
	Gerente  Gerente
	Comissao comissao.Comissao
	Venda    venda.Venda
}
