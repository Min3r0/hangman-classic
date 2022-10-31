package RequestUsr

import (
	"bufio"
	"os"
)

func AskASCII() string {
	scanner := bufio.NewScanner(os.Stdin)
	print("choose a model of ASCII (Standard, Shadow, Thinkertoy):\n")
	scanner.Scan()          // Lancement du scanner
	ASCII := scanner.Text() // Stockage du résultat du scanner dans une variable
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
