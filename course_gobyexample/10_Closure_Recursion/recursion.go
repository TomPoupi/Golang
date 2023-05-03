package main

import "fmt"

// fonciton qui fait n!
func fact(n int) int {
	if n == 0 {
		return 1 //condition de sortie
	}
	return n * fact(n-1) // boucle recursive
}

func main() {
	fmt.Println(fact(7))

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n // condition de sortie
		}

		return fib(n-1) + fib(n-2) // boucle recursive
	}

	fmt.Println(fib(7))
}
