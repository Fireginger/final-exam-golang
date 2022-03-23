package piscine

import "github.com/01-edu/z01"

func Diamond(index int) {
	if index == 1 {
		z01.PrintRune('*')
	}
	if index < 1 {
		z01.PrintRune('\n')
	}
	count := index - 1
	count2 := 0
	if index > 1 {
		for i := count; i >= 0; i-- {
			z01.PrintRune(' ')
		}
		z01.PrintRune('*')
		z01.PrintRune(10)
	}
	for i := index - 1; i > 0; i-- {
		for j := count - 1; j >= 0; j-- {
			z01.PrintRune(' ')
		}
		count--
		count2++
		z01.PrintRune('*')
		for k := count2 - 1; k >= 0; k-- {
			z01.PrintRune(' ')
		}
		count2++
		z01.PrintRune('*')
		z01.PrintRune(10)
	}
	count3 := index - 1
	count4 := index - 1
	if index > 1 {
		z01.PrintRune(' ')
		z01.PrintRune(' ')
		z01.PrintRune('*')
		for i := count3; i >= 0; i-- {
			z01.PrintRune(' ')
		}
		z01.PrintRune('*')
		z01.PrintRune(10)
	}
	for i := 0; i < index-2; i++ {
		for j := count3 - 2; j >= 0; j-- {
			z01.PrintRune(' ')
		}
		count3++
		count4--
		z01.PrintRune('*')
		for k := count4 - 1; k >= 0; k-- {
			z01.PrintRune(' ')
		}
		count4--
		z01.PrintRune('*')
		z01.PrintRune(10)
	}
}
