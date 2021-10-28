package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	Read()
}

type HangManData struct {
	Word             string     // Word composed of '_', ex: Hll
	ToFind           string     // Final word chosen by the program at the beginning. It is the word to find
	Attempts         int        // Number of attempts left
	HangmanPositions [10]string // It can be the array where the positions parsed in "hangman.txt" are stored
	Correct          string
}

var s, entree string
var tab []string
var col []rune

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

	mot := tab[i]
	fmt.Println(mot)
	tabmot := []rune(mot)
	fmt.Println(tabmot)
	for u := 0; u < len(tabmot); u++ {
		fmt.Printf("")
		col = append(col, '_')

	}
	for {
		reader := bufio.NewReader(os.Stdin)
		if err != nil {
			println(err)
		}
		input, _ := reader.ReadString('\n')
		// fmt.Print(input)
		tabrun := []rune(input)
		for x := 0; x < len(tabmot); x++ {
			if tabmot[x] == tabrun[0] {
				col[x] = tabmot[x]
				fmt.Println(string(col))
			}
		}
	}
}
