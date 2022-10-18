package RequestUsr

import (
	"bufio"
	"hangman-classic/Verify"
	"os"
)

func AskLetter(ListLetterUsed []string) string { // Fonction qui demande à l'utilisateur une lettre
	scanner := bufio.NewScanner(os.Stdin) // création du scanner capturant une entrée utilisateur
	print("Choose: ")
	scanner.Scan()           //lancement du scanner
	Letter := scanner.Text() // stockage du résultat du scanner dans une variable
	if Letter < "A" || Letter > "Z" {
		print("this character  is not capital, this is not valid! \n")
		return "0"
	} else if Verify.VerifUsedLetter(Letter, ListLetterUsed) {
		return "0"
	} else {
		return Letter
	}
}
