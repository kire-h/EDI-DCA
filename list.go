package main

import (
	"errors"
	"fmt"
)

// funcoes para o tipo list
type List interface {
	Size() int
	Get(index int) (int, error)
	Add(e int)
	AddOnIndex(e int, index int) error
	Remove(index int) error
	Invert()
	Set(value int, index int) error
}

type ArrayList struct {
	inserted int
	v        []int
}

type Node struct {
	val  int
	next *Node
}

type LinkedList struct {
	head     *Node
	inserted int
}

type Node2 struct {
	val  int
	next *Node2
	prev *Node2
}

type DoubleLinkedList struct {
	head     *Node2
	tail     *Node2
	inserted int
}

// FUNCOES SIZE
func (list *ArrayList) Size() int {
	return list.inserted
}

func (list *LinkedList) Size() int {
	return list.inserted
}

func (list *DoubleLinkedList) Size() int {
	return list.inserted
}

func (list *ArrayList) Init(size int) {
	list.v = make([]int, size)
}

// FUNCOES GET
func (list *ArrayList) Get(index int) (int, error) {
	if index >= 0 && index < list.inserted {
		return list.v[index], nil
	}
	return -1, errors.New("Indice invalido")
}

func (list *LinkedList) Get(index int) (int, error) {
	if index >= 0 && index < list.inserted {
		aux := list.head
		for i := 0; i < index; i++ {
			aux = aux.next
		}
		return aux.val, nil
	}
	return -1, errors.New("Indice invalido")
}

func (list *DoubleLinkedList) Get(index int) (int, error) {
	if index >= 0 && index < list.inserted {
		if index < list.inserted/2 {
			aux := list.head
			for i := 0; i < index; i++ {
				aux = aux.next
			}
			return aux.val, nil
		} else {
			aux := list.tail
			for i := list.inserted - 1; i > index; i-- {
				aux = aux.prev
			}
			return aux.val, nil
		}
	}
	return -1, errors.New("Indice invalido")
}

func (list *ArrayList) doubleV() {
	newV := make([]int, list.Size()*2)
	for i := 0; i < list.inserted; i++ {
		newV[i] = list.v[i]
	}
	list.v = newV
}

// FUNCOES ADD
func (list *ArrayList) Add(e int) {
	if list.inserted == len(list.v) {
		list.doubleV()
	}
	list.v[list.inserted] = e
	list.inserted++
}

func (list *LinkedList) Add(e int) {
	newNode := &Node{
		val:  e,
		next: nil,
	}
	if list.head == nil {
		list.head = newNode
		list.inserted++
		return
	}
	aux := list.head
	for i := 0; i < list.inserted-1; i++ {
		aux = aux.next
	}
	aux.next = newNode
	list.inserted++
}

func (list *DoubleLinkedList) Add(e int) {
	newNode := &Node2{
		val:  e,
		next: nil,
		prev: nil,
	}
	if list.head == nil {
		list.head = newNode
		list.tail = newNode
		list.inserted++
		return
	} else {
		list.tail.next = newNode
		newNode.prev = list.tail
		list.tail = newNode
		list.inserted++
		return
	}
}

// FUNCOES ADDONINDEX
func (list *ArrayList) AddOnIndex(val int, index int) error {
	if index >= 0 && index <= list.inserted {
		if list.inserted == len(list.v) {
			list.doubleV()
		}
		for i := list.inserted; i > index; i-- {
			list.v[i] = list.v[i-1]
			//[1][2][3][4][]
		}
		list.v[index] = val
		list.inserted++
		return nil
	} else {
		return errors.New("Indice invalido")
	}
}

func (list *LinkedList) AddOnIndex(val int, index int) error {
	newNode := &Node{
		val:  val,
		next: nil,
	}

	if index >= 0 && index <= list.inserted {
		if index == 0 {
			newNode.next = list.head
			list.head = newNode
			list.inserted++
			return nil
		} else {
			aux := list.head
			for i := 0; i < index-1; i++ {
				aux = aux.next
				//[0][1][2][3][4]
			}
			newNode.next = aux.next
			aux.next = newNode
			list.inserted++
			return nil
		}
	} else {
		return errors.New("Indice invalido")
	}
}

func (list *DoubleLinkedList) AddOnIndex(val int, index int) error {
	newNode := &Node2{
		val:  val,
		next: nil,
		prev: nil,
	}

	if index >= 0 && index <= list.inserted {
		if list.head == nil && list.tail == nil {
			list.head = newNode
			list.tail = newNode
			list.inserted++
			return nil
		} else if index == 0 {
			newNode.next = list.head
			list.head.prev = newNode
			list.head = newNode
			list.inserted++
			return nil
		} else if index == list.inserted {
			newNode.prev = list.tail
			list.tail.next = newNode
			list.tail = newNode
			list.inserted++
			return nil
		} else if index < list.inserted/2 {
			aux := list.head
			for i := 0; i < index-1; i++ {
				aux = aux.next
			}
			newNode.next = aux.next
			newNode.prev = aux
			aux.next.prev = newNode
			aux.next = newNode
			list.inserted++
			return nil
		} else {
			aux := list.tail
			for i := list.inserted - 1; i > index; i-- {
				aux = aux.prev
			}
			newNode.prev = aux.prev
			newNode.next = aux
			aux.prev.next = newNode
			aux.prev = newNode
			list.inserted++
			return nil
		}
	} else {
		return errors.New("Indice invalido")
	}
}

