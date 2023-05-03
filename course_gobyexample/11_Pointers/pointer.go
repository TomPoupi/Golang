package main

import "fmt"

// prend une variable et la change à 0 sans la return
func zeroval(ival int) {
	ival = 0
}

// la valeur pointer d'une address d'une variable et change la valeur pointer
// ce qui change ça valeur même s'il n'y a pas de return
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	//zeroval ne change pas la valeur
	zeroval(i)
	fmt.Println("zeroval:", i)

	//zeroprt change la pointer valeur par l'address de &i
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// la valeur de l'address
	fmt.Println("pointer:", &i)
}
