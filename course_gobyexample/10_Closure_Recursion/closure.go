package main

import "fmt"

// la fonction return une fonction closure ( qui n'a pas de nom )
// une des application est de pouvoir garder en mémoire une variable local ( ici i )
// et pour qu'à chaque execusion elle ne soit réinitialiser
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	// nextInt à comme valeur une fonction closure
	nextInt := intSeq()

	// chaque fois qu'on appelle nextInt on change la valeur de i qui n'est jamais re-initialiser
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}
