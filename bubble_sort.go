package main

import "fmt"

func BubbleSort(v []int) []int {
	for varredura := 0; varredura < len(v)-1; varredura++ {
		for i := 0; i < len(v)-varredura-1; i++ {
			if v[i] > v[i+1] {
				v[i], v[i+1] = v[i+1], v[i]
			}
		}
	}
	return v
}

func main() {
	// Test array
	arr := []int{64, 34, 25, 12, 22, 11, 90}

	fmt.Println("Original array:", arr)

	// Call BubbleSort
	sortedArr := BubbleSort(arr)

	fmt.Println("Sorted array:", sortedArr)
}
