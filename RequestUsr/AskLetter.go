package RequestUsr

import (
	"bufio"
	"hangman-classic/CreateList"
	"hangman-classic/Print"
	"hangman-classic/Verify"
	"os"
)

func AskLetter(ListLetterUsed []string, ListASCII []string, CharList []string) string { // Fonction qui demande à l'utilisateur une lettre
	Scanner := bufio.NewScanner(os.Stdin) // création du Scanner capturant une entrée utilisateur
	print("Choose: ")
	Scanner.Scan()           //lancement du Scanner
	Letter := Scanner.Text() // stockage du résultat du Scanner dans une variable
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
