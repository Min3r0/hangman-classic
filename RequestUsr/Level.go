package RequestUsr

import (
	"bufio"
	"os"
)

func Level(save bool) string { //Fonction qui demande le niveau de difficulté
	scanner := bufio.NewScanner(os.Stdin) // Création du scanner capturant une entrée utilisateur
	scanner.Scan()                        // Lancement du scanner
	lv := scanner.Text()                  // Stockage du résultat du scanner dans une variable
	if save == true {
		print("You can continue (Continue) or Choose your level (Easy, Medium, Hard):\n")
		if lv == "Easy" {
			return "words.txt"
		} else if lv == "Medium" {
			return "words2.txt"
		} else if lv == "Hard" {
			return "words3.txt"
		} else if lv == "Continue" {
			return "Save.txt"
		} else {
			print("Your choice is not valid, try again!\n")
			return ""
		}
	} else {
		print("Choose your level (Easy, Medium, Hard):\n")
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
}
