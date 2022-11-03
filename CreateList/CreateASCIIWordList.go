package CreateList

func CreateASCIIWordList(message string, ListASCII []string, CharList []string) []string { //Créer transforme une chaine de charactère en ASCII
	WordASCII := []string{"", "", "", "", "", "", ""}
	var Index []int
	ListWord := CreateListWord(message)
	for i := range ListWord {
		for j := range CharList {
			if ListWord[i] == CharList[j] {
				Index = append(Index, j)
			}
		}
	}
	for x := 0; x < 7; x++ {
		for y := 0; y < len(Index); y++ {
			WordASCII[x] += ListASCII[1+((9*Index[y])+x)]
		}
	}
	return WordASCII
}
