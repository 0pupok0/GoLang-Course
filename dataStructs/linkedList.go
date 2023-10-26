package dataStructs

import (
	"errors"
	"fmt"
)

type node struct {
	data int
	next *node
}

type LinkedList struct {
	head *node
	tail *node
}

// New Инициализирует новый список. Количество элементов q
func New(q int) LinkedList {
	var l LinkedList
	for i := 0; i < q; i++ {
		l.Add(0)
	}
	return l
}

// Add Добавляет элемент в конец списка
func (l *LinkedList) Add(val int) {
	var newNode node
	newNode.data = val
	if l.tail == nil {
		l.head = &newNode
		l.tail = &newNode
	} else {
		l.tail.next = &newNode
		l.tail = &newNode
	}
}

// Pop Удаляет элемент из конца списка
func (l *LinkedList) Pop() error {
	el := l.head
	if el == nil {
		return errors.New("no elements")
	}
	if el.next == nil {
		l.head = nil
		l.tail = nil
		return nil
	}
	for el.next != l.tail {
		el = el.next
	}
	el.next = nil
	l.tail = el
	return nil
}

// At Получает значение на позиции pos
func (l *LinkedList) At(pos int) (int, error) {
	el := l.head
	if pos < 0 {
		return 0, errors.New("no negative pos")
	}
	for i := 0; i < pos; i++ {
		if el == nil {
			return 0, errors.New("no element on such index")
		}
		el = el.next
	}
	return el.data, nil
}

// Size Возвращает длину списка
func (l *LinkedList) Size() int {
	counter := 0
	el := l.head
	for el != nil {
		counter++
		el = el.next
	}
	return counter
}

// DeleteFrom Удаляет элемент на позиции pos
func (l *LinkedList) DeleteFrom(pos int) error {
	el := l.head
	for i := 0; i < pos; i++ {
		if el == nil {
			return errors.New("no element on such index")
		}
		el = el.next
	}
	if el.next == nil {
		return l.Pop()
	}
	el.next = el.next.next
	return nil
}

// UpdateAt Делает значение на позиции pos равным val
func (l *LinkedList) UpdateAt(pos, val int) error {
	el := l.head
	if pos < 0 {
		return errors.New("no negative pos")
	}
	for i := 0; i <= pos; i++ {
		if el == nil {
			return errors.New("no element on such index")
		}
		el = el.next
	}
	el.data = val
	return nil
}

// NewFromSlice Создаёт связанный список из среза
func NewFromSlice(s []int) LinkedList {
	var l LinkedList
	for _, v := range s {
		l.Add(v)
	}
	return l
}

// Println Выводит список
func (l *LinkedList) Println() {
	fmt.Print("[")
	for ptr := l.head; ptr != nil; ptr = ptr.next {
		fmt.Print(ptr.data, ", ")
	}
	fmt.Println("]")
}
