package Print

import "fmt"

func DrawHangman(IndexOfDeath int, HangmanList []string) { //Fonction qui affiche la bonne étape du hangman
	for i := IndexOfDeath; i < IndexOfDeath+7; i++ {
		fmt.Print(HangmanList[i])
		print("\n")
	}
}
