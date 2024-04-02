package main

import (
	"errors"
	"fmt"
)

func main() {
	exercicioUm()
	exercicioDois()
	exericicoTres()
	exercicioQuatro()
	exercicioCinco()
}

// *********************
// INICIO EXERCÍCIO 1 👇🏼
// *********************

func exercicioUm() {
	fmt.Println("EXERCÍCIO 1")
	printSalarioImposto := "Salário: R$ %.2f\nImposto de Renda: R$ %.2f\n"
	salario1 := 5000.0
	salario2 := 100000.0
	salario3 := 1000000.0
	imposto1 := impostoRenda(salario1)
	imposto2 := impostoRenda(salario2)
	imposto3 := impostoRenda(salario3)
	fmt.Printf(printSalarioImposto, salario1, imposto1)
	fmt.Printf(printSalarioImposto, salario2, imposto2)
	fmt.Printf(printSalarioImposto, salario3, imposto3)
}

func impostoRenda(salario float64) float64 {
	if salario <= 50000 {
		return 0
	} else if salario <= 150000 {
		return salario * 0.17
	}
	return salario * 0.27
}

// *********************
// INICIO EXERCÍCIO 2 👇🏼
// *********************

func exercicioDois() {
	fmt.Println("\nEXERCÍCIO 2")
	media, err := calculaMedia(7.5, 8.0, 9.0, 10.0)
	if err == nil {
		fmt.Printf("Média: %.2f\n", media)
		return
	}
	fmt.Println(err)
}

func calculaMedia(notas ...float64) (float64, error) {
	var total float64 = 0
	for _, nota := range notas {
		if nota < 0 {
			return 0, fmt.Errorf("nota %.2f inválida", nota)
		}
		total += nota
	}

	return total / float64(len(notas)), nil
}

// *********************
// INICIO EXERCÍCIO 3 👇🏼
// *********************

func exericicoTres() {
	fmt.Println("\nEXERCÍCIO 3")
	funcionarios := []map[string]interface{}{
		{
			"categoria": "A",
			"horas":     172.0,
		},
		{
			"categoria": "B",
			"horas":     176.0,
		},
		{
			"categoria": "C",
			"horas":     162.0,
		},
		{
			"categoria": "D",
			"horas":     202.0,
		},
	}
	for _, funcionario := range funcionarios {
		salario, error := calculaSalario(funcionario)
		if error == nil {
			fmt.Printf("Salário: R$ %.2f\n", salario)
		} else {
			fmt.Println(error)
		}
	}
}

func calculaSalario(funcionario map[string]interface{}) (float64, error) {
	categoria := funcionario["categoria"].(string)
	horas := funcionario["horas"].(float64)
	var valorHora float64
	var adicional float64 = 1.0
	switch categoria {
	case "A":
		valorHora = 3000.0
		if horas > 160 {
			adicional = 1.5
		}
	case "B":
		valorHora = 1500.0
		if horas > 160 {
			adicional = 1.2
		}
	case "C":
		valorHora = 1000.0
	default:
		return 0.0, errors.New("categoria inválida")
	}
	return valorHora * horas * adicional, nil
}

// *********************
// INICIO EXERCÍCIO 4 👇🏼
// *********************

const (
	minimum = "minimum"
	maximum = "maximum"
	average = "average"
)

func exercicioQuatro() {
	fmt.Println("\nEXERCÍCIO 4")
	minFunc, _ := operation(minimum)
	maxFunc, _ := operation(maximum)
	avgFunc, _ := operation(average)
	_, err := operation("error")

	minValue := minFunc(9, 2, 6, 4, 5, 17, 1, 3, 14)
	maxValue := maxFunc(9, 2, 6, 4, 5, 17, 1, 3, 14)
	avgValue := avgFunc(9, 2, 6, 4, 5, 17, 1, 3, 14)

	fmt.Printf("Mínimo: %.1f\n", minValue)
	fmt.Printf("Máximo: %.1f\n", maxValue)
	fmt.Printf("Média: %.2f\n", avgValue)
	if err != nil {
		fmt.Println(err)
	}
}

func operation(operation string) (func(values ...float64) float64, error) {
	switch operation {
	case minimum:
		return minimumFunc, nil
	case maximum:
		return maximumFunc, nil
	case average:
		return averageFunc, nil
	default:
		return nil, errors.New("operação não encontrada")
	}

}

func minimumFunc(values ...float64) float64 {
	min := values[0]
	for _, value := range values {
		if value < min {
			min = value
		}
	}
	return min
}

func maximumFunc(values ...float64) float64 {
	max := values[0]
	for _, value := range values {
		if value > max {
			max = value
		}
	}
	return max
}

func averageFunc(values ...float64) float64 {
	total := 0.0
	for _, value := range values {
		total += value
	}
	return total / float64(len(values))
}

// *********************
// INICIO EXERCÍCIO 5 👇🏼
// *********************

const (
	dog       = "cachorro"
	cat       = "gato"
	hamster   = "hamster"
	tarantula = "tarântula"
)

func exercicioCinco() {
	fmt.Println("\nEXERCÍCIO 5")
	animals := []string{dog, cat, hamster, tarantula, "fish"}
	for i, animal := range animals {
		food, err := animalFood(animal)
		if err == nil {
			fmt.Printf("Para alimentar %d %s(s) é preciso %d gramas de alimento.\n", (i+1)*(i+1), animal, food((i+1)*(i+1)))
		} else {
			fmt.Println(err)
		}
	}
}

func animalFood(animal string) (func(int) int, error) {
	switch animal {
	case dog:
		return dogFood, nil
	case cat:
		return catFood, nil
	case hamster:
		return hamsterFood, nil
	case tarantula:
		return tarantulaFood, nil
	default:
		return nil, errors.New("animal não encontrado")
	}
}

func dogFood(quant int) int {
	return quant * 10000
}

func catFood(quant int) int {
	return quant * 5000
}

func hamsterFood(quant int) int {
	return quant * 250
}

func tarantulaFood(quant int) int {
	return quant * 150
}
