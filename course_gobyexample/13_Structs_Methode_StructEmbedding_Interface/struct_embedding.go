package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

func (b base) desc() {
	fmt.Printf("base with num=%v", b.num)
}

// on fait un struct qui fait appel un struct comme args
type container struct {
	base
	str string
}

func main() {

	co := container{
		base: base{ // on creer en arg de base l'objet base
			num: 1,
		},
		str: "some name",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	fmt.Println("also num:", co.base.num) // co.base pour avoir objet base puis .num pour avoir le int assicier à l'arg num de base

	fmt.Println("describe:", co.describe()) // pas besoin de préciser le .base pour utliser la méthode describe()
	// Since container embeds base,
	//the methods of base also become methods of a container.

	// interface comme on a vu avant, ça marche aussi avec les interface
	type describer interface {
		describe() string
		desc()
	}

	// on associe notre objet container à notre decriber
	var d describer = co
	fmt.Println("describer:", d.describe())
	d.desc()
}
