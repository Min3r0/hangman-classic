package Print

func UsedLetter(letter string, LetterUsed string) string { //Fonction qui stock et rajoute les lettre ou mot utilisé
	LetterUsed += letter + ", "
	return LetterUsed
}
