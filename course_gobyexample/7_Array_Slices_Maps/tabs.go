package main

import "fmt"

func main() {
	//------------------------------------------------------
	fmt.Println("")
	fmt.Println("Array--------------------------------")
	var d [5]string //... string " "
	fmt.Println("emp perso:", d)

	var a [5]int // initialiser un tableau vide int : 0 de taille 5
	fmt.Println("emp:", a)

	a[4] = 100 // comment remplir un tabs manière classique
	fmt.Println("set:", a)
	fmt.Println("get:", a[4]) // selectionner l'element 4 de a

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5} // manière de remplir un tableau
	fmt.Println("dcl:", b)

	//on ne peut pas append un Array il faut use slice

	var twoD [2][3]int // tableau 3 lignes 2 colonne (ou on peut le voir 2 tab de 3 element)
	for i := 0; i <= 1; i++ {
		for j := 0; j <= 2; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d: ", twoD)

	fmt.Println("other view tab 2D: ")
	for i := 0; i <= 1; i++ {
		if i != 0 {
			fmt.Println("")
		}
		for j := 0; j <= 2; j++ {
			fmt.Print(twoD[i][j], " ")
		}
	}
	fmt.Println("")
	fmt.Println("--------------------------------")
	//------------------------------------------------------

	//------------------------------------------------------
	fmt.Println("")
	fmt.Println("Slices--------------------------------")

	s := make([]string, 3) // make permet de faire des slices, []string : tableau 1D string, 3 : 3elements
	fmt.Println("emp:", s)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	s = append(s, "d") // là on peut append
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s) // on peut copier des élement de s dans c
	fmt.Println("cpy:", c)
	m := make([][]string, 0)
	m = append(m, c)
	m = append(m, s)
	fmt.Println("append_tab:", m)

	l := s[2:5] // prendre les élement de la 2 jusqu'à la 5
	fmt.Println("sl1:", l)

	l = s[:5] // de la 0 jusqu'à la 5
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"} // avec cette écriture la différence avec array c'est qu'on définit une taille limite , slice prends autant d'élement qu'on veut
	fmt.Println("dcl:", t)

	twoD2 := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD2[i] = make([]int, innerLen) // faire tableau de tableau
		for j := 0; j < innerLen; j++ {
			twoD2[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD2)

	fmt.Println("")
	fmt.Println("--------------------------------")
	//------------------------------------------------------

	//------------------------------------------------------
	fmt.Println("")
	fmt.Println("Maps--------------------------------")

	e := make(map[string]int) // init map , map : key, element

	e["k1"] = 7  // key: k1 = element: 7
	e["k2"] = 13 // key: k2 = element:8

	fmt.Println("map:", e)

	v1 := e["k1"] // selectionner élement de la key "k1"
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(e))

	delete(e, "k2") // on peut delete l'élement de la key "k2"
	fmt.Println("map:", e)

	_, prs := e["k2"]        // _,prs permet de mettre élement e["k2"] dans pers
	fmt.Println("prs:", prs) // sauf qu'il n'y pas de key k2 ou d'element associer à k2 renvoie false
	key, prs := e["k1"]
	fmt.Println("key:", key, prs)

	n := map[string]int{"foo": 1, "bar": 2} // on peut aussi init comme ça
	fmt.Println("map:", n)

	fmt.Println("")
	fmt.Println("--------------------------------")

	//------------------------------------------------------
}
