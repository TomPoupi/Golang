package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// rune = charactere in golang , c'est un int qui represente un code Unicode
	const s = "สวัสดี" // thai charac

	fmt.Println("Len:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i]) // voir les byte en hexa du string s
	}
	fmt.Println()

	// compte le nombre de rune, Some Thai characters are represented by multiple UTF-8 code points
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// par exemple rune value : U+0E2A
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}

func examineRune(r rune) {

	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
