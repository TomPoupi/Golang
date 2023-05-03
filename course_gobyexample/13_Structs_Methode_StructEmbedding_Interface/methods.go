package main

import "fmt"

type rect struct {
	width, height int
}

// methode fonction intégré dans la struct plutot pratique
// il faut préciser (r *rect) avant le noms de la foncitons
// au sinon fonction classique on peut mettre des truc en entre
// même si c'est pas montré

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
