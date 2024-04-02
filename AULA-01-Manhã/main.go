package main

import "fmt"

func main() {
	exercicioUm()
	exercicioDois()
	exercicioTres()
	exercicioQuatro()
}

func exercicioUm() {
	fmt.Printf("EXERCÍCIO 1\n")
	var nome string = "Gabriel"
	var idade int = 23
	fmt.Printf("Olá, meu nome é %s, tenho %d anos de idade.\n", nome, idade)
}

func exercicioDois() {
	fmt.Printf("\nEXERCÍCIO 2\n")
	var temperatura int = 27
	var umidade float32 = 0.35
	var pressao int = 1016
	fmt.Printf("Temperatura: %d°C\nUmidade: %.0f\nPressão: %d mb\n", temperatura, umidade*100, pressao)
}

func exercicioTres() {
	fmt.Printf("\nEXERCÍCIO 3\n")
	/*
		CRIADAS DE MANEIRA ERRADA 👇🏼
			var 1nome string 				-> Não pode começar com número.
			var int idade 					-> Ordem está errada, primeiro o nome da variável e depois o tipo.
			1sobrenome := 6 				-> Não pode começar com número, e o nome da variável não é semântico, sobrenome está como um inteiro ao invés de texto.
			var licenca_para_dirigir = true -> Deveria utilizar camelCase
			var estatuda da pessoa int		-> As palavras não podem estar separadas, devem estar todas juntas em camelCase

		CRIADAS DE MANEIRA CORRETA 👇🏼
			var sobrenome string
			quantidadeDeFilhos := 2
			var licenca_para_dirigir = true -> Não está errada, o GO aceita, porém não segue as convenções.
	*/

	//LISTA COM TODAS CORRETAS 👇🏼
	var nome string
	var idade int
	sobrenome := "texto"
	var licencaParaDirigir = true
	var estaturaDaPessoa int
	quantidadeDeFilhos := 2
	fmt.Printf("Nome: %s\nSobrenome: %s\nIdade: %d\nLicença para dirigir: %t\nEstatura: %d\nQuabtidade de filhos: %d\n", nome, sobrenome, idade, licencaParaDirigir, estaturaDaPessoa, quantidadeDeFilhos)
}

func exercicioQuatro() {
	fmt.Printf("\nEXERCÍCIO 4\n")
	permissao := false
	var sobrenome string = "Silva"
	var idade int = 25
	var salario float64 = 4585.90
	var nome string = "Fellipe"
	fmt.Printf("Sobrenome: %s\nIdade: %d\nPermissão: %t\nSalário: %.2f\nNome: %s\n", sobrenome, idade, permissao, salario, nome)
}
