package Verify

func VerifUsedLetter(letter string, ListLetterUsed []string) int {
	for i := range ListLetterUsed {
		if letter == ListLetterUsed[i] {
			print("You cannot use this letter again!")
			return 1
		}
	}
	return 0
}
