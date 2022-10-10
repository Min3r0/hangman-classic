package Verify

func VerifOneTap(letter string, ListWordCap []string) bool { //Fonction qui vérifie si la tentative de mettre le mot en entier est réusite ou pas
	var box string
	if len(letter) > 1 {
		for y := range ListWordCap {
			box += ListWordCap[y]
			if letter == box {
				return true
			}
		}
	}
	return false
}
