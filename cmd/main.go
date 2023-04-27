package main

import (
	"fmt"

	"varejo.org.br/domain/comissao"
	"varejo.org.br/domain/funcionario"
	"varejo.org.br/domain/venda"

)

func main() {
	var venda venda.Venda = venda.Venda{
		Valor:  5000,
		Status: "Faturada",
	}

	var john = funcionario.Gerente{
		Pessoa: funcionario.Pessoa{
			Nome:     "John",
			Sexo:     "Masculino",
			Idade:    35,
			Telefone: "+55 011 99101-9211",
			Email:    "john@gmail.com",
		},
		Comissao: comissao.Comissao{
			Percents: 6,
		},
		Venda: venda,
	}

	var mary = funcionario.Vendedor{
		Pessoa: funcionario.Pessoa{
			Nome:     "Mary",
			Sexo:     "Feminino",
			Idade:    25,
			Telefone: "+55 011 99301-1234",
			Email:    "mary@gmail.com",
		},
		Comissao: comissao.Comissao{
			Percents: 3,
		},
		Venda: venda,
	}

	var comissao_gerente, _ = venda.Execute(&john)
	var comissao_vendedor, _ = venda.Execute(&mary)
	fmt.Println("Comissão gerente:", comissao_gerente)
	fmt.Println("Comissão vendedor:", comissao_vendedor)
}
