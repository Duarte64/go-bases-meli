package main

import (
	"errors"
	"fmt"
)

func main() {
	exercicioUm()
	exercicioDois()
	exercicioTres()
	exercicioQuatro()
}

// *********************
// INICIO EXERCÍCIO 1 👇🏼
// *********************

type salarioError struct{}

func (e *salarioError) Error() string {
	return "erro: O salário digitado não está dentro do valor mínimo"
}

func verificaSalario1(salario int) (string, error) {
	if salario < 15000 {
		return "", &salarioError{}
	}
	return "Necessário pagamento de imposto", nil
}

func exercicioUm() {
	fmt.Println("EXERCÍCIO 1")
	salario := 16000
	msg, err := verificaSalario1(salario)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(msg)
}

// *********************
// INICIO EXERCÍCIO 2 👇🏼
// *********************

func verificaSalario2(salario int) (string, error) {
	if salario < 15000 {
		return "", errors.New("erro: O salário digitado não está dentro do valor mínimo")
	}
	return "Necessário pagamento de imposto", nil
}

func exercicioDois() {
	fmt.Println("\nEXERCÍCIO 2")
	salario := 14000
	msg, err := verificaSalario2(salario)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(msg)
}

// *********************
// INICIO EXERCÍCIO 3 👇🏼
// *********************

func verificaSalario3(salario int) (string, error) {
	if salario < 15000 {
		return "", fmt.Errorf("erro: o mínimo tributável é 15.000 e o salário informado é: %d", salario)
	}
	return "Necessário pagamento de imposto", nil
}

func exercicioTres() {
	fmt.Println("\nEXERCÍCIO 3")
	salario := 14000
	msg, err := verificaSalario3(salario)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(msg)
}

// *********************
// INICIO EXERCÍCIO 4 👇🏼
// *********************

type Funcionario struct {
	valorHora        float64
	horasTrabalhadas float64
}

func calculaSalario(horas, valor float64) (float64, error) {
	if horas < 80 {
		return 0, errors.New("erro: o trabalhador não pode ter trabalhado menos de 80 horas por mês")
	}
	return horas * valor, nil
}

func (f *Funcionario) verificaSalario4() (float64, error) {
	salario, err := calculaSalario(f.horasTrabalhadas, f.valorHora)
	if err != nil {
		return 0, err
	} else if salario < 15000 {
		return salario, fmt.Errorf("erro: o mínimo tributável é 15.000 e o salário informado é: %.2f", salario)
	} else if salario > 20000 {
		return salario * 0.9, nil
	}
	return salario, nil
}

func exercicioQuatro() {
	fmt.Println("\nEXERCÍCIO 4")
	funcionario := Funcionario{
		valorHora:        200,
		horasTrabalhadas: 100,
	}
	salario, err := funcionario.verificaSalario4()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Salário: %.2f\n", salario)
}
