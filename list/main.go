package main

import (
	list "example/list/storage"
	"fmt"
)

func main() {
	l := list.NewList()

	l.Add(4)
	l.Add(8)
	l.Add(34)
	l.Add(8)
	l.Add(7)
	l.Add(79)
	l.Add(117)
	l.Add(4)
	l.Add(56)
	l.Add(12)

	fmt.Println("Длина листа:", l.Len())

	values, ok := l.GetAll()
	if ok {
		fmt.Println("Элементы листа:", values)
	} else {
		fmt.Println("Пусто")
	}

	l.RemoveByIndex(2)
	fmt.Println("Длина листа после удаления элемента по индексу:", l.Len())

	values, ok = l.GetAll()
	if ok {
		fmt.Println("Элементы в листе после удаления:", values)
	} else {
		fmt.Println("Пусто")
	}
	l.RemoveAllByValue(8)
	l.RemoveByValue(56)
	values, ok = l.GetAll()
	if ok {
		fmt.Println("Элементы в листе после удаления по значению:", values)
	} else {
		fmt.Println("Пусто")
	}

	fmt.Println(l.GetByIndex(3))
	fmt.Println(l.GetByValue(117))
	fmt.Println(l.GetAllByValue(12))

	fmt.Println("Вывод элементов листа:")
	l.Print()

	l.Clear()
	fmt.Println("Длина листа после очистки:", l.Len())
}
