/*Exercício 4 - Ordenamento

Uma empresa de sistemas precisa analisar que algoritmos de ordenamento utilizar para seus serviços.

Para eles é necessário instanciar 3 arrays com valores aleatórios não ordenados

uma matriz de inteiros com 100 valores
uma matriz de inteiros com 1000 valores
uma matriz de inteiros com 10.000 valores

Para instanciar as variáveis, utilize o rand*/

package ex04

import (
	"fmt"
	"math/rand"
	"time"
)

func ExercicioQuatro() {
	fmt.Println("\nEXERCÍCIO 4")
	c1 := make(chan time.Duration)
	c2 := make(chan time.Duration)
	c3 := make(chan time.Duration)
	variavel1 := rand.Perm(10000)
	variavel2 := rand.Perm(10000)
	variavel3 := rand.Perm(10000)
	go insertionSort(&variavel1, c1)
	go bubbleSort(&variavel2, c2)
	go selectionSort(&variavel3, c3)
	fmt.Printf("Execução do insertion sort levou %s\n", <-c1)
	fmt.Printf("Execução do bubble sort levou %s\n", <-c2)
	fmt.Printf("Execução do selection sort levou %s\n", <-c3)
}

func insertionSort(arr *[]int, c chan time.Duration) {
	start := time.Now()
	for i := 1; i < len(*arr); i++ {
		key := (*arr)[i]
		j := i - 1

		for j >= 0 && (*arr)[j] > key {
			(*arr)[j+1] = (*arr)[j]
			j = j - 1
		}
		(*arr)[j+1] = key
	}
	c <- time.Since(start)
}

func bubbleSort(arr *[]int, c chan time.Duration) {
	start := time.Now()
	for i := 0; i < len(*arr); i++ {
		for j := 0; j < len(*arr)-i-1; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
	}
	c <- time.Since(start)
}

func selectionSort(arr *[]int, c chan time.Duration) {
	start := time.Now()
	for i := 0; i < len(*arr); i++ {
		minIndex := i
		for j := i + 1; j < len(*arr); j++ {
			if (*arr)[j] < (*arr)[minIndex] {
				minIndex = j
			}
		}
		(*arr)[i], (*arr)[minIndex] = (*arr)[minIndex], (*arr)[i]
	}
	c <- time.Since(start)
}
