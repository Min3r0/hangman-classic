package CreateList

func CreateListWordCap(Word string) []string { //Fonction qui retourne une liste de lettre en majuscule
	var ListLettre []string
	for i := 0; i < len(Word); i++ {
		ListLettre = append(ListLettre, string(Word[i]-32))
	}
	return ListLettre
}
