package main

import (
	"errors"
	"fmt"
)

func (queue *ArrayQueue) init(size int) {
	queue.v = make([]int, size)
	queue.rear = -1
	queue.front = -1
}

type queue interface {
	Enqueue(val int)
	Dequeue() (int, error)
	Size() int
	Front() (int, error)
}

type ArrayQueue struct {
	v     []int
	front int
	rear  int
}

type Node3 struct {
	next *Node3
	val  int
}

type LinkedQueue struct {
	front    *Node3
	rear     *Node3
	inserted int
}

func (Queue *LinkedQueue) Enqueue(val int) {
	newNode := &Node3{
		next: nil,
		val:  val,
	}
	if Queue.front == nil && Queue.rear == nil {
		Queue.front = newNode
		Queue.rear = newNode
	} else {
		Queue.rear.next = newNode
		Queue.rear = newNode
	}
	Queue.inserted++
}

func (Queue *ArrayQueue) Enqueue(val int) {
	if Queue.Size() == len(Queue.v) {
		Queue.doubleV()
	}
	if Queue.front == -1 && Queue.rear == -1 {
		Queue.front++
		Queue.rear++
	} else {
		// tam = 5 [1][2][3][4][5]
		// rear = 0 f           r
		Queue.rear = (Queue.rear + 1) % len(Queue.v)
		// rear = 1%5 = 1
	}
	Queue.v[Queue.rear] = val
}

func (queue *ArrayQueue) doubleV() {
	newV := make([]int, len(queue.v)*2)
	newI := 0
	for i := queue.front; newI < len(queue.v); newI++ {
		newV[newI] = queue.v[i]
		i = (i + 1) % len(queue.v)
	}
	queue.front = 0
	queue.rear = queue.Size() - 1
	queue.v = newV
}

func (Queue *LinkedQueue) Dequeue() (int, error) {
	if Queue.front == nil && Queue.rear == nil {
		return -1, errors.New("Fila Vazia")
	} else if Queue.front == Queue.rear {
		aux := Queue.front.val
		Queue.front = nil
		Queue.rear = nil
		Queue.inserted--
		return aux, nil
	} else {
		aux := Queue.front.val
		Queue.front = Queue.front.next
		Queue.inserted--
		return aux, nil
	}
}

func (Queue *ArrayQueue) Dequeue() (int, error) {
	if Queue.front == -1 && Queue.rear == -1 {
		return -1, errors.New("Fila Vazia")
	} else if Queue.front == Queue.rear {
		aux := Queue.v[Queue.front]
		Queue.front = -1
		Queue.rear = -1
		return aux, nil
	} else {
		aux := Queue.v[Queue.front]
		Queue.front = (Queue.front + 1) % len(Queue.v)
		return aux, nil
	}
}

func (Queue *LinkedQueue) Size() int {
	return Queue.inserted
}

func (Queue *ArrayQueue) Size() int {
	if Queue.front == -1 && Queue.rear == -1 {
		return 0
	}
	return (Queue.rear-Queue.front+len(Queue.v))%len(Queue.v) + 1
}

func (Queue *LinkedQueue) Front() (int, error) {
	if Queue.front != nil {
		return Queue.front.val, nil
	} else {
		return -1, errors.New("Fila Vazia")
	}
}

func (Queue *ArrayQueue) Front() (int, error) {
	if Queue.front == -1 && Queue.rear == -1 {
		return -1, errors.New("Fila Vazia")
	}
	return Queue.v[Queue.front], nil
}

func main() {
	fmt.Println("=== Testando LinkedQueue ===")
	lq := &LinkedQueue{}
	lq.Enqueue(10)
	lq.Enqueue(20)
	lq.Enqueue(30)

	fmt.Println("Front:", must(lq.Front()))
	fmt.Println("Size:", lq.Size())

	val, _ := lq.Dequeue()
	fmt.Println("Dequeued:", val)
	fmt.Println("Front:", must(lq.Front()))
	fmt.Println("Size:", lq.Size())

	fmt.Println("\n=== Testando ArrayQueue ===")
	aq := &ArrayQueue{v: make([]int, 5), front: -1, rear: -1}
	aq.Enqueue(100)
	aq.Enqueue(200)
	aq.Enqueue(300)

	fmt.Println("Front:", must(aq.Front()))
	fmt.Println("Size:", aq.Size())

	val2, _ := aq.Dequeue()
	fmt.Println("Dequeued:", val2)
	fmt.Println("Front:", must(aq.Front()))
	fmt.Println("Size:", aq.Size())
}

// Função auxiliar para ignorar erro de Front
func must(val int, err error) int {
	if err != nil {
		panic(err)
	}
	return val
}
