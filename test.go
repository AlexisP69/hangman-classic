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
	Word             string     // Word composed of '_', ex: H_ll_
	ToFind           string     // Final word chosen by the program at the beginning. It is the word to find
	Attempts         int        // Number of attempts left
	HangmanPositions [10]string // It can be the array where the positions parsed in "hangman.txt" are stored
	Correct          string
}

var tab []string
var col []rune
var tab2 []string

func Read() {

	isdommage := true

	file, err := os.Open("words.txt")
	// file2, erreur := os.Open("hangman.txt")

	if err != nil { //fichier mot
		log.Fatal(err)
	}

	// if erreur != nil { //fichier position (hangman)
	// 	log.Fatal(erreur)
	// }

	defer file.Close() //fermer le fichier mot

	// defer file2.Close() //fermer le fichier position (hangman)

	scanner := bufio.NewScanner(file) //Lis mon fichier words
	for scanner.Scan() {
		tab = append(tab, scanner.Text())
		if err := scanner.Err(); err != nil {
			log.Fatalln(err)
		}
	}

	// position := bufio.NewScanner(file2) //Lis mon fichier position (hangman)
	// for position.Scan() {
	// 	tab2 = append(tab2, position.Text())
	// 	if erreur := position.Err(); erreur != nil {
	// 		log.Fatalln(erreur)
	// 	}
	// }

	rand.Seed(time.Now().UnixNano())
	min := 0
	max := len(tab)
	i := rand.Intn((max - min))

	mot := tab[i] // Mot aléatoire du fichier words
	fmt.Println(mot)
	tabmot := []rune(mot) // fmt.Println(tabmot) Affiche les rune des underscore
	for u := 0; u < len(tabmot); u++ {
		fmt.Printf("_")
		col = append(col, '_')
	}

	fmt.Print("\n")
	fmt.Print("Choissisez votre lettre: ")

	for {
		reader := bufio.NewReader(os.Stdin)
		if err != nil {
			println(err)
		}
		input, _ := reader.ReadString('\n')
		tabrun := []rune(input) //la rune de notre input

		isdommage = true

		for x := 0; x < len(tabmot); x++ {
			if tabmot[x] == tabrun[0] { // compare l'index du mot a l'index de mon input
				col[x] = tabmot[x]
				isdommage = false
			}
		}
		fmt.Println(string(col)) //Affiche le mot avec les underscore modifier
		fmt.Print("Choissisez votre lettre: ")

		if isdommage == true {
			fmt.Println("DOMMAGE CETTE LETTRE N'EST PAS DANS CE MOTS")
			fmt.Print("\n")
			// fmt.Println(tab2)
			fmt.Println(string(col))
			fmt.Print("Choissisez votre lettre: ")
		}
		if mot == string(col) {
			fmt.Print("\n")
			fmt.Println("Congratulation You found the word")
		}
	}
}
