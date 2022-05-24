package main

import (
	"bufio"
	"fmt"
	"os"
)

func translate(text string) {
	var word string
	for j := 0; j < len(text); j++ {
		c := text[j]
		switch c {
		case 'a':
			word = word + ".- "
		case 'b':
			word = word + "-... "
		case 'c':
			word = word + "-.-. "
		case 'd':
			word = word + "-.. "
		case 'e':
			word = word + ". "
		case 'f':
			word = word + "..-. "
		case 'g':
			word = word + "--. "
		case 'h':
			word = word + ".... "
		case 'i':
			word = word + ".. "
		case 'j':
			word = word + ".--- "
		case 'k':
			word = word + "-.- "
		case 'l':
			word = word + ".-.. "
		case 'm':
			word = word + "-- "
		case 'n':
			word = word + "-. "
		case 'o':
			word = word + "--- "
		case 'p':
			word = word + ".--. "
		case 'q':
			word = word + "--.- "
		case 'r':
			word = word + ".-. "
		case 's':
			word = word + "... "
		case 't':
			word = word + "- "
		case 'u':
			word = word + "..- "
		case 'w':
			word = word + ".-- "
		case 'x':
			word = word + "-..- "
		case 'y':
			word = word + "-.-- "
		case 'z':
			word = word + "--. "
		case ' ':
			word = word + "\n"
		default:
			fmt.Printf("введите фразу на латинице\n")
		}
	}
	fmt.Println(word)
}
func main() {
	var text string
	fmt.Println("Введите слово на английском для перевода")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		text = scanner.Text()
	}
	fmt.Printf("Input was: %q\n", text)
	translate(text)
}
