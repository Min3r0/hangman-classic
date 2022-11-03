package CreateList

func AddLettreInDashList(Lettre string, DashList []string, IndexLettre []int) []string { //fonction qui rajoute la Lettre dans la liste de tiret
	for i := range IndexLettre {
		DashList[IndexLettre[i]] = Lettre
	}
	return DashList
}
