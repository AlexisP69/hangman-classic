package Hangman

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func Read() {

	var (
		tab      []string
		col      []rune
		attempts int
		essai    int
	)

	essai = 9
	isfalse := true

	file, err := os.Open("words.txt")

	if err != nil { //fichier mot
		log.Fatal(err)
	}

	defer file.Close() //fermer le fichier mot

	scanner := bufio.NewScanner(file) //Lis mon fichier words
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

		isfalse = true

		for x := 0; x < len(tabmot); x++ {
			if tabmot[x] == tabrun[0] { // compare l'index du mot a l'index de mon input
				col[x] = tabmot[x]
				isfalse = false
			}
		}
		fmt.Println(string(col)) //Affiche le mot avec les underscore modifier
		fmt.Print("Choissisez votre lettre: ")

		if isfalse == true {
			fmt.Println("Dommage cette lettre n'est pas dans ce mot")
			fmt.Print("\n")
			fmt.Println("Il vous reste", essai, "essai")
			essai = essai - 1
			Draw(attempts)
			attempts = attempts + 1
			fmt.Println(string(col))
			fmt.Print("Choissisez votre lettre: ")
		}
		if essai == -1 {
			fmt.Println("\n")
			fmt.Println("You're dead")
		}
		if mot == string(col) {
			fmt.Print("\n")
			fmt.Println("Congratulation You found the word")
		}
	}
}

func Draw(attempts int) {
	count := 0

	file2, erreur := os.Open("hangman.txt")

	if erreur != nil { //fichier mot
		log.Fatal(erreur)
	}

	defer file2.Close() //fermer le fichier mot

	position := bufio.NewScanner(file2) //Lis mon fichier words
	for position.Scan() {
		if count >= attempts*8 && count < (attempts*8)+8 {
			fmt.Println(position.Text())
		}
		if erreur := position.Err(); erreur != nil {
			log.Fatalln(erreur)
		}
		count++
	}
}
