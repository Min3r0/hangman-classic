package CreateList

import "math/rand"

func RandomLetter(ListWordCap []string) string {
	LenListWordCap := len(ListWordCap)
	i := rand.Intn(LenListWordCap)
	return ListWordCap[i]
}
