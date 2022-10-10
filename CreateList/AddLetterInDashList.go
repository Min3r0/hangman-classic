package CreateList

func AddLettreInDashList(lettre string, DashList []string, IndexLettre []int) []string {
	for i := range IndexLettre {
		DashList[IndexLettre[i]] = lettre
	}
	return DashList
}
