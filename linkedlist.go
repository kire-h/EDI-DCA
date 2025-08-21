package main

import (
	"fmt"
	"errors"
)
//funcoes para o tipo list
type List interface {
	Size() int
	Get(index int) (int, error)
	Add(e int) 
	AddOnIndex(e int, index int) error
	Remove(index int) error
}
//struct node que armazena um valor e um ponteiro para o proximo node
type Node struct {
	val int
	next * Node
}
//struct linked list que armazena o ponteiro para o node head da lista e o inserted para o numero de elementos
type LinkedList struct {
	head * Node
	inserted int
}
//retorna o inserted do objeto list que é uma linkedlist
func (list *LinkedList) Size() int{ //Theta(1)
	return list.inserted
}
//retorna o valor de um indice percorrendo os nodes da minha lista
func (list *LinkedList) Get(index int) (int,error){ //O(n), Ômega(1) 
	if index>=0 && index < list.inserted {
		aux := list.head
		for i:=0; i < index; i++ {
			aux = aux.next
		} 		
		return aux.val, nil
	} else{
		return -1, errors.New(fmt.Sprintf("Index inválido: %d", index))
	}
}

func (list *LinkedList) Add(val int){
	newNode := &Node{
		val: val,
		next: nil,

	}
	if list.head == nil{
		list.head = newNode
	} else{
		aux := list.head
		for aux.next != nil{
			aux = aux.next
		}
		aux.next = newNode
	}
	list.inserted++
}

func (list *LinkedList) AddOnIndex(val int, index int) error{
	if index>=0 && index < list.inserted {
		newNode := &Node{
		val: val,
		next: nil,
		}
		if index == 0{
			newNode.next = list.head
			list.head = newNode
		} else{
			aux:= list.head
			for i:=0;i<index-1;i++{
				aux = aux.next
			}
			newNode.next = aux.next
			aux.next = newNode
		}
	} else{
		return errors.New(fmt.Sprintf("Index inválido: %d", index))
	}
	list.inserted++
	return nil
}


func main(){
	l := &ArrayList{}
	l.Init(10)
	for i:=1; i <= 50; i++{
		l.Add(i)
	}
	val, _ := l.Get(0)
	fmt.Println("Valor na posicao 0: ",val)
	
	val, _ = l.Get(49)
	fmt.Println("Valor na posicao 49: ",val)

	l.AddOnIndex(-1,0)
	
	val, _ = l.Get(0)
	fmt.Println("Valor na posicao 0: ",val)
	
	l.Remove(0)
	
	val, _ = l.Get(0)
	fmt.Println("Valor na posicao 0: ",val)

}
