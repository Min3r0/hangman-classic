package Print

import "fmt"

func PrintASCII(ListWordASCII []string) {
	for i := range ListWordASCII {
		fmt.Println(ListWordASCII[i])
	}
}
