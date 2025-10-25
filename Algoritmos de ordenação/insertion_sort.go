package main

import (
	"fmt"
)

func insertion_sort(v []int) {
	for i := 1; i < len(v); i++ { //(n-1) vezes vou começar do 1 para verificar o numero que vem antes dele e vou fazendo isso em todo o vetor
		for j := i; j > 0; j-- { //verificando atras do numero se o numero anterior a ele é maior que o atual se for troca para inserir no lugar certo
			if v[j] < v[j-1] { // a repetição do for interno vai ocorrer 1 + 2 + 3 + ... + (n-1) = n(n-1)/2
				v[j], v[j-1] = v[j-1], v[j]
			} else {
				break
			}
		}
	}
}

func main() {
	vetor := []int{2, 1, 0, 4, 3, 10, 1, 4, 2}
	insertion_sort(vetor)
	fmt.Printf("%d", vetor)
}
