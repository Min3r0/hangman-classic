package Verify

import (
	"hangman-classic/CreateList"
	"hangman-classic/Print"
)

func VerifUsedLetter(Letter string, ListLetterUsed []string, ListASCII []string, CharList []string) bool { // Fonction qui vérifie si la lettre renter à déjà etait rentrée
	for i := range ListLetterUsed {
		if Letter == ListLetterUsed[i] {
			Message := "You cannot use this letter again!"
			ListWordASCII := CreateList.CreateASCIIWordList(Message, ListASCII, CharList)
			Print.PrintASCII(ListWordASCII)
			print("\n")
			return true
		}
	}
	return false
}
