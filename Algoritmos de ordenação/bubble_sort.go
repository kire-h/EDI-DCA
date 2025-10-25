package main

import "fmt"

func BubbleSort(v []int) {
	for varredura := 0; varredura < len(v)-1; varredura++ {
		trocou := false
		for i := 0; i < len(v)-varredura-1; i++ {
			if v[i] > v[i+1] {
				v[i], v[i+1] = v[i+1], v[i]
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
