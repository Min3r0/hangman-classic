package RequestUsr

import (
	"bufio"
	"os"
)

func AskASCII() string {
	Scanner := bufio.NewScanner(os.Stdin)
	print("choose a model of ASCII (Standard, Shadow, Thinkertoy):\n")
	Scanner.Scan()          // Lancement du Scanner
	ASCII := Scanner.Text() // Stockage du résultat du Scanner dans une variable
	if ASCII == "Standard" {
		return "standard.txt"
	} else if ASCII == "Shadow" {
		return "shadow.txt"
	} else if ASCII == "Thinkertoy" {
		return "thinkertoy.txt"
	} else {
		print("Your choice is not valid, try again!\n")
		return ""
	}
}
