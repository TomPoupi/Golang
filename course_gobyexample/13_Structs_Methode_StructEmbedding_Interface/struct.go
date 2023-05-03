package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {

	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	// on est pas obliger de tout déclarer qu'on fait appel à un objet d'une struct, elle seront mis par default si rien n'est initialiser

	// on n'est pas obliger de préciser les noms des attibuts
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{age: 40})

	//opérations classique sur les structures
	p := newPerson("Jon")
	fmt.Println(*p)
	fmt.Println(p.name)

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}
