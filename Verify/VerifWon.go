package Verify

func VerifWon(DashList []string) bool {
	for i := range DashList {
		if DashList[i] == "_" {
			return false
		}
	}
	return true
}
