package main

import (
	"flag"
	"hangman-classic/CreateList"
	"hangman-classic/RequestUsr"
)

func init() {
	const (
		Save  = "Save.txt"
		usage = "init du jeu"
	)
	flag.String("SaveType", Save, usage)
}

func main() {
	var ListLetterUsed []string
	IndexOfDeath, LineHangman := 0, 0
	lev := RequestUsr.Level()
	for lev == "" {
		lev = RequestUsr.Level()
	}
	ListWord := CreateList.ReadFile(lev)
	ListWordCap := CreateList.CreateListWordCap(ListWord[RequestUsr.Random(ListWord)])
	DashList := CreateList.CreateDashList(ListWordCap)
	Game(LineHangman, ListLetterUsed, DashList, IndexOfDeath, ListWordCap)
}
