package CreateList

func CreateASCIIWordList(message string, ListASCII []string, CharList []string) []string {
	WordASCII := []string{"", "", "", "", "", "", ""}
	var index []int
	ListWord := CreateListWord(message)
	for i := range ListWord {
		for j := range CharList {
			if ListWord[i] == CharList[j] {
				index = append(index, j)
			}
		}
	}
	for x := 0; x < 7; x++ {
		for y := 0; y < len(index); y++ {
			WordASCII[x] += ListASCII[1+((9*index[y])+x)]
		}
	}
	return WordASCII
}
