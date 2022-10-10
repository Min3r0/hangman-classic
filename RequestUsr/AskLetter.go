package RequestUsr

import (
	"bufio"
	"os"
)

func AskLetter() string { // Fonction qui demande à l'utilisateur une lettre
	scanner := bufio.NewScanner(os.Stdin) // création du scanner capturant une entrée utilisateur
	print("Choose: ")
	scanner.Scan()           //lancement du scanner
	letter := scanner.Text() // stockage du résultat du scanner dans une variable
	if letter >= "A" && letter <= "Z" {
		return letter
	} else {
		print("this character is not capital, this is not valid! \n")
		return "0"
	}
}
