package piscine

func Decipher(s string) string {
	phrase := ""
	tableau := []rune(s)
	for i := 0; i < len(tableau); i++ {
		if s[i] < 91 && s[i] > 64 {
			if s[i]-16 > 64 {
				phrase = phrase + string(s[i]-16)
			} else if s[i]-16 < 64 {
				tableau[i] = tableau[i] - 16
				tableau[i] = tableau[i] + 26
				phrase = phrase + string(tableau[i])
			}
		}
		if s[i] != 32 {
			if s[i]-16 > 96 {
				phrase = phrase + string(s[i]-16)
			} else if s[i]-16 < 96 && s[i]-16 > 80 {
				tableau[i] = tableau[i] - 16
				tableau[i] = tableau[i] + 26
				phrase = phrase + string(tableau[i])
			}
		} else {
			phrase = phrase + " "
		}
	}
	return phrase
}
