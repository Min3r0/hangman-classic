package StartAndStop

import (
	"encoding/json"
	"fmt"
	"hangman-classic/Structure"
	"io/ioutil"
	"os"
)

func Stop(LineHangman int, ListLetterUsed []string, DashList []string, IndexOfDeath int) {
	data := Structure.Hangman{
		LineHangman:    LineHangman,
		ListLetterUsed: ListLetterUsed,
		DashList:       DashList,
		IndexOfDeath:   IndexOfDeath}
	Data, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Party saved!")
	ioutil.WriteFile("Save.txt", Data, 0644)
	os.Exit(0)
}
