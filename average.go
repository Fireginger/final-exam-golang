package piscine

func Average(tab []int) int {
	count := 0
	moyenne := 0
	for i := 0; i < len(tab); i++ {
		count++
		if tab[i] == -1 {
			i++
		}
		moyenne = moyenne + tab[i]
	}
	moyenne = moyenne / count
	return moyenne
}
