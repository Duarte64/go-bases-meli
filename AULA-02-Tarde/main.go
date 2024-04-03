package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	exercicioUm()
	exercicioDois()
}

// *********************
// INICIO EXERCÍCIO 1 👇🏼
// *********************

type Aluno struct {
	Nome           string
	Sobrenome      string
	RG             string
	DataDeAdmissao time.Time
}

func (a Aluno) detalhes() {
	aluno := reflect.TypeOf(a)
	for i := 0; i < aluno.NumField(); i++ {
		field := aluno.Field(i)
		fmt.Printf("%s: %s\n", field.Name, reflect.ValueOf(a).Field(i))
	}
}
func exercicioUm() {
	fmt.Println("EXERCÍCIO 1")
	aluno := Aluno{
		Nome:           "Gabriel",
		Sobrenome:      "Duarte",
		RG:             "504267558",
		DataDeAdmissao: time.Now(),
	}
	aluno.detalhes()
}

// *********************
// INICIO EXERCÍCIO 2 👇🏼
// *********************

type Produto struct {
	Tipo  string
	Nome  string
	Preco float64
}

type iProduto interface {
	CalcularCusto() float64
}

func (p Produto) CalcularCusto() float64 {
	if p.Tipo == "Grande" {
		return p.Preco*1.06 + 2500
	} else if p.Tipo == "Medio" {
		return p.Preco * 1.03
	}
	return p.Preco
}

type Loja struct {
	Produtos []iProduto
}

type Ecommerce interface {
	Total() float64
	Adicionar(p Produto)
}

func (l Loja) Total() float64 {
	var total float64
	for _, produto := range l.Produtos {
		total += produto.CalcularCusto()
	}
	return total
}

func (l *Loja) Adicionar(p Produto) {
	l.Produtos = append(l.Produtos, p)
}

func novoProduto(tipo, nome string, preco float64) *Produto {
	return &Produto{
		Tipo:  tipo,
		Nome:  nome,
		Preco: preco,
	}
}

func novaLoja() Ecommerce {
	return &Loja{}
}

func exercicioDois() {
	fmt.Println("\nEXERCÍCIO 2")
	loja := novaLoja()
	loja.Adicionar(*novoProduto("Pequeno", "Mousepad", 21))
	loja.Adicionar(*novoProduto("Grande", "Notebook", 1000))
	loja.Adicionar(*novoProduto("Medio", "Mouse", 100))
	loja.Adicionar(*novoProduto("Pequeno", "Teclado", 50))
	loja.Adicionar(*novoProduto("Grande", "Monitor", 500))
	loja.Adicionar(*novoProduto("Medio", "Cadeira", 200))
	loja.Adicionar(*novoProduto("Pequeno", "Mesa", 100))

	fmt.Println(loja.Total())
}
