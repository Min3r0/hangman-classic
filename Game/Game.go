package Game

import (
	"fmt"
	"hangman-classic/CreateList"
	"hangman-classic/Print"
	"hangman-classic/RequestUsr"
	"hangman-classic/StartAndStop"
	"hangman-classic/Verify"
)

func Game(LineHangman int, ListLetterUsed []string, DashList []string, IndexOfDeath int, ListWordCap []string) {
	var StringLetterUsed string
	HangmanList := CreateList.ReadFile("hangman.txt")
	RandomLetter := CreateList.RandomLetter(ListWordCap)
	IndexRandomLetter := Verify.VerifyLetter(RandomLetter, ListWordCap)
	CreateList.AddLettreInDashList(RandomLetter, DashList, IndexRandomLetter)
	ListLetterUsed = append(ListLetterUsed, RandomLetter)
	Print.DrawHangman(LineHangman, HangmanList)
	Print.PrintDashList(DashList)
	print("\n")
	for IndexOfDeath < 10 {
		if Verify.VerifWon(DashList) == false {
			letter := RequestUsr.AskLetter(ListLetterUsed)
			ListLetterUsed = append(ListLetterUsed, letter)
			StillPlaying := true
			for letter == "0" {
				letter = RequestUsr.AskLetter(ListLetterUsed)
				ListLetterUsed = append(ListLetterUsed, letter)
			}
			if letter == "STOP" {
				StartAndStop.Stop(LineHangman, ListLetterUsed, DashList, IndexOfDeath, ListWordCap)
				StillPlaying = false
			}
			if !StillPlaying {
				break
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
				DashList = CreateList.AddLettreInDashList(letter, DashList, IndexLetter)
			}
			fmt.Println("il vous reste, ", 10-IndexOfDeath, " essai.")
			Print.DrawHangman(LineHangman, HangmanList)
			print("\n")
			fmt.Println("[", StringLetterUsed, "]")
			if Verify.VerifOneTap(letter, ListWordCap) {
				IndexOfDeath = 10
				print(letter)
				print("\n")
				print("GG, you're the best player i've ever seen, WOW!")
			} else if Verify.VerifWon(DashList) {
				DashList = CreateList.AddLettreInDashList(letter, DashList, IndexLetter)
				Print.PrintDashList(DashList)
				print("\n")
				print("GG, You Win")
				IndexOfDeath = 10
			} else {
				Print.PrintDashList(DashList)
				print("\n")
			}
		} else {
			print("You Loose")
			IndexOfDeath = 10
		}
	}
}
