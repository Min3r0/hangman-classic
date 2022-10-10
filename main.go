package main

import (
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
	Print.DrawHangman(LineHangman, HangmanList)
	Print.PrintDashList(DashList)
	print("\n")
	for IndexOfDeath < 9 {
		if Verify.VerifWon(DashList) == false {
			letter := RequestUsr.AskLetter()
			for Verify.VerifUsedLetter(letter, ListLetterUsed) == 1 || letter == "0" {
				ListLetterUsed = append(ListLetterUsed, letter)
				letter = RequestUsr.AskLetter()
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
			print("il vous reste, ", 9-IndexOfDeath, " essai.")
			print("\n")
			Print.DrawHangman(LineHangman, HangmanList)
			Print.PrintDashList(DashList)
			print("\n[")
			print(StringLetterUsed)
			print("]\n")
			if Verify.VerifOneTap(letter, ListWordCap) {
				IndexOfDeath = 10
				CreateList.AddLettreInDashList(letter, DashList, IndexLetter)
				print("GG, you're the best player i've ever seen, WOW!")
			}
		} else {
			IndexOfDeath = 10
		}
	}
	if Verify.VerifWon(DashList) {
		print("GG, You Win")
	} else {
		print("You Loose")
	}
}
