package main

import (
	"flag"
	"hangman-classic/CreateList"
	"hangman-classic/Game"
	"hangman-classic/RequestUsr"
	"hangman-classic/StartAndStop"
)

func init() {
	const (
		Save  = "Save.txt"
		usage = "init du jeu"
	)
	flag.String("startWith", Save, usage)
}

func main() {
	if len(CreateList.ReadFile("save.txt")) >= 1 {
		Level := RequestUsr.Level(true)
		for Level == "" {
			Level = RequestUsr.Level(true)
		}
		if Level == "Continue" {

			Game.Game(StartAndStop.Start())
		} else {
			var ListLetterUsed []string
			IndexOfDeath, LineHangman := 0, 0
			ListWord := CreateList.ReadFile(Level)
			ListWordCap := CreateList.CreateListWordCap(ListWord[RequestUsr.Random(ListWord)])
			DashList := CreateList.CreateDashList(ListWordCap)
			Game.Game(LineHangman, ListLetterUsed, DashList, IndexOfDeath, ListWordCap)
		}
	} else {
		var ListLetterUsed []string
		IndexOfDeath, LineHangman := 0, 0
		Level := RequestUsr.Level(false)
		for Level == "" {
			Level = RequestUsr.Level(false)
		}
		ListWord := CreateList.ReadFile(Level)
		ListWordCap := CreateList.CreateListWordCap(ListWord[RequestUsr.Random(ListWord)])
		DashList := CreateList.CreateDashList(ListWordCap)
		Game.Game(LineHangman, ListLetterUsed, DashList, IndexOfDeath, ListWordCap)
	}
}
