package list

import (
	"fmt"
)

type List struct {
	firstnode *node
	len       int64
}

// NewList создает новый список
func NewList() (l *List) {
	return &List{}
}

// Len возвращает длину списка
func (l *List) Len() (len int64) {
	return l.len
}

// Add добавляет элемент в список и возвращает его индекс
func (l *List) Add(value int64) (id int64) {
	newNode := &node{value: value}
	if l.firstnode == nil {
		l.firstnode = newNode
		l.len++
		l.firstnode.index = l.len - 1
		return l.firstnode.index
	}
	cur := l.firstnode
	for ; cur.next != nil; cur = cur.next {
	}
	cur.next = newNode
	l.len++
	cur.next.index = l.len - 1
	return cur.next.index
}

// RemoveByIndex удаляет элемент из списка по индексу
func (l *List) RemoveByIndex(id int64) {
	if l.firstnode == nil {
		return
	}
	if id == 0 {
		l.firstnode = l.firstnode.next
		l.len--
		return
	}
	cur := l.firstnode
	for cur.next != nil && cur.next.index != id {
		cur = cur.next
	}
	cur.next = cur.next.next
	l.len--
	return
}

// RemoveByValue удаляет элемент из списка по значению
func (l *List) RemoveByValue(value int64) {
	if l.firstnode == nil {
		return
	}
	if l.firstnode.value == value {
		l.firstnode = l.firstnode.next
		l.len--
		return
	}
	cur := l.firstnode
	for cur.next != nil && cur.next.value != value {
		cur = cur.next
	}
	cur.next = cur.next.next
	l.len--

	return
}

// RemoveAllByValue удаляет все элементы из списка по значению
func (l *List) RemoveAllByValue(value int64) {
	if l.firstnode == nil {
		return
	}
	if l.firstnode != nil && l.firstnode.value == value {
		l.firstnode = l.firstnode.next
		l.len--
		return
	}
	cur := l.firstnode
	for cur != nil && cur.next != nil {
		if cur.next.value == value {
			cur.next = cur.next.next
			l.len--
		}
		cur = cur.next
	}
	return
}

// GetByIndex возвращает значение элемента по индексу.
//
// Если элемента с таким индексом нет, то возвращается 0 и false.
func (l *List) GetByIndex(id int64) (value int64, ok bool) {
	if id < 0 {
		return 0, false
	}
	cur := l.firstnode
	for cur != nil {
		if id == 0 {
			return cur.value, true
		}
		cur = cur.next
		id++
	}
	return 0, false
}

// GetByValue возвращает индекс первого найденного элемента по значению.
//
// Если элемента с таким значением нет, то возвращается 0 и false.
func (l *List) GetByValue(value int64) (id int64, ok bool) {
	cur := l.firstnode
	id = 0
	for cur != nil {
		if cur.value == value {
			return id, true
		}
		cur = cur.next
		id++
	}

	return 0, false
}

// GetAllByValue возвращает индексы всех найденных элементов по значению
//
// Если элементов с таким значением нет, то возвращается nil и false.
func (l *List) GetAllByValue(value int64) (ids []int64, ok bool) {
	cur := l.firstnode
	for cur != nil {
		if cur.value == value {
			ids = append(ids, cur.index)
		}
		cur = cur.next
	}
	if len(ids) == 0 {
		return nil, false
	}

	return ids, true
}

// GetAll возвращает все элементы списка
//
// Если список пуст, то возвращается nil и false.
func (l *List) GetAll() (values []int64, ok bool) {
	if l.len == 0 {
		return nil, false
	}
	cur := l.firstnode
	for cur != nil {
		values = append(values, cur.value)
		cur = cur.next
	}
	return values, true
}

// Clear очищает список
func (l *List) Clear() {
	l.firstnode = nil
	l.len = 0
}

// Print выводит список в консоль
func (l *List) Print() {
	if l.len == 0 {
		return
	}
	fmt.Print("My List: ")
	cur := l.firstnode
	for cur != nil {
		fmt.Printf("%d:%d, ", cur.index, cur.value)
		cur = cur.next
	}

	fmt.Println()
}
