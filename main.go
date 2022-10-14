package main

import (
	"fmt"
	"hangman-classic/CreateList"
	"hangman-classic/Print"
	"hangman-classic/RequestUsr"
	"hangman-classic/Verify"
	"strconv"
)

func main() {
	var StringLetterUsed string
	var ListLetterUsed []string
	CharList := []string{" ", "!", "\"", "#", "$", "%", "&", "'", "(", ")", "*", "+", ",", "-", ".", "/", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", ":", ";", "<", "=", ">", "?", "@", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "[", "\\", "]", "^", "_", "'", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "{", "|", "}", "~"}
	ListASCII := CreateList.ReadFile("standard (1).txt")
	IndexOfDeath, LineHangman := 0, 0
	lev := RequestUsr.Level(ListASCII, CharList)
	for lev == "" {
		lev = RequestUsr.Level(ListASCII, CharList)
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
				message := "GG, you're the best player I've ever seen!"
				ListWordASCII := CreateList.CreateASCIIWordList(message, ListASCII, CharList)
				Print.PrintASCII(ListWordASCII)
				//print("GG, you're the best player i've ever seen, WOW!")
			} else if Verify.VerifWon(DashList) {
				CreateList.AddLettreInDashList(letter, DashList, IndexLetter)
				Print.PrintDashList(DashList)
				print("\n")
				message := "GG, you win"
				ListWordASCII := CreateList.CreateASCIIWordList(message, ListASCII, CharList)
				Print.PrintASCII(ListWordASCII)
				//print("GG, You Win")
				IndexOfDeath = 10
			} else {
				Print.PrintDashList(DashList)
				print("\n")
			}
		} else {
			IndexOfDeath = 10
			message := "You Loose"
			ListWordASCII := CreateList.CreateASCIIWordList(message, ListASCII, CharList)
			Print.PrintASCII(ListWordASCII)
			//print("You Loose")
		}
	}
}
