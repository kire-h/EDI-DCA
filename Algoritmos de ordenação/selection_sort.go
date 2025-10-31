package main

import (
	"fmt"
)

func selection_sort(v []int) {
	for i := 0; i < len(v)-1; i++ { // numero de iterações necessarias para ordenar
		menor := i                        // quando eu faço a ultima iteração o ultimo elemento ja está em seu lugar
		for j := i + 1; j < len(v); j++ { //percorrendo todos os elementos do vetor a direita do i para encontrar o menor
			if v[j] < v[menor] {
				menor = j // eu encontro alguem menor atualizo o indice
			}
		}
		v[i], v[menor] = v[menor], v[i] // realizo a troca do indice atual com o indice do menor
	}
}

// teste:
// 8, 2, 4, 3, 7
// 2, 8, 4, 3, 7
// 2, 3, 4, 8, 7
// 2, 3, 4, 7, 8

func main() {
	vetor := []int{2, 1, 0, 4, 3, 10, 1, 4, 2}
	selection_sort(vetor)
	fmt.Printf("%d", vetor)
}
