package Verify

func VerifOneTap(Letter string, ListWordCap []string) bool { //Fonction qui vérifie si la tentative de mettre le mot en entier est réusite ou pas
	var Box string
	if len(Letter) > 1 {
		for y := range ListWordCap {
			Box += ListWordCap[y]
			if Letter == Box {
				return true
			}
		}
	}
	return false
}
