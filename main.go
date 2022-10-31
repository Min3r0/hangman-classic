package main

import (
	"hangman-classic/CreateList"
	"hangman-classic/Game"
	"hangman-classic/RequestUsr"
)

func main() {

	var ListLetterUsed []string
	IndexOfDeath, LineHangman := 0, 0
	Level := RequestUsr.Level()
	for Level == "" {
		Level = RequestUsr.Level()
	}
	ListWord := CreateList.ReadFile(Level)
	ListWordCap := CreateList.CreateListWordCap(ListWord[RequestUsr.Random(ListWord)])
	DashList := CreateList.CreateDashList(ListWordCap)
	Game.Game(LineHangman, ListLetterUsed, DashList, IndexOfDeath, ListWordCap)
}
