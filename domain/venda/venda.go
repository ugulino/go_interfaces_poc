package venda

type Venda struct {
	Valor    float64
	Status   string
	Calculos CalculosVenda
}

type CalculosVenda interface {
	CalculaComissao() (float64, error)
}

func (v *Venda) Execute(c CalculosVenda) (float64, error) {
	var comissao float64 = 0.00
	var err error = nil
	comissao, err = c.CalculaComissao()

	if err != nil {
		return 0.00, err
	}

	return comissao, nil
}
