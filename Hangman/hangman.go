package Hangman

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

var tab []string

func Read() {

	file, err := os.Open("words.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tab = append(tab, scanner.Text())
		if err := scanner.Err(); err != nil {
			log.Fatalln(err)
		}
	}
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := len(tab)
	i := rand.Intn((max - min))
	fmt.Println(tab[i])
}
