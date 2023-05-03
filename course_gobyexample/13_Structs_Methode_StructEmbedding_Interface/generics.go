package main

import "fmt"

// permet de préciser le typage en entré
// comme ça je peux prendre n'importe quelle type de map entré
func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

// permet de mettre n'importe quelle type délement dans ma liste
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var m = map[int]string{1: "2", 2: "4", 4: "8"}

	fmt.Println("keys:", MapKeys(m))

	//fmt.Println(MapKeys[int, string](m))

	//je peux définir directement le type que va utilser ma structure en entré
	lst := List[any]{}
	lst.Push(10)
	lst.Push("13")
	lst.Push(true)
	fmt.Println("list:", lst.GetAll())
	fmt.Println("list:", lst.tail.val)
}
