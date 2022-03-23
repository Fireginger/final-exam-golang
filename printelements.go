package piscine

func PrintElements(tab []int, str []string) string {
	phrase := ""
	for i := 0; i < len(tab); i++ {
		if tab[i] > len(str) {
			phrase = phrase + "!"
			if tab[i] != len(tab)-1 {
				phrase = phrase + " "
			}
		} else {
			phrase = phrase + str[tab[i]]
			phrase = phrase + " "
		}
	}
	phrase += "\n"
	return phrase
}
