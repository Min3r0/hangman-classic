package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func level() string { //Fonction qui demande le niveau de difficulté
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

func DrawHangman(IndexOfDeath int, HangmanList []string) {
	for i := IndexOfDeath; i < IndexOfDeath+7; i++ {
		fmt.Print(HangmanList[i])
		print("\n")
	}
}

func readfile(name string) []string { //Fonction qui met tous les mot du .txt choisi dans une liste de mot.
	var list []string
	f, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		list = append(list, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return list
}

func Random(list []string) int { //Fonction qui retourne un nombre aléatoire entre 1 et la taille d'une liste
	rand.Seed(time.Now().UnixNano())
	Len := len(list)
	x := rand.Intn(Len)
	return x
}

func CreateListWordCap(Word string) []string { //Fonction qui retourne une liste de lettre en majuscule
	var ListLettre []string
	for i := 0; i < len(Word); i++ {
		ListLettre = append(ListLettre, string(Word[i]-32))
	}
	return ListLettre
}

func CreateDashList(list []string) []string { // Fonction qui créer une liste de tiret aussi longue que le nombre de lettre du mot choisi
	var DashListLetter []string
	for dash := 0; dash < len(list); dash++ {
		DashListLetter = append(DashListLetter, "_")
	}
	return DashListLetter
}

func AskLetter() string {
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

func VerifyLetter(letter string, ListWordCap []string) []int { // Fonction vérifie si la lettre choisi et bien dans le mot et renvois une liste d'entier composer de la position de la lettre.
	var list []int
	for i := range ListWordCap {
		if letter == ListWordCap[i] {
			list = append(list, i)
		}
	}
	return list
}

func AddLettreInDashList(lettre string, DashList []string, IndexLettre []int) []string {
	for i := range IndexLettre {
		DashList[IndexLettre[i]] = lettre
	}
	return DashList
}

func PrintDashList(DashList []string) {
	for i := range DashList {
		fmt.Print(DashList[i])
	}
}

func VerifWon(DashList []string) bool {
	for i := range DashList {
		if DashList[i] == "_" {
			return false
		}
	}
	return true
}

func UsedLetter(letter string, LetterUsed string) string {
	LetterUsed += letter + ", "
	return LetterUsed
}

func verifonetap(letter string, ListWordCap []string) bool {
	var boite string
	if len(letter) > 1 {
		for y := range ListWordCap {
			boite += ListWordCap[y]
			if letter == boite {
				return true
			}
		}
	}
	return false
}

func VerifUsedLetter(letter string, ListLetterUsed []string) int {
	for i := range ListLetterUsed {
		if letter == ListLetterUsed[i] {
			print("You cannot use this letter again!")
			return 1
		}
	}
	return 0
}

func main() {
	var StringLetterUsed string
	var ListLetterUsed []string
	IndexOfDeath, LineHangman := 0, 0
	lev := level()
	for lev == "" {
		lev = level()
	}
	HangmanList := readfile("hangman.txt")
	ListWord := readfile(lev)
	ListWordCap := CreateListWordCap(ListWord[Random(ListWord)])
	DashList := CreateDashList(ListWordCap)
	DrawHangman(LineHangman, HangmanList)
	PrintDashList(DashList)
	print("\n")
	for IndexOfDeath < 9 {
		if VerifWon(DashList) == false {
			letter := AskLetter()
			for VerifUsedLetter(letter, ListLetterUsed) == 1 || letter == "0" {
				ListLetterUsed = append(ListLetterUsed, letter)
				letter = AskLetter()
			}
			StringLetterUsed = UsedLetter(letter, StringLetterUsed)
			IndexLetter := VerifyLetter(letter, ListWordCap)
			if len(IndexLetter) == 0 && len(letter) == 1 {
				LineHangman = LineHangman + 8
				IndexOfDeath++
			} else if len(IndexLetter) == 0 && len(letter) >= 2 {
				LineHangman = LineHangman + 16
				IndexOfDeath += 2
			} else {
				AddLettreInDashList(letter, DashList, IndexLetter)
			}
			print("il vous reste, ", 9-IndexOfDeath, " essai.")
			print("\n")
			DrawHangman(LineHangman, HangmanList)
			PrintDashList(DashList)
			print("\n[")
			print(StringLetterUsed)
			print("]\n")
			if verifonetap(letter, ListWordCap) {
				IndexOfDeath = 10
				AddLettreInDashList(letter, DashList, IndexLetter)
				print("GG, you're the best player i've ever seen, WOW!")
			}
		} else {
			IndexOfDeath = 10
		}
	}
	if VerifWon(DashList) {
		print("GG, You Win")
	} else {
		print("You Loose")
	}
}
