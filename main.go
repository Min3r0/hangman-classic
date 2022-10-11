package main

import (
	"fmt"
	"hangman-classic/CreateList"
	"hangman-classic/Print"
	"hangman-classic/RequestUsr"
	"hangman-classic/Verify"
)

func main() {
	var StringLetterUsed string
	var ListLetterUsed []string
	IndexOfDeath, LineHangman := 0, 0
	lev := RequestUsr.Level()
	for lev == "" {
		lev = RequestUsr.Level()
	}
	HangmanList := CreateList.ReadFile("hangman.txt")
	ListWord := CreateList.ReadFile(lev)
	ListWordCap := CreateList.CreateListWordCap(ListWord[RequestUsr.Random(ListWord)])
	DashList := CreateList.CreateDashList(ListWordCap)
	RandomLetter := CreateList.RandomLetter(ListWordCap)
	IndexRandomLetter := Verify.VerifyLetter(RandomLetter, ListWordCap)
	CreateList.AddLettreInDashList(RandomLetter, DashList, IndexRandomLetter)
	ListLetterUsed = append(ListLetterUsed, RandomLetter)
	Print.DrawHangman(LineHangman, HangmanList)
	Print.PrintDashList(DashList)
	print("\n")
	for IndexOfDeath < 9 {
		if Verify.VerifWon(DashList) == false {
			letter := RequestUsr.AskLetter(ListLetterUsed)
			ListLetterUsed = append(ListLetterUsed, letter)
			for letter == "0" {
				letter = RequestUsr.AskLetter(ListLetterUsed)
				ListLetterUsed = append(ListLetterUsed, letter)
			}
			StringLetterUsed = Print.UsedLetter(letter, StringLetterUsed)
			IndexLetter := Verify.VerifyLetter(letter, ListWordCap)
			if len(IndexLetter) == 0 && len(letter) == 1 {
				LineHangman = LineHangman + 8
				IndexOfDeath++
			} else if len(IndexLetter) == 0 && len(letter) >= 2 {
				LineHangman = LineHangman + 16
				IndexOfDeath += 2
			} else {
				CreateList.AddLettreInDashList(letter, DashList, IndexLetter)
			}
			fmt.Println("il vous reste, ", 9-IndexOfDeath, " essai.")
			Print.DrawHangman(LineHangman, HangmanList)
			print("\n")
			fmt.Println("[", StringLetterUsed, "]")
			if Verify.VerifOneTap(letter, ListWordCap) {
				IndexOfDeath = 10
				print(letter)
				print("\n")
				print("GG, you're the best player i've ever seen, WOW!")
			} else if Verify.VerifWon(DashList) {
				CreateList.AddLettreInDashList(letter, DashList, IndexLetter)
				Print.PrintDashList(DashList)
				print("\n")
				print("GG, You Win")
				IndexOfDeath = 10
			} else {
				Print.PrintDashList(DashList)
				print("\n")
			}
		} else {
			IndexOfDeath = 10
			print("You Loose")
		}
	}
}
