package main

import (
	"errors"
	"fmt"
)

type dequeue interface {
	push_front(val int)
	push_back(val int)
	pull_front() (int, error)
	pull_back() (int, error)
	size() int
	Front() (int, error)
	Back() (int, error)
}

type ArrayDequeue struct {
	rear  int
	front int
	v     []int
}

func (deck *ArrayDequeue) init(size int) {
	deck.v = make([]int, size)
	deck.rear = -1
	deck.front = -1
}

func (deck *ArrayDequeue) push_front(val int) {
	if deck.size() == len(deck.v) {
		deck.doubleV()
	}
	if deck.front == -1 && deck.rear == -1 {
		deck.front = 0
		deck.rear = 0
	} else {
		if deck.front-1 < 0 {
			deck.front = len(deck.v) - 1
		} else {
			deck.front--
		}
	}
	deck.v[deck.front] = val
}

func (deck *ArrayDequeue) push_back(val int) {
	if deck.size() == len(deck.v) {
		deck.doubleV()
	}
	if deck.front == -1 && deck.rear == -1 {
		deck.front = 0
		deck.rear = 0
	} else {
		deck.rear = (deck.rear + 1) % len(deck.v)
	}
	deck.v[deck.rear] = val
}

func (deck *ArrayDequeue) doubleV() {
	newV := make([]int, len(deck.v)*2)
	newI := 0
	for i := deck.front; newI < len(deck.v); newI++ {
		newV[newI] = deck.v[i]
		i = (i + 1) % len(deck.v)
	}
	deck.front = 0
	deck.rear = deck.size() - 1
	deck.v = newV
}

func (deck *ArrayDequeue) pull_front() (int, error) {
	if deck.front == -1 && deck.rear == -1 {
		return -1, errors.New("Fila Vazia")
	} else if deck.front == deck.rear {
		aux := deck.v[deck.front]
		deck.front = -1
		deck.rear = -1
		return aux, nil
	} else {
		aux := deck.v[deck.front]
		deck.front = (deck.front + 1) % len(deck.v)
		return aux, nil
	}
}

func (deck *ArrayDequeue) pull_back() (int, error) {
	if deck.front == -1 && deck.rear == -1 {
		return -1, errors.New("Fila Vazia")
	} else if deck.front == deck.rear {
		aux := deck.v[deck.rear]
		deck.front = -1
		deck.rear = -1
		return aux, nil
	} else {
		aux := deck.v[deck.rear]
		if deck.rear-1 < 0 {
			deck.rear = len(deck.v) - 1
		} else {
			deck.rear--
		}
		return aux, nil
	}
}

func (deck *ArrayDequeue) size() int {
	if deck.front == -1 && deck.rear == -1 {
		return 0
	} else if deck.rear >= deck.front {
		return deck.rear - deck.front + 1
	} else {
		return len(deck.v) - deck.front + deck.rear + 1
	}
}

func (deck *ArrayDequeue) Front() (int, error) {
	if deck.front == -1 && deck.rear == -1 {
		return -1, errors.New("Fila Vazia")
	}
	return deck.v[deck.front], nil
}

func (deck *ArrayDequeue) Back() (int, error) {
	if deck.front == -1 && deck.rear == -1 {
		return -1, errors.New("Fila Vazia")
	}
	return deck.v[deck.rear], nil
}

func (deck *ArrayDequeue) print() {
	sz := deck.size()
	if sz == 0 {
		fmt.Println("Deque vazio")
		return
	}
	i := deck.front
	for count := 0; count < sz; count++ {
		fmt.Printf("%d ", deck.v[i])
		i = (i + 1) % len(deck.v)
	}
	fmt.Println()
}

func main() {
	d := &ArrayDequeue{}
	d.init(4) // inicializa com capacidade 4

	// Teste inserções push_back
	d.push_back(10)
	d.push_back(20)
	d.push_back(30)
	fmt.Printf("Deque após push_back: size=%d\n", d.size())
	f, _ := d.Front()
	b, _ := d.Back()
	fmt.Println("Front:", f, "Back:", b)

	// Teste inserção push_front
	d.push_front(5)
	fmt.Printf("Deque após push_front(5): size=%d\n", d.size())
	f, _ = d.Front()
	b, _ = d.Back()
	fmt.Println("Front:", f, "Back:", b)

	// Forçar o doubleV
	d.push_back(40)
	d.push_back(50)
	fmt.Printf("Deque após push_back (dobrar capacidade): size=%d\n", d.size())
	f, _ = d.Front()
	b, _ = d.Back()
	fmt.Println("Front:", f, "Back:", b)

	// Teste pull_front
	val, err := d.pull_front()
	if err == nil {
		fmt.Println("pull_front:", val)
	}
	fmt.Printf("Deque após pull_front: size=%d\n", d.size())

	// Teste pull_back
	val, err = d.pull_back()
	if err == nil {
		fmt.Println("pull_back:", val)
	}
	fmt.Printf("Deque após pull_back: size=%d\n", d.size())

	d.print()
}
