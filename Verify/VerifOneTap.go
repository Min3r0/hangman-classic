package Verify

func VerifOneTap(letter string, ListWordCap []string) bool {
	var boite string
	if len(letter) > 1 {
		for y := range ListWordCap {
			boite += ListWordCap[y]
			if letter == boite {
				return true
			}
		}
	}
	print("\n")
	return false
}
