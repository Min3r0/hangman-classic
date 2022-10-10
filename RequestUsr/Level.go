package RequestUsr

import (
	"bufio"
	"os"
)

func Level() string { //Fonction qui demande le niveau de difficulté
	scanner := bufio.NewScanner(os.Stdin) // Création du scanner capturant une entrée utilisateur
	print("Choose your level (Easy, Medium, Hard):\n")
	scanner.Scan()       // Lancement du scanner
	lv := scanner.Text() // Stockage du résultat du scanner dans une variable
	if lv == "Easy" {
		return "words.txt"
	} else if lv == "Medium" {
		return "words2.txt"
	} else if lv == "Hard" {
		return "words3.txt"
	} else {
		print("Your choice is not valid, try again!\n")
		return ""
	}
}
