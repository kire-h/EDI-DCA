import (
	"fmt"
	"errors"
)

type queue interface {
	Enqueue(val int)
	Dequeue() int, error
	Size() int
	Front() int,error
}

type ArrayQueue struct {
	v [] int
	front int
	rear int
}

func (Queue *ArrayQueue) Enqueue(val int){
	if Queue.front == -1 && Queue.rear == -1{
		Queue.front++ 
		Queue.rear++
	} else {
		// tam = 5 [1][2][3][4][5]
		// rear = 0 f           r
		Queue.rear = (Queue.rear+1)%len(Queue.v)
		// rear = 1%5 = 1
	}
	Queue.v[Queue.rear] = val
}

func (Queue *ArrayQueue) Dequeue() (int,error){
	if Queue.front == -1 && Queue.rear == -1{
		return -1,errors.New("Fila Vazia")
	} else if Queue.front == Queue.rear {
		aux := Queue.v[Queue.front]
		Queue.front = -1
		Queue.rear = -1
		return aux,nil
	} else {
		aux:= Queue.v[Queue.front]
		Queue.front = (Queue.front+1)%len(Queue.v)
		return aux, nil
	} 
}

func (Queue *ArrayQueue) Size() int{
	return (fila.rear - fila.front + len(Queue.v)) % len(Queue.v) + 1
}

func (Queue *ArrayQueue) Front() int{
	return Queue.v[Queue.front]
}


