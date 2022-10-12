package Verify

func VerifUsedLetter(letter string, ListLetterUsed []string) bool { // Fonction qui vérifie si la lettre renter à déjà etait rentrée
	for i := range ListLetterUsed {
		if letter == ListLetterUsed[i] {
			print("You cannot use this letter again! \n")
			return true
		}
	}
	return false
}
