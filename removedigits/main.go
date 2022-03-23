package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	phrase := "Arguments "
	arguments := os.Args
	var count rune = '0'
	if len(arguments) < 1 {
		return
	}
	for i := 1; i < len(arguments); i++ {
		count++
		for _, mot := range arguments[i] {
			if mot >= 48 && mot <= 57 {
				i++
				count++
			}
			if mot+1 >= 48 && mot+1 <= 57 {
				i++
				count++
			}
		}
		for i := 0; i < len(phrase); i++ {
			z01.PrintRune(rune(phrase[i]))
		}
		z01.PrintRune(count)
		z01.PrintRune(' ')
		z01.PrintRune(':')
		z01.PrintRune(' ')
		for j := 0; j < len(arguments[i]); j++ {
			tableau := []rune(arguments[i])
			z01.PrintRune(tableau[j])
		}
		z01.PrintRune(10)
	}
}
