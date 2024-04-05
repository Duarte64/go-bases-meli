package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	exercicioUm()
	exercicioDois()
}

// *********************
// INICIO EXERCÍCIO 1 👇🏼
// *********************

func readFile(name string) {
	file, err := os.ReadFile(name)
	if err != nil {
		panic("o arquivo indicado não foi encontrado ou está danificado")
	}
	fmt.Println(string(file))
}

func exercicioUm() {
	readFile("customers.txt")
}

// *********************
// INICIO EXERCÍCIO 2 👇🏼
// *********************

type Customer struct {
	Arquivo   int
	Nome      string
	Sobrenome string
	RG        string
	Telefone  string
	Endereco  string
}

func lerDocumento() ([]Customer, error) {
	file, err := os.Open("customers.txt")
	if err != nil {
		return nil, fmt.Errorf("erro: o arquivo indicado não foi encontrado ou está danificado")
	}
	defer file.Close()

	var customers []Customer
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ",")
		arquivo, err := strconv.Atoi(fields[0])

		if err != nil {
			return nil, err
		}

		customer := Customer{
			Arquivo:  arquivo,
			Nome:     fields[1],
			RG:       fields[2],
			Telefone: fields[3],
			Endereco: fields[4],
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

func inserirCliente(customers []Customer) error {
	file, err := os.Create("customers.txt")
	if err != nil {
		return fmt.Errorf("erro ao criar o arquivo")
	}
	defer file.Close()

	for _, customer := range customers {
		_, err := fmt.Fprintf(file, "%d,%s,%s,%s,%s\n", customer.Arquivo, customer.Nome, customer.RG, customer.Telefone, customer.Endereco)
		if err != nil {
			return fmt.Errorf("erro ao escrever no arquivo")
		}
	}

	return nil
}

func exercicioDois() {
	customers, err := lerDocumento()
	if err != nil {
		panic(err)
	}

	cliente := Customer{
		Arquivo:   len(customers) + 1,
		Nome:      "Gabriel",
		Sobrenome: "Duarte",
		RG:        "50426756-0",
		Telefone:  "999999999",
		Endereco:  "Teste",
	}

	customers = append(customers, cliente)

	err = inserirCliente(customers)
	if err != nil {
		panic(err)
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()
}
