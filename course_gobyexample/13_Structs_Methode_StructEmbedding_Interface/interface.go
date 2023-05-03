package main

import (
	"fmt"
	"math"
)

// les interfaces permet de centraliser les méthode
type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// on a juste a appeler notre interface pour avoir les méthodes
// comme quand on appelle la fonction measure on aura juste à mettre la variable de l'objet créer par notre struct
// en entré pour avoir les méthode associer

// différente utilisation des méthodes

func measure() func(g geometry) {

	i := 0
	return func(g geometry) { // variante personnelle j'ai fait une closure pour incrémenté i en interne
		fmt.Println("forme géométrique n°", i)
		fmt.Println(g)
		fmt.Println(g.area())
		fmt.Println(g.perim())
		i = i + 1
	}
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	p := measure()
	p(r)
	p(c)
}
