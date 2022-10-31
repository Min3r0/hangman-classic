package RequestUsr

import (
	"bufio"
	"hangman-classic/CreateList"
	"hangman-classic/Print"
	"os"
)

func Level(ListASCII []string, CharList []string) string { //Fonction qui demande le niveau de difficulté
	Scanner := bufio.NewScanner(os.Stdin) // Création du Scanner capturant une entrée utilisateur
	print("Choose your level (Easy, Medium, Hard):\n")
	Scanner.Scan()              // Lancement du Scanner
	LevelName := Scanner.Text() // Stockage du résultat du Scanner dans une variable
	if LevelName == "Easy" {
		return "words.txt"
	} else if LevelName == "Medium" {
		return "words2.txt"
	} else if LevelName == "Hard" {
		return "words3.txt"
	} else {
		Message := "Your choice is not valid, try again!"
		ListWordASCII := CreateList.CreateASCIIWordList(Message, ListASCII, CharList)
		Print.PrintASCII(ListWordASCII)
		print("\n")
		return ""
	}
}
