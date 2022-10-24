package StartAndStop

import (
	"encoding/json"
	"hangman-classic/Structure"
	"io/ioutil"
)

func Start() (int, []string, []string, int, []string) {
	var save Structure.Hangman
	data, _ := ioutil.ReadFile("save.txt")
	json.Unmarshal(data, save)
	return save.LineHangman, save.ListLetterUsed, save.DashList, save.IndexOfDeath, save.ListWordCap
}
