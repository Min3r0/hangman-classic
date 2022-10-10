package Verify

func VerifUsedLetter(letter string, ListLetterUsed []string) int { // Fonction qui vérifie si la lettre renter à déjà etait rentrée
	for i := range ListLetterUsed {
		if letter == ListLetterUsed[i] {
			print("You cannot use this letter again!")
			return 1
		}
	}
	return 0
}
