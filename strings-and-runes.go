package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	const s = "สวัสดี"

	// fmt.Println(s)
	// fmt.Println("strings == []bytes,so the value below represents bytes\n", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	fmt.Println("Rune Count", utf8.RuneCountInString(s))

	for index,runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, index)
	}

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, width)
		w = width
		examineRune(runeValue)
	}


	// ch := "你好"

	// fmt.Println(ch)
	// fmt.Println("len of raw bytes: ", len(ch))
}

func examineRune(r rune) {

	if r == 't' {
		fmt.Println("Found t!")
	} else if r == 'ส' {
		fmt.Println("Found so sua!")
		fmt.Printf("%#U\n", r)
	}
}

