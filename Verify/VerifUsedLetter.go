package Verify

import (
	"hangman-classic/CreateList"
	"hangman-classic/Print"
)

func VerifUsedLetter(letter string, ListLetterUsed []string, ListASCII []string, CharList []string) bool { // Fonction qui vérifie si la lettre renter à déjà etait rentrée
	for i := range ListLetterUsed {
		if letter == ListLetterUsed[i] {
			message := "You cannot use this letter again!"
			ListWordASCII := CreateList.CreateASCIIWordList(message, ListASCII, CharList)
			Print.PrintASCII(ListWordASCII)
			print("\n")
			return true
		}
	}
	return false
}
