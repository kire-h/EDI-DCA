package main

import "fmt"

func BubbleSort(v []int) {
	for varredura := 0; varredura < len(v)-1; varredura++ { //vai varrer a quantidade de vezes necessárias para realizar todas as trocas
		trocou := false                           // bool responsavel por saber quando o vetor já está ordenado para melhor caso ser omega(n)
		for i := 0; i < len(v)-varredura-1; i++ { //varre até o final do vetor -1 (já que eu preciso analisar o proximo elemento (i+1))
			if v[i] > v[i+1] { // - varredura porque diminui a quantidade que eu tenho que percorrer na parte de tras do vetor
				v[i], v[i+1] = v[i+1], v[i] // a cada iteração o laço interno repete (n-1) + (n-2) + (n-3) + ... + 1 = n(n-1)/2 O(n^2)
				trocou = true
			}
		}
		// Se nenhuma troca foi feita, o vetor já está ordenado
		if !trocou {
			return
		}
	}
}

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}

	fmt.Println("Original array:", arr)

	BubbleSort(arr) // ordena "in place"

	fmt.Println("Sorted array:", arr)
}
