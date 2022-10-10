package Print

import "fmt"

func PrintDashList(DashList []string) { //Fonction qui affiche la liste de tiret donc l'évolution du mot à chercher
	for i := range DashList {
		fmt.Print(DashList[i])
	}
}
