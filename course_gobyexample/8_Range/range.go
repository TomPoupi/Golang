package main

import "fmt"

func main() {

	nums := []int{2, 3, 4}
	sum := 0
	sumi := 0
	// range comme dans python de parcourir les élement d'un slice ou Array
	// le for i, num : le i permet d'avoir les itération et le num les élements
	for i, num := range nums {
		sum += num
		sumi += i
	}
	fmt.Println("les sums:", sum, sumi)

	sum = 0
	// si tu t'en fous des indew tu peux mettre _, et inversement
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum: ", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// tu peux range sur une map
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v) // %s c'est pour les strings
	}

	//prendre uniquement les keys
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// range sur un string permet d'avoir unicode des caractère (rune)
	// i : index of rune
	// c : rune
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
