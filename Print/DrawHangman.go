package Print

import "fmt"

func DrawHangman(IndexOfDeath int, HangmanList []string) {
	for i := IndexOfDeath; i < IndexOfDeath+7; i++ {
		fmt.Print(HangmanList[i])
		print("\n")
	}
}
