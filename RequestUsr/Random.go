package RequestUsr

import (
	"math/rand"
	"time"
)

func Random(List []string) int { //Fonction qui retourne un nombre aléatoire entre 1 et la taille d'une liste
	rand.Seed(time.Now().UnixNano())
	Len := len(List)
	x := rand.Intn(Len)
	return x
}
