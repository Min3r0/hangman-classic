package RequestUsr

import (
	"bufio"
	"hangman-classic/CreateList"
	"hangman-classic/Print"
	"hangman-classic/Verify"
	"os"
)

func AskLetter(ListLetterUsed []string, ListASCII []string, CharList []string) string { // Fonction qui demande à l'utilisateur une lettre
	scanner := bufio.NewScanner(os.Stdin) // création du scanner capturant une entrée utilisateur
	print("Choose: ")
	scanner.Scan()           //lancement du scanner
	Letter := scanner.Text() // stockage du résultat du scanner dans une variable
	if Letter < "A" || Letter > "Z" {
		message := "This character is not capital!"
		ListWordASCII := CreateList.CreateASCIIWordList(message, ListASCII, CharList)
		Print.PrintASCII(ListWordASCII)
		print("\n")
		return "0"
	} else if Verify.VerifUsedLetter(Letter, ListLetterUsed, ListASCII, CharList) {
		return "0"
	} else {
		return Letter
	}
}
