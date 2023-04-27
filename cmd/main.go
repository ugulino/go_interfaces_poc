package main

import (
	"fmt"
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

	fmt.Println("Comissão gerente:", venda.Execute(&john))
	fmt.Println("Comissão vendedor:", venda.Execute(&mary))
}
