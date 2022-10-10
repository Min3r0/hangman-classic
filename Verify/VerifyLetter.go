package Verify

func VerifyLetter(letter string, ListWordCap []string) []int { // Fonction vérifie si la lettre choisi et bien dans le mot et renvois une liste d'entier composer de la position de la lettre.
	var list []int
	for i := range ListWordCap {
		if letter == ListWordCap[i] {
			list = append(list, i)
		}
	}
	return list
}
