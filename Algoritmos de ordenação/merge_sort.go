package main

import "fmt"

func merge(v []int, e []int, d []int) { //juntar os elementos da esquerda e da direita em um vetor v em ordem
	eindex := 0
	dindex := 0
	i := 0
	for eindex < len(e) && dindex < len(d) {
		if e[eindex] < d[dindex] {
			v[i] = e[eindex]
			eindex++
		} else {
			v[i] = d[dindex]
			dindex++
		}
		i++
	}
	for eindex < len(e) {
		v[i] = e[eindex]
		eindex++
		i++
	}
	for dindex < len(d) {
		v[i] = d[dindex]
		dindex++
		i++
	}
}

func merge_sort(v []int) { //dividir o vetor em dois vetores esquerda e direita e quando possivel aplicar o merge
	if len(v) > 1 {
		meio := len(v) / 2
		e := make([]int, meio)
		d := make([]int, len(v)-meio)
		i := 0
		for eindex := 0; eindex < len(e); eindex++ {
			e[eindex] = v[i]
			i++
		}
		for dindex := 0; dindex < len(d); dindex++ {
			d[dindex] = v[i]
			i++
		}
		merge_sort(e)
		merge_sort(d)
		merge(v, e, d)
	}
}

func main() {
	v := []int{9, 4, 3, 6, 3, 2, 5, 7, 1, 8}

	merge_sort(v)

	fmt.Printf("%d", v)
}
