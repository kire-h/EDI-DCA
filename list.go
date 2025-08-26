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
}

// struct node que armazena um valor e um ponteiro para o proximo node
type Node struct {
	val  int
	next *Node
}

// struct linked list que armazena o ponteiro para o node head da lista e o inserted para o numero de elementos
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

// retorna o inserted do objeto list que é uma linkedlist
func (list *LinkedList) Size() int { //Theta(1)
	return list.inserted
}

// retorna o valor de um indice percorrendo os nodes da minha lista
func (list *LinkedList) Get(index int) (int, error) { //O(n), Ômega(1)
	if index >= 0 && index < list.inserted {
		aux := list.head
		for i := 0; i < index; i++ {
			aux = aux.next
		}
		return aux.val, nil
	} else {
		return -1, errors.New(fmt.Sprintf("Index inválido: %d", index))
	}
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
	} else {
		return -1, errors.New(fmt.Sprintf("Index inválido: %d", index))
	}
}

func (list *DoubleLinkedList) Add(val int) {
	newNode := &Node2{
		val:  val,
		next: nil,
		prev: nil,
	}
	if list.inserted == 0 {
		list.head = newNode
		list.tail = newNode
	} else {
		list.tail.next = newNode
		newNode.prev = list.tail
		list.tail = newNode
	}
	list.inserted++
}

func (list *DoubleLinkedList) AddOnIndex(val int, index int) error {
	if index >= 0 && index < list.inserted {
		newNode := &Node2{
			val:  val,
			next: nil,
			prev: nil,
		}

		if index == 0 {
			newNode.next = list.head
			list.head.prev = newNode
			list.head = newNode
		} else {
			if index < list.inserted/2 {
				aux := list.head
				for i := 0; i < index-1; i++ {
					aux = aux.next
				}
				newNode.prev = aux
				newNode.next = aux.next
				aux.next.prev = newNode
				aux.next = newNode
			} else {
				aux := list.tail
				for i := list.inserted - 1; i > index; i-- {
					aux = aux.prev
				}
				newNode.prev = aux.prev
				newNode.next = aux
				aux.prev.next = newNode
				aux.prev = newNode
			}
		}
	} else {
		if index == list.inserted {
			list.Add(val)
			return nil
		} else {
			return errors.New(fmt.Sprintf("Index inválido: %d", index))
		}
	}
	list.inserted++
	return nil
}

func (list *LinkedList) Add(val int) {
	newNode := &Node{
		val:  val,
		next: nil,
	}
	if list.head == nil {
		list.head = newNode
	} else {
		aux := list.head
		for aux.next != nil {
			aux = aux.next
		}
		aux.next = newNode
	}
	list.inserted++
}

func (list *LinkedList) AddOnIndex(val int, index int) error {
	if index >= 0 && index < list.inserted {
		newNode := &Node{
			val:  val,
			next: nil,
		}
		if index == 0 {
			newNode.next = list.head
			list.head = newNode
		} else {
			aux := list.head
			for i := 0; i < index-1; i++ {
				aux = aux.next
			}
			newNode.next = aux.next
			aux.next = newNode
		}
	} else {
		return errors.New(fmt.Sprintf("Index inválido: %d", index))
	}
	list.inserted++
	return nil
}

func (list *LinkedList) Remove(index int) error {
	if index >= 0 && index < list.inserted {
		if index == 0 {
			list.head = list.head.next
		} else {
			aux := list.head
			for i := 0; i < index-1; i++ {
				aux = aux.next
			}
			aux.next = aux.next.next
		}
	} else {
		return errors.New(fmt.Sprintf("Index inválido: %d", index))
	}
	list.inserted--
	return nil
}

func (list *DoubleLinkedList) Remove(index int) error {
    if index < 0 || index >= list.inserted {
        return errors.New(fmt.Sprintf("Index inválido: %d", index))
    }
    if list.inserted == 1 {
        list.head = nil
        list.tail = nil
        list.inserted--
        return nil
    }
    if index == 0 {
        list.head = list.head.next
        list.head.prev = nil
        list.inserted--
        return nil
    }
    if index == list.inserted-1 {
        list.tail = list.tail.prev
        list.tail.next = nil
        list.inserted--
        return nil
    }
    if index < list.inserted/2 {
        aux := list.head
        for i := 0; i < index-1; i++ {
            aux = aux.next
        }
        aux.next = aux.next.next
        aux.next.next.prev = aux
    } else {
        aux := list.tail
        for i := list.inserted-1; i > index; i--{
            aux = aux.prev
        }
        aux.prev.next = aux.next
        aux.next.prev = aux.prev
    }
    list.inserted--
    return nil
}

func (list *DoubleLinkedList) PrintForward() {
	fmt.Print("Head -> ")
	for aux := list.head; aux != nil; aux = aux.next {
		fmt.Printf("%d ", aux.val)
	}
	fmt.Println("<- Tail")
}

func (list *DoubleLinkedList) PrintBackward() {
	fmt.Print("Tail -> ")
	for aux := list.tail; aux != nil; aux = aux.prev {
		fmt.Printf("%d ", aux.val)
	}
	fmt.Println("<- Head")
}



func main() {
	list := &DoubleLinkedList{}

	// Adiciona no final
	list.Add(10)
	list.Add(20)
	list.Add(30)
	fmt.Println("Após Add:")
	list.PrintForward()
	list.PrintBackward()

	// Adiciona em índice
	list.AddOnIndex(15, 1)
	list.AddOnIndex(5, 0)
	list.AddOnIndex(35, 5)
	fmt.Println("\nApós AddOnIndex:")
	list.PrintForward()
	list.PrintBackward()

	// Remove elementos
	list.Remove(0) // remove o primeiro
	list.Remove(2) // remove do meio
	list.Remove(list.inserted - 1) // remove o último
	fmt.Println("\nApós Remove:")
	list.PrintForward()
	list.PrintBackward()
}
