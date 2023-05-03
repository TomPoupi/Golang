package main

import "fmt"

func main() {
	i := 1
	//boucle while
	fmt.Println("Boucle while :")
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	fmt.Println(" ")
	fmt.Println("Boucle while avec break :")
	r := 0
	for {
		fmt.Println("loop")
		r = r + 1
		if r == 4 {
			break // break quand r = 4
		}
	}

	fmt.Println(" ")
	fmt.Println("Boucle for classique :")
	for j := 0; j <= 5; j++ {

		fmt.Println(j)
	}

	fmt.Println(" ")
	fmt.Println("Boucle for avec continue:")
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue // skip la boucle actuelle si je suis pair
		}
		fmt.Println(n)
	}

}
