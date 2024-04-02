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
	var palavra string = "BOOTCAMP"
	var tamanhoPalavra int = len(palavra)
	fmt.Printf("A palavra %s tem %d letras.\n", palavra, tamanhoPalavra)
	for i := 0; i < tamanhoPalavra; i++ {
		fmt.Printf("%d° letra: %c\n", i+1, palavra[i])
	}
}

func exercicioDois() {
	fmt.Printf("\nEXERCÍCIO 2\n")
	clientes := []map[string]interface{}{
		{
			"nome":          "João",
			"idade":         23,
			"ehEmpregado":   true,
			"anosAtividade": 4,
			"salario":       100000.1,
		},
		{
			"nome":          "Maria",
			"idade":         22,
			"ehEmpregado":   false,
			"anosAtividade": 0,
			"salario":       5000.0,
		},
		{
			"nome":          "José",
			"idade":         25,
			"ehEmpregado":   true,
			"anosAtividade": 2,
			"salario":       90000.0,
		},
	}
	for _, cliente := range clientes {
		nome := cliente["nome"].(string)
		idade := cliente["idade"].(int)
		ehEmpregado := cliente["ehEmpregado"].(bool)
		anosAtividade := cliente["anosAtividade"].(int)
		salario := cliente["salario"].(float64)
		if idade > 22 && ehEmpregado && anosAtividade > 1 {
			if salario > 100000 {
				fmt.Printf("%s, seu empréstimo foi aprovado SEM juros.\n", nome)
			} else {
				fmt.Printf("%s, seu empréstimo aprovado COM juros.\n", nome)
			}
		} else {
			fmt.Printf("%s, seu empréstimo foi negado.\n", nome)
		}
	}
}

func exercicioTres() {
	fmt.Printf("\nEXERCÍCIO 3\n")
	const monthMessage = "O mês %d é %s\n"
	mesMap := map[int]string{
		1:  "Janeiro",
		2:  "Fevereiro",
		3:  "Março",
		4:  "Abril",
		5:  "Maio",
		6:  "Junho",
		7:  "Julho",
		8:  "Agosto",
		9:  "Setembro",
		10: "Outubro",
		11: "Novembro",
		12: "Dezembro",
	}
	fmt.Printf(monthMessage, 1, mesMap[1])
	fmt.Printf(monthMessage, 4, mesMap[4])
	fmt.Printf(monthMessage, 10, mesMap[10])
}

func exercicioQuatro() {
	fmt.Printf("\nEXERCÍCIO 4\n")
	var maior21 int = 0
	var employees = map[string]int{
		"Benjamin": 20,
		"Manuel":   26,
		"Brenda":   19,
		"Dario":    44,
		"Pedro":    30,
	}
	fmt.Printf("Benjamin tem %d anos\n", employees["Benjamin"])
	for _, age := range employees {
		if age > 21 {
			maior21++
		}
	}
	fmt.Printf("Existem %d pessoas com mais de 21 anos\n", maior21)
	employees["Federico"] = 25
	delete(employees, "Pedro")
	for name, age := range employees {
		fmt.Printf("%s tem %d anos\n", name, age)
	}
}
