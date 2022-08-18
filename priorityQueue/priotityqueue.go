package main

import (
	"fmt"
)

type QueueElement struct {
	value    interface{} //Valor
	priority int         //Prioridad
}

type PriorityQueue struct {
	maxSize int
	data    []QueueElement
	size    int
}

//estbalece el maximo tamaño de la priority queue
func (p *PriorityQueue) SetMaxSize(maxSize int) {
	p.maxSize = maxSize
}

// comprueba que la cola no este vacia (colaprioridadvacia)
func (p *PriorityQueue) IsEmpty() bool {
	return p.size == 0
}

// Insertar un nuevo elemento haciendo uso de su prioridad(inseEnprioridad)
func (p *PriorityQueue) Enqueue(el interface{}, priority int) {
	newElement := QueueElement{el, priority}
	if p.size == p.maxSize {
		panic("La cola ha superado su maxima capacidad.")
	}

	if p.IsEmpty() {
		p.data = append(p.data, newElement)
		p.size++
		return
	}

	p.data = append(p.data, newElement)
	i := p.size - 1
	for i >= 0 {
		if p.data[i].priority > priority {
			p.data[i+1] = p.data[i]
			i--
		} else {
			break
		}
	}
	p.data[i+1] = newElement
	p.size++
}

// elimina y retorna el primer elemeno de la priority enqueue (elementoMin y quitarMin)
func (p *PriorityQueue) Dequeue() QueueElement {
	if p.IsEmpty() {
		panic("La cola se encuentra vacia.")
	}
	dequeued := p.data[0]

	p.data = p.data[1:]
	p.size--
	return dequeued
}

// Retorna el primer elemento sin modificar la cola
func (p *PriorityQueue) Peek() interface{} {
	if p.IsEmpty() {
		panic("La cola se encuentra vacia.")
	}

	return p.data[0]
}

// muestra los elementos en  un arary
func (p *PriorityQueue) Display() {
	if p.IsEmpty() {
		panic("La cola se encuentra vacia.")
	}

	arr := make([]interface{}, p.size)
	i := 0
	for i < p.size {
		arr[i] = p.data[i].value
		i++
	}
	fmt.Println("elementos de la priority queue: ", arr)
}

func main() {
	pq := PriorityQueue{}
	pq.SetMaxSize(12)
	pq.Enqueue(2, 3)
	pq.Enqueue(6, 7)
	pq.Enqueue(6, 9)
	pq.Enqueue(12, 10)
	pq.Enqueue(15, 4)
	pq.Enqueue(9, 12)

	pq.Display()
	fmt.Println(pq.Dequeue())
	fmt.Println(pq.Dequeue())
	fmt.Println(pq.Dequeue())
	fmt.Println(pq.Dequeue())
	fmt.Println(pq.Peek())
	pq.Display()
}
