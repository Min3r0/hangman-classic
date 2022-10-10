package Verify

func VerifWon(DashList []string) bool { // Fonction qui vérifie si le mot a était découvert en entier
	for i := range DashList {
		if DashList[i] == "_" {
			return false
		}
	}
	return true
}
