package main

import (
	"fmt"
)

// le ...int : j'accepte en entré autant élement int que tu veux même des tableau de int et je le met dans nums
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

// je peux faire quelque chose d'encore plus what the fuck
// n'importe quelle entrée de n'importe quel type
// BOOOM
func nimp(nums ...interface{}) {
	fmt.Println(nums, "")

	for _, num := range nums {
		fmt.Printf("%T ", num)
	}
	fmt.Println("")

}

func main() {

	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	// il faut faire slices... pour mettre en entrée un slice
	sum(nums...)

	nimp(3, "cocou", true)
	tabnimp := make([]interface{}, 3)
	tabnimp[0] = "coucou"
	tabnimp[1] = 2
	tabnimp[2] = false
	nimp(tabnimp...)

}
