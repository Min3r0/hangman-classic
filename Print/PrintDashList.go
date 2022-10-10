package Print

import "fmt"

func PrintDashList(DashList []string) {
	for i := range DashList {
		fmt.Print(DashList[i])
	}
}
