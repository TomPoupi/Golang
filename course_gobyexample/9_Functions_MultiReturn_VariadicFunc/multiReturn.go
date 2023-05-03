package main

import "fmt"

// on peut return plusieur variable
// écriture spéciale vu qu'on return directement pas besoin de mettre de variable juste le type suffit
// 1er int = 3 2nd int =7
func vals() (int, int) {
	return 3, 7
}

func main() {

	//a=3 et b=7
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// si on veut uniquement le 2nd return il faut précisé le _,var
	// c = 7
	_, c := vals()
	fmt.Println(c)
}
