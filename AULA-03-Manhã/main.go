package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
	exercicioUm()
	exercicioDois()
}

// *********************
// INICIO EXERCÍCIO 1 👇🏼
// *********************

type Produto struct {
	ID         string
	Preco      float64
	Quantidade int
}

func (p Produto) toString() string {
	return fmt.Sprintf("\n%s,%.2f,%d", p.ID, p.Preco, p.Quantidade)
}

func exercicioUm() {
	fmt.Println("EXERCÍCIO 1")
	produtos := []Produto{
		{
			ID:         "1",
			Preco:      10.00,
			Quantidade: 5,
		},
		{
			ID:         "2",
			Preco:      20.00,
			Quantidade: 10,
		},
		{
			ID:         "3",
			Preco:      30.00,
			Quantidade: 15,
		},
		{
			ID:         "4",
			Preco:      40.00,
			Quantidade: 20,
		},
		{
			ID:         "5",
			Preco:      50.00,
			Quantidade: 25,
		},
	}
	criaCSV(produtos)
}

func criaCSV(produtos []Produto) {
	var text string = "ID,Preco,Quantidade"
	for _, produto := range produtos {
		text += produto.toString()
	}
	byteText := []byte(text)
	err := os.WriteFile("produtos.csv", byteText, 0644)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Arquivo criado com sucesso!")
}

// *********************
// INICIO EXERCÍCIO 2 👇🏼
// *********************

func exercicioDois() {
	fmt.Println("\nEXERCÍCIO 2")
	text := csvToText("produtos.csv")

	total := 0.0
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		total += calculateLineTotal(line)
		fields := strings.Split(line, ",")
		printFields(fields)
	}
	fmt.Printf("%25.2f\n", total)
}

func csvToText(fileName string) string {
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}
	return string(data)
}

func printFields(fields []string) {
	for i, field := range fields {
		if i == 0 {
			fmt.Printf("%s\t", field)
		} else {
			fmt.Printf("%15s\t", field)
		}
	}
	fmt.Println()
}

func calculateLineTotal(line string) float64 {
	fields := strings.Split(line, ",")
	if len(fields) > 1 {
		preco, errPreco := strconv.ParseFloat(fields[1], 64)
		quantidade, errQuant := strconv.ParseFloat(fields[2], 64)
		if errPreco == nil && errQuant == nil {
			return preco * quantidade
		}
		return 0
	}
	return 0
}
