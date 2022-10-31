package StartAndStop

import (
	"encoding/json"
	"fmt"
	"hangman-classic/Structure"
	"os"
)

func Stop(LineHangman int, ListLetterUsed []string, DashList []string, IndexOfDeath int, ListWordCap []string) {
	data := Structure.Hangman{
		LineHangman,
		ListLetterUsed,
		DashList,
		IndexOfDeath,
		ListWordCap}
	Data, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Party saved!")
	os.WriteFile("Save.json", Data, 0644)
	os.Exit(0)
}
