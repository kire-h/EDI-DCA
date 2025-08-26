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
	newS:= len(stack.v)
	if newS == 0{
		newS = 1
	}
	newV := make([] int, newS*2)
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
	stack.inserted--
	return stack.v[stack.top],nil
}

func (stack *ArrayStack) Top() (int, error){
	if stack.IsEmpty(){
		return -1, errors.New(fmt.Sprintf("lista vazia"))
	}
	return stack.v[stack.top-1],nil
}

func (stack *ArrayStack) IsEmpty() bool{
	if stack.inserted == 0 {
		return true
	} else {
		return false
	}
}

func (stack *ArrayStack) Size() int{
	return stack.inserted
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

func testStack(s Stack, name string) {
	fmt.Println("===================================")
	fmt.Println("🔹 Testando:", name)

	fmt.Println("A pilha está vazia?", s.IsEmpty())
	fmt.Println("Tamanho da pilha:", s.Size())

	fmt.Println("➡️ Inserindo elementos: 10, 20, 30")
	s.Push(10)
	s.Push(20)
	s.Push(30)

	fmt.Println("A pilha está vazia?", s.IsEmpty())
	fmt.Println("Tamanho da pilha:", s.Size())

	if top, err := s.Top(); err == nil {
		fmt.Println("Topo da pilha:", top)
	} else {
		fmt.Println("Erro ao acessar topo:", err)
	}

	fmt.Println("🔁 Desempilhando elementos:")
	for !s.IsEmpty() {
		if val, err := s.Pop(); err == nil {
			fmt.Println("Desempilhado:", val)
		} else {
			fmt.Println("Erro:", err)
		}
	}

	fmt.Println("A pilha está vazia?", s.IsEmpty())
	fmt.Println("Tamanho final da pilha:", s.Size())

	if _, err := s.Pop(); err != nil {
		fmt.Println("⚠️ Erro esperado ao desempilhar de pilha vazia:", err)
	}
}

func main() {
	// Testando ArrayStack (precisa de slice inicial para não dar panic)
	testStack(&ArrayStack{v: make([]int, 1)}, "ArrayStack")

	// Testando LinkedStack
	testStack(&LinkedStack{}, "LinkedStack")
}


