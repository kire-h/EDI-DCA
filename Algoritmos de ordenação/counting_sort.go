package main

import (
	"fmt"
)

func counting_sort(v []int) []int {
	//passo 1 - achar o maior e o menor valor do vetor
	menor, maior := v[0], v[0]
	for i := 1; i < len(v); i++ {
		if v[i] < menor {
			menor = v[i]
		}
		if v[i] > maior {
			maior = v[i]
		}
	}
	//passo 2 - criar um vetor de ocorrencias de cada digito no intervalo maior-menor+1
	c := make([]int, maior-menor+1)
	//passo 3 - incrementar o vetor de contagem de ocorrencias de cada valor
	for i := 0; i < len(v); i++ {
		c[v[i]-menor]++
	}
	//passo 4 - fazer a soma comulativa no vetor de contagem
	for i := 0; i < len(c)-1; i++ {
		c[i+1] += c[i]
	}
	//passo 5 - criar outro vetor que será usando para posicionar os elementos de maneira ordenada
	ord := make([]int, len(v))
	//passo 6 - mapear os valores de forma ordenada usando c
	for i := 0; i < len(v); i++ {
		ord[c[v[i]-menor]-1] = v[i]
		c[v[i]-menor]--
	}
	return ord
}

func main() {
	// Exemplo de vetor para ordenar
	v := []int{5, 2, 9, 1, 5, 6}

	// Exibindo o vetor original
	fmt.Println("Vetor original:", v)

	// Chamando o CountingSort
	ordenado := counting_sort(v)

	// Exibindo o vetor ordenado
	fmt.Println("Vetor ordenado:", ordenado)
}