// FUNCOES REMOVE
func (list *ArrayList) Remove(index int) error {
	if index >= 0 && index < list.inserted {
		if list.inserted == 0 {
			return errors.New("Lista Vazia")
		} else {
			for i := index; i < list.inserted-1; i++ {
				list.v[i] = list.v[i+1]
				//[1][3][4][4]
			}
			list.inserted--
			return nil
		}
	} else {
		return errors.New("Indice invalido")
	}
}

func (list *LinkedList) Remove(index int) error {
	if index >= 0 && index < list.inserted {
		if list.inserted == 0 {
			return errors.New("Lista Vazia")
		} else {
			if list.inserted == 1 {
				list.head = nil
				list.inserted--
				return nil
			} else if index == 0 {
				list.head = list.head.next
				list.inserted--
				return nil
			} else {
				aux := list.head
				for i := 0; i < index-1; i++ {
					aux = aux.next
				}
				aux.next = aux.next.next
				list.inserted--
				return nil
			}
		}
	} else {
		return errors.New("Indice invalido")
	}
}

func (list *DoubleLinkedList) Remove(index int) error {
	if index >= 0 && index < list.inserted {
		if list.inserted == 0 {
			return errors.New("Lista Vazia")
		} else {
			if list.inserted == 1 {
				list.head = nil
				list.tail = nil
				list.inserted--
				return nil
			} else if index == 0 {
				list.head = list.head.next
				list.head.prev = nil
				list.inserted--
				return nil
			} else if index == list.inserted-1 {
				list.tail = list.tail.prev
				list.tail.next = nil
				list.inserted--
				return nil
			} else {
				if index < list.inserted/2 {
					aux := list.head
					for i := 0; i < index-1; i++ {
						aux = aux.next
					}
					aux.next = aux.next.next
					aux.next.prev = aux
					list.inserted--
					return nil
				} else {
					aux := list.tail
					for i := list.inserted - 1; i > index+1; i-- {
						aux = aux.prev
					}
					aux.prev = aux.prev.prev
					aux.prev.next = aux
					list.inserted--
					return nil
				}
			}
		}
	} else {
		return errors.New("Indice invalido")
	}
}

//FUNCOES INVERT

func (list *ArrayList) Invert() {
	for i := 0; i < list.inserted/2; i++ {
		aux := list.v[i]
		list.v[i] = list.v[list.inserted-1-i]
		list.v[list.inserted-1-i] = aux
	}
	//[1][2][3][4]
	//[4][3][2][1]
}

func (list *LinkedList) Invert() {
	var prev *Node
	var curr = list.head
	var next *Node

	for i := 0; i < list.inserted; i++ {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	list.head = prev
}

func (list *DoubleLinkedList) Invert() {
	var curr = list.head
	var temp *Node2

	for i := 0; i < list.inserted; i++ {
		temp := curr.prev
		curr.prev = curr.next
		curr.next = temp
		curr = curr.prev
	}
	temp = list.head
	list.head = list.tail
	list.tail = temp
}

// FUNCOES SET

func (list *ArrayList) Set(value int, index int) error {
	if index >= 0 && index < list.inserted {
		list.v[index] = value
		return nil
	} else {
		return errors.New("Indice invalido")
	}
}

func (list *LinkedList) Set(value int, index int) error {
	if index >= 0 && index < list.inserted {
		aux := list.head
		for i := 0; i < index; i++ {
			aux = aux.next
		}
		aux.val = value
		return nil
	} else {
		return errors.New("Indice invalido")
	}
}

func (list *DoubleLinkedList) Set(value int, index int) error {
	if index >= 0 && index < list.inserted {
		if index < list.inserted/2 {
			aux := list.head
			for i := 0; i < index; i++ {
				aux = aux.next
			}
			aux.val = value
			return nil
		} else {
			aux := list.tail
			for i := list.inserted - 1; i > index; i-- {
				aux = aux.prev
			}
			aux.val = value
			return nil
		}

	} else {
		return errors.New("Indice invalido")
	}
}

// FUNCOES PRINT

func (list *ArrayList) Print() {
	for i := 0; i < list.inserted; i++ {
		fmt.Printf("%d ", list.v[i])
	}
}

func (list *LinkedList) Print() {
	aux := list.head
	for i := 0; i < list.inserted; i++ {
		fmt.Printf("%d ", aux.val)
		aux = aux.next
	}
}

func (list *DoubleLinkedList) Print() {
	aux := list.head
	for i := 0; i < list.inserted; i++ {
		fmt.Printf("%d ", aux.val)
		aux = aux.next
	}
}

func main() {
	list := &DoubleLinkedList{}

	// Adiciona no final
	list.Add(10)
	list.Add(20)
	list.Add(30)
	list.Print()

	// Adiciona em índice
	list.AddOnIndex(15, 1)
	list.AddOnIndex(5, 0)
	list.AddOnIndex(35, 5)
	fmt.Println("\nApós AddOnIndex:")
	list.Print()

	// Remove elementos
	list.Remove(0)                 // remove o primeiro
	list.Remove(2)                 // remove do meio
	list.Remove(list.inserted - 1) // remove o último
	fmt.Println("\nApós Remove:")
	list.Print()
	fmt.Println("Após Invert:")
	list.Invert()
	list.Print()
}
