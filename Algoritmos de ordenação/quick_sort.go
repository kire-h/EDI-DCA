package main

import (
	"fmt"
	"math/rand"
	"time"
)

func quick_sort(v []int, ini int, fim int) { // vou dividir o vetor até ele n ser mais divisivel de acordo com o pivot que vai ficar no meio do vetor
	if ini >= fim {
		return
	}
	indexpivot := partition(v, ini, fim)
	quick_sort(v, ini, indexpivot-1)
	quick_sort(v, indexpivot+1, fim)
}

func partition(v []int, ini int, fim int) int {
	// Gerar um valor aleatório entre ini e fim (inclusive)
	randIndex := rand.Intn(fim-ini+1) + ini

	// Trocar os elementos v[randIndex] e v[fim]
	v[randIndex], v[fim] = v[fim], v[randIndex]

	pivot := v[fim]
	pIndex := ini
	for i := ini; i < fim; i++ {
		if v[i] <= pivot {
			v[pIndex], v[i] = v[i], v[pIndex]
			pIndex++
		}
	}
	v[pIndex], v[fim] = v[fim], v[pIndex]
	return pIndex
}

func main() {
	// Seed para números aleatórios (necessário para pivô aleatório)
	rand.Seed(time.Now().UnixNano())

	// Vetor de teste
	vetor := []int{9, 3, 7, 2, 8, 5, 1, 6, 4}

	fmt.Println("Vetor original:", vetor)

	// Chamada do Quick Sort
	quick_sort(vetor, 0, len(vetor)-1)

	fmt.Println("Vetor ordenado:", vetor)
}
