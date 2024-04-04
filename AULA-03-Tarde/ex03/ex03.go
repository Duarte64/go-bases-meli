/*Uma empresa nacional é responsável pela venda de produtos, serviços e manutenção.

Para isso, eles precisam realizar um programa que seja responsável por calcular o preço total dos Produtos, Serviços e Manutenção. Devido à forte demanda e para otimizar a velocidade, eles exigem que o cálculo da soma seja realizado em paralelo por meio de 3 go routines.


Precisamos de 3 estruturas:

Produtos: nome, preço, quantidade.
Serviços: nome, preço, minutos trabalhados.
Manutenção: nome, preço.

Precisamos de 3 funções:

Somar Produtos: recebe um array de produto e devolve o preço total (preço * quantidade).
Somar Serviços: recebe uma array de serviço e devolve o preço total (preço * média hora trabalhada, se ele não vier trabalhar por 30 minutos, ele será cobrado como se tivesse trabalhado meia hora).
Somar Manutenção: recebe um array de manutenção e devolve o preço total.

Os 3 devem ser executados concomitantemente e ao final o valor final deve ser mostrado na tela (somando o total dos 3).*/

package ex03

import "fmt"

type ProdutoEX3 struct {
	Nome       string
	Preco      float64
	Quantidade int
}

type ServicoEX3 struct {
	Nome               string
	Preco              float64
	MinutosTrabalhados int
}

type ManutencaoEX3 struct {
	Nome  string
	Preco float64
}

func somarProdutos(produtos []ProdutoEX3, canal chan float64) {
	var total float64
	for _, produto := range produtos {
		total += produto.Preco * float64(produto.Quantidade)
	}
	canal <- total
}

func somarServicos(servicos []ServicoEX3, canal chan float64) {
	var total float64
	for _, servico := range servicos {
		if servico.MinutosTrabalhados < 30 {
			total += servico.Preco * 0.5
		} else {
			total += servico.Preco * (float64(servico.MinutosTrabalhados / 60))
		}
	}
	canal <- total
}

func somarManutencao(manutencoes []ManutencaoEX3, canal chan float64) {
	var total float64
	for _, manutencao := range manutencoes {
		total += manutencao.Preco
	}
	canal <- total
}

func ExerciciosTres() {
	fmt.Println("\nEXERCÍCIO 3")
	produtos := []ProdutoEX3{
		{Nome: "Produto 1", Preco: 10.0, Quantidade: 5},
		{Nome: "Produto 2", Preco: 20.0, Quantidade: 10},
	}

	manutencoes := []ManutencaoEX3{
		{Nome: "Manutencao 1", Preco: 100.0},
		{Nome: "Manutencao 2", Preco: 200.0},
	}

	servicos := []ServicoEX3{
		{Nome: "Servico 1", Preco: 1000.0, MinutosTrabalhados: 120},
		{Nome: "Servico 2", Preco: 2000.0, MinutosTrabalhados: 20},
	}

	canal := make(chan float64)

	go somarProdutos(produtos, canal)
	go somarManutencao(manutencoes, canal)
	go somarServicos(servicos, canal)

	s1, s2, s3 := <-canal, <-canal, <-canal

	fmt.Printf("Valor total: %.2f\n", s1+s2+s3)
}
