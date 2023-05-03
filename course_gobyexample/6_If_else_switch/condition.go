package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println(" ")
	fmt.Println("if else-------------------------------")
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := -3; num < 0 { //comme si je faisais num = -3 puis je fais mon if mais construction possible très bizarre ( init dans le if)
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
	fmt.Println("-------------------------------")

	fmt.Println(" ")
	fmt.Println("Switch-------------------------------")

	//switch classique
	p := 2
	fmt.Print("Write ", p, " as ")
	switch p {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday: // on peut fusioner des cases pour leur donner un même résultat
		fmt.Println(time.Now().Weekday(), ", It's the weekend")
	default:
		fmt.Println(time.Now().Weekday(), ", It's a weekday") // le défault correspond à toute les autres réponses
	}

	t := time.Now()
	switch {
	case t.Hour() < 12: // on peut mettre des condition à la if dans les switch case
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
	// je créer la fonction whatAmI qui prend n'importe quelle type
	// en entré "interface{}" et donne en print le type de la variable
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Other type %T\n", t) // le %T ça permet d'incruster type Type dans une phrase
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
	fmt.Println("-------------------------------")
}
