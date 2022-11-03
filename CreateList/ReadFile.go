package CreateList

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(name string) []string { //Fonction qui met tous les mot du .txt choisi dans une liste de mot.
	var List []string
	f, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		List = append(List, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return List
}
