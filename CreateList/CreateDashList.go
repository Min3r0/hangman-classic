package CreateList

func CreateDashList(list []string) []string { // Fonction qui créer une liste de tiret aussi longue que le nombre de lettre du mot choisi
	var DashListLetter []string
	for Dash := 0; Dash < len(list); Dash++ {
		DashListLetter = append(DashListLetter, "_")
	}
	return DashListLetter
}
