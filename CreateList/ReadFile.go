package CreateList

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(name string) []string { //Fonction qui met tous les mot du .txt choisi dans une liste de mot.
	var list []string
	f, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		list = append(list, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return list
}
