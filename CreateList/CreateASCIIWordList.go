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
	for i := 0; i > 9; i++ {
		for j := range index {
			WordASCII[i] += ListASCII[2+((9*index[j])+i)]
			WordASCII[i] += "/"
		}
	}
	return WordASCII
}
