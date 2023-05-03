package main

import "fmt"

// on créer la fonction plus qui prends en entré 2 int (a int ,b int) et return un int
func plus(a int, b int) int {

	return a + b
}

// on peut l'écrire de cette manière , globaliser les int si a b c doit être des int en entré
func plusPlus(a, b, c int, d bool) int {
	if d == true {
		fmt.Println("ça marche")
	}
	return a + b + c
}

func main() {

	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3, true)
	fmt.Println("1+2+3 =", res)
}
