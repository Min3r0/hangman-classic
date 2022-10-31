package Game

import (
	"fmt"
	"hangman-classic/CreateList"
	"hangman-classic/Print"
	"hangman-classic/RequestUsr"
	"hangman-classic/Verify"
	"strconv"
)

func Game(LineHangman int, ListLetterUsed []string, DashList []string, IndexOfDeath int, ListWordCap []string, ListASCII []string, CharList []string) {
	var StringLetterUsed string
	HangmanList := CreateList.ReadFile("hangman.txt")
	RandomLetter := CreateList.RandomLetter(ListWordCap)
	IndexRandomLetter := Verify.VerifyLetter(RandomLetter, ListWordCap)
	CreateList.AddLettreInDashList(RandomLetter, DashList, IndexRandomLetter)
	ListLetterUsed = append(ListLetterUsed, RandomLetter)
	Print.DrawHangman(LineHangman, HangmanList)
	Print.PrintDashList(DashList)
	print("\n")
	for IndexOfDeath < 9 {
		if Verify.VerifWon(DashList) == false {
			letter := RequestUsr.AskLetter(ListLetterUsed, ListASCII, CharList)
			ListLetterUsed = append(ListLetterUsed, letter)
			for letter == "0" {
				letter = RequestUsr.AskLetter(ListLetterUsed, ListASCII, CharList)
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
			NbDeathInString := strconv.FormatInt(int64(9-IndexOfDeath), 10)
			message := "you have " + NbDeathInString + " attempts left"
			ListWordASCII := CreateList.CreateASCIIWordList(message, ListASCII, CharList)
			Print.PrintASCII(ListWordASCII)
			//print("il vous reste, ", 9-IndexOfDeath, " essai.")
			print("\n")
			Print.DrawHangman(LineHangman, HangmanList)
			print("\n")
			fmt.Println("[", StringLetterUsed, "]")
			if Verify.VerifOneTap(letter, ListWordCap) {
				IndexOfDeath = 10
				print(letter)
				print("\n")
				CreateList.AddLettreInDashList(letter, DashList, IndexLetter)
				message := "GG, you're the best player !"
				ListWordASCII := CreateList.CreateASCIIWordList(message, ListASCII, CharList)
				Print.PrintASCII(ListWordASCII)
			} else if Verify.VerifWon(DashList) {
				CreateList.AddLettreInDashList(letter, DashList, IndexLetter)
				Print.PrintDashList(DashList)
				print("\n")
				message := "GG, you win"
				ListWordASCII := CreateList.CreateASCIIWordList(message, ListASCII, CharList)
				Print.PrintASCII(ListWordASCII)
				IndexOfDeath = 10
			} else {
				Print.PrintDashList(DashList)
				print("\n")
			}
		} else {
			IndexOfDeath = 10
			message := "You Loose "
			ListWordASCII := CreateList.CreateASCIIWordList(message, ListASCII, CharList)
			Print.PrintASCII(ListWordASCII)
			//print("You Loose")
		}
	}
}
