package main

import (
	"errors"
	"fmt"
)

type node struct {
	val   int
	left  *node
	right *node
}

type Binary_tree struct {
	root     *node
	inserted int
}

type binary_tree interface {
	Add(val int) //feito
	Remove(val int)
	Height() int         //feito
	Min() (int, error)   //feito
	Max() (int, error)   //feito
	InOrder()            //feito
	PreOrder()           //feito
	PosOrder()           //feito
	LevelOrder()         //feito
	Search(val int) bool //feito
}

// função auxiliar para criar noh
func createNode(val int) *node {
	return &node{val: val, left: nil, right: nil}
}

// função add
func (bst *Binary_tree) Add(val int) {
	if bst.root == nil {
		bst.root = createNode(val)
	} else {
		bst.root.AddNode(val)
	}
	bst.inserted++
}

func (no *node) AddNode(val int) {
	if val <= no.val { // se o valor for menor adiciono na parte no noh filho da esquerda
		if no.left == nil { // se eu tiver na folha eu adiciono o novo noh
			no.left = createNode(val)
		} else {
			no.left.AddNode(val) //se n eu percorro pra esquerda dnv e adciono o noh na folha
		}
	} else {
		if no.right == nil {
			no.right = createNode(val)
		} else {
			no.right.AddNode(val)
		}
	}
}

// função search
func (bst *Binary_tree) Search(val int) bool {
	if bst.root == nil {
		return false
	} else {
		return bst.root.SearchNode(val)
	}
}

func (no *node) SearchNode(val int) bool {
	if no.val == val {
		return true
	} else if val < no.val {
		if no.left == nil {
			return false
		} else {
			return no.left.SearchNode(val)
		}
	} else {
		if no.right == nil {
			return false
		} else {
			return no.right.SearchNode(val)
		}
	}
}

// funções max e min
func (bst *Binary_tree) Min() (int, error) {
	node := bst.root
	if bst.root == nil { // a arvore é vazia, se for retorna erro se não vou percorrer a arvore até no no.left ser nil
		return -1, errors.New("Arvore vazia")
	} else {
		for node.left != nil {
			node = node.left
		}
		return node.val, nil
	}
}

func (bst *Binary_tree) Max() (int, error) {
	node := bst.root
	if bst.root == nil {
		return -1, errors.New("Arvore vazia")
	} else {
		for node.right != nil {
			node = node.right
		}
		return node.val, nil
	}
}

func (no *node) Min() int {
	for no.left != nil {
		no = no.left
	}
	return no.val
}

//função height

func (bst *Binary_tree) Height() int {
	if bst.root == nil {
		return 0
	} else {
		return bst.root.Height()
	}
}

func (no *node) Height() int {
	if no.left == nil && no.right == nil { //caso base da recursão (se no.left e no.right forem nil)
		return 0
	}
	hleft := 0 //crio uma variavel hleft e se no.left != nil eu retorno 1 + no.left.Height()
	if no.left != nil {
		hleft = 1 + no.left.Height()
	}
	hright := 0 //mesma coisa aqui
	if no.right != nil {
		hright = 1 + no.right.Height()
	}
	if hleft >= hright { //retorno a altura maior do dois
		return hleft
	} else {
		return hright
	}
}

//funcoes order

func (no *node) PreOrder() {
	fmt.Print(no.val)
	if no.left != nil {
		no.left.PreOrder()
	}
	if no.right != nil {
		no.right.PreOrder()
	}
}

func (no *node) InOrder() {
	if no.left != nil {
		no.left.InOrder()
	}
	fmt.Print(no.val)
	if no.right != nil {
		no.right.InOrder()
	}
}

func (no *node) PosOrder() {
	if no.left != nil {
		no.left.PosOrder()
	}
	if no.right != nil {
		no.right.PosOrder()
	}
	fmt.Print(no.val)
}

func (no *node) LevelOrder() { // se noh é nill retorno
	if no == nil {
		return
	}
	queue := []*node{no} //crio o vetor e coloco noh nele
	for len(queue) > 0 { // enquanto tiver elementos na fila
		n := len(queue) // numero de elementos a serem processados
		for i := 0; i < n; i++ {
			no := queue[0]           //coloco o primeiro elemento em no
			fmt.Printf("%d", no.val) //processo
			queue = queue[1:]        //tiro da fila
			if no.left != nil {      //coloco os elementos a esqueda e a direita de no na fila se existirem para o processamento
				queue = append(queue, no.left)
			}
			if no.right != nil {
				queue = append(queue, no.right)
			}
		}
		fmt.Println()
	}
}

// função remove
func (bst *Binary_tree) Remove(val int) bool {
	if bst.root == nil {
		return false
	} else {
		var removed bool
		bst.root, removed = bst.root.Remove(val)
		if removed {
			bst.inserted--
		}
		return removed
	}
}

func (no *node) Remove(val int) (*node, bool) {
	if no == nil { //se noh atual é nil retorno false e não foi encontrado o valor
		return nil, false
	}
	var removed bool
	if val < no.val { //navego em toda a arvore procurando o noh e atualizando o removed
		no.left, removed = no.left.Remove(val)
	} else if val > no.val {
		no.right, removed = no.right.Remove(val)
	} else {
		//noh encontrado
		removed = true
		// se ele é um noh folha eu retorno nil para o lugar seu pai
		if no.left == nil && no.right == nil {
			return nil, removed
		} else if no.left != nil && no.right == nil { // se ele é um noh com um filho só eu retorno seu unico filho para o lugar de seu pai
			return no.left, removed
		} else if no.left == nil && no.right != nil { //mesmo caso
			return no.right, removed
		} else { // se ele tem dois filhos
			min := no.right.Min()              //procuro o menor valor da subarvore da direta do noh que eu quero remover
			no.val = min                       //atualizo o valor dele para o valor do menor
			no.right, _ = no.right.Remove(min) // removo o menor da subarvore direita com uma chamada recursiva
		}
	}
	return no, removed //retorno o no e o bool se foi removido ou não
}

// func isBst
func (no *node) IsBst(min int, max int) bool {
	if no == nil { //se no == nil returna true
		return true
	}
	if no.val >= max || no.val < min { //verifica se o noh está dentro do intervalo
		return false //se estiver fora retorna false
	}
	return no.left.IsBst(min, no.val) && no.right.IsBst(no.val, max)
	//retorna que todos os valores da esquerda tem q ser menores que no.val
	//e todos os valores a direita maiores que no.val
}

func (no *node) Size() int {
	countL := 0
	countR := 0
	if no.left != nil {
		countL += 1 + no.left.Size()
	}
	if no.right != nil {
		countR += 1 + no.right.Size()
	}
	return 1 + countR + countL
}

// func para transformar um vetor em uma arvore
func convertToBalancedBst(v []int, ini int, fim int) *node {
	if ini > fim {
		return nil
	}
	mid := (ini + fim) / 2
	no := &node{
		val:   v[mid],
		left:  convertToBalancedBst(v, ini, mid-1),
		right: convertToBalancedBst(v, mid+1, fim),
	}
	return no
}

// func para contar o numero de pares numa arvore
func (no *node) Par() int {
	countL := 0
	countR := 0
	if no == nil {
		return 0
	}
	if no.left != nil {
		countL += no.left.Par()
	}
	if no.right != nil {
		countR += no.right.Par()
	}
	if no.val%2 == 0 {
		return 1 + countL + countR
	} else {
		return countL + countR
	}
}
