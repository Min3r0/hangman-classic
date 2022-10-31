package main

import (
	"hangman-classic/CreateList"
	"hangman-classic/Game"
	"hangman-classic/RequestUsr"
)

func main() {
	var ListLetterUsed []string
	CharList := []string{" ", "!", "\"", "#", "$", "%", "&", "'", "(", ")", "*", "+", ",", "-", ".", "/", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", ":", ";", "<", "=", ">", "?", "@", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "[", "\\", "]", "^", "_", "'", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "{", "|", "}", "~"}
	ASCII := RequestUsr.AskASCII()
	for ASCII == "" {
		ASCII = RequestUsr.AskASCII()
	}
	ListASCII := CreateList.ReadFile(ASCII)
	IndexOfDeath, LineHangman := 0, 0
	lev := RequestUsr.Level(ListASCII, CharList)
	for lev == "" {
		lev = RequestUsr.Level(ListASCII, CharList)
	}
	ListWord := CreateList.ReadFile(lev)
	ListWordCap := CreateList.CreateListWordCap(ListWord[RequestUsr.Random(ListWord)])
	DashList := CreateList.CreateDashList(ListWordCap)
	Game.Game(LineHangman, ListLetterUsed, DashList, IndexOfDeath, ListWordCap, ListASCII, CharList)
}
