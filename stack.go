package main

import (
	"errors"
	"fmt"
)

type Stack interface{
	Push(value int)
	Pop()(int , error)
	Top()(int , error)
	IsEmpty() bool
	Size() int
}

type ArrayStack struct{
	inserted int
	top int
	v []int
}

type Node3 struct{
	value int
	next * Node3
}

type LinkedStack struct{
	top * Node3
	inserted int
}
func (stack *ArrayStack) Push(value int){
	if stack.inserted == len(stack.v){
		stack.doubleV()
	}
	stack.v[stack.top] = value
	stack.top++
	stack.inserted++
}

func (stack *ArrayStack) doubleV(){
	newV := make([] int, len(stack.v)*2)
	for i:=0;i<stack.inserted;i++{
		newV[i] = stack.v[i]
	}
	stack.v = newV
}

func (stack *ArrayStack) Pop() (int, error){
	if stack.IsEmpty(){
		return -1, errors.New(fmt.Sprintf("lista vazia"))
	}
	stack.top = stack.top-1
	return stack.v[stack.top],nil
}

func (stack *LinkedStack) Push(value int){
	newNode := &Node3{
		value: value,
		next: nil,
	}
	newNode.next = stack.top
	stack.top = newNode
	stack.inserted++
}

func (stack *LinkedStack) Pop() (int,error){
	if stack.IsEmpty(){
		return -1, errors.New(fmt.Sprintf("lista vazia"))
	}
	aux:= stack.top.value
	stack.top = stack.top.next
	stack.inserted--
	return aux,nil 
}

func (stack *LinkedStack) Top() (int,error){
	if stack.IsEmpty(){
		return -1, errors.New(fmt.Sprintf("lista vazia"))
	}
	return stack.top.value,nil
}

func (stack *LinkedStack) IsEmpty() bool{
	if stack.inserted == 0{
		return true
	} else{
		return false
	}
}

func (stack *LinkedStack) Size() int{
	return stack.inserted
}

func main() {
	var s Stack = &LinkedStack{}

	fmt.Println("🔹 Testando pilha com lista ligada...")

	fmt.Println("A pilha está vazia?", s.IsEmpty()) // true
	fmt.Println("Tamanho da pilha:", s.Size())      // 0

	fmt.Println("➡️ Inserindo elementos: 10, 20, 30")
	s.Push(10)
	s.Push(20)
	s.Push(30)

	fmt.Println("A pilha está vazia?", s.IsEmpty()) // false
	fmt.Println("Tamanho da pilha:", s.Size())      // 3

	top, err := s.Top()
	if err != nil {
		fmt.Println("Erro ao acessar o topo:", err)
	} else {
		fmt.Println("Topo da pilha:", top) // 30
	}

	fmt.Println("🔁 Desempilhando elementos:")
	for !s.IsEmpty() {
		val, err := s.Pop()
		if err != nil {
			fmt.Println("Erro ao desempilhar:", err)
		} else {
			fmt.Println("Desempilhado:", val)
		}
	}

	fmt.Println("A pilha está vazia?", s.IsEmpty()) // true
	fmt.Println("Tamanho final da pilha:", s.Size()) // 0

	_, err = s.Pop()
	if err != nil {
		fmt.Println("⚠️ Erro esperado ao tentar desempilhar de uma pilha vazia:", err)
	}
}


