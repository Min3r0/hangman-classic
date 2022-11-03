package CreateList

import "math/rand"

func RandomLetter(ListWordCap []string) string {
	i := rand.Intn(len(ListWordCap))
	return ListWordCap[i]
}
