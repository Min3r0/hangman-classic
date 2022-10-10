package CreateList

func AddLettreInDashList(lettre string, DashList []string, IndexLettre []int) []string { //fonction qui rajoute la lettre dans la liste de tiret
	for i := range IndexLettre {
		DashList[IndexLettre[i]] = lettre
	}
	return DashList
}
