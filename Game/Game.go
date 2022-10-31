package Game

import (
	"fmt"
	"hangman-classic/CreateList"
	"hangman-classic/Print"
	"hangman-classic/RequestUsr"
	"hangman-classic/StartAndStop"
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
	for IndexOfDeath < 10 { //Tant que l'indicateur de mort n'a pas atteint 10
		if Verify.VerifWon(DashList) == false { //Vérification de la condition de victoire
			Letter := RequestUsr.AskLetter(ListLetterUsed, ListASCII, CharList)
			ListLetterUsed = append(ListLetterUsed, Letter)
			StillPlaying := true
			for Letter == "0" {
				Letter = RequestUsr.AskLetter(ListLetterUsed, ListASCII, CharList)
				ListLetterUsed = append(ListLetterUsed, Letter)
			}
			if Letter == "STOP" {
				StartAndStop.Stop(LineHangman, ListLetterUsed, DashList, IndexOfDeath, ListWordCap)
				StillPlaying = false
			}
			if !StillPlaying {
				break
			}
			StringLetterUsed = Print.UsedLetter(Letter, StringLetterUsed)
			IndexLetter := Verify.VerifyLetter(Letter, ListWordCap)
			if len(IndexLetter) == 0 && len(Letter) == 1 {
				LineHangman = LineHangman + 8
				IndexOfDeath++
			} else if len(IndexLetter) == 0 && len(Letter) >= 2 {
				LineHangman = LineHangman + 16
				IndexOfDeath += 2
			} else {
				CreateList.AddLettreInDashList(Letter, DashList, IndexLetter)
			}
			if IndexOfDeath == 11 { // si il reste 1 tentative et le joueur essais un mot
				IndexOfDeath = 10
				LineHangman -= 8
			}
			NbDeathInString := strconv.FormatInt(int64(10-IndexOfDeath), 10)
			message := "You have " + NbDeathInString + " attempts left"
			ListWordASCII := CreateList.CreateASCIIWordList(message, ListASCII, CharList)
			Print.PrintASCII(ListWordASCII)
			print("\n")
			Print.DrawHangman(LineHangman, HangmanList)
			print("\n")
			fmt.Println("[", StringLetterUsed, "]")
			if Verify.VerifOneTap(Letter, ListWordCap) {
				IndexOfDeath = 12
				println(Letter)
				CreateList.AddLettreInDashList(Letter, DashList, IndexLetter)
				message := "GG, you're the best player !"
				ListWordASCII := CreateList.CreateASCIIWordList(message, ListASCII, CharList)
				Print.PrintASCII(ListWordASCII)
			} else if Verify.VerifWon(DashList) {
				IndexOfDeath = 12
				CreateList.AddLettreInDashList(Letter, DashList, IndexLetter)
				Print.PrintDashList(DashList)
				print("\n")
				message := "GG, you win"
				ListWordASCII := CreateList.CreateASCIIWordList(message, ListASCII, CharList)
				Print.PrintASCII(ListWordASCII)
			} else {
				Print.PrintDashList(DashList)
				print("\n")
			}
		}
	}
	if IndexOfDeath >= 9 && IndexOfDeath < 12 {
		Message := "You Loose "
		ListWordASCII := CreateList.CreateASCIIWordList(Message, ListASCII, CharList)
		Print.PrintASCII(ListWordASCII)

	}
}
