package RequestUsr

import (
	"math/rand"
	"time"
)

func Random(list []string) int { //Fonction qui retourne un nombre aléatoire entre 1 et la taille d'une liste
	rand.Seed(time.Now().UnixNano())
	Len := len(list)
	x := rand.Intn(Len)
	return x
}
