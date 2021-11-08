package Hangman

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func Pendu() {

	var (
		tab        []string
		underscore []rune
		attempts   int
		essai      int
	)

	essai = 9
	isfalse := true
	// isfalse2 := true

	file, err := os.Open("words.txt") //ouvre le fichier de mots

	if err != nil { //fichier avec les mots
		log.Fatal(err)
	}

	defer file.Close()                //fermer la lecture du fichier des mots
	scanner := bufio.NewScanner(file) //Scan tout les mots dans le fichier mots
	for scanner.Scan() {
		tab = append(tab, scanner.Text())
		if err := scanner.Err(); err != nil { //retourne une erreur s'il n'arrive pas à scan les mots dans le fichier des mots
			log.Fatalln(err)
		}
	}

	rand.Seed(time.Now().UnixNano()) //choisi de façon aléatoire 1 mot
	min := 0
	max := len(tab)
	i := rand.Intn((max - min))

	mot := tab[i] // Mot aléatoire du fichier mots dans un tableau
	fmt.Println(mot)
	tabmot := []rune(mot) //initialise tabmot comme tableau de rune du mot
	for u := 0; u < len(tabmot); u++ {
		underscore = append(underscore, '_') //remplace le mot par des underscores
	}
	for g := 0; g < (len(mot)/2 - 1); g++ { //affichage la moitié -1 des lettres du mot
		A := rand.Intn(len(mot)) //choisi aléatoirement les lettres à afficher du mot
		underscore[A] = tabmot[A]
	}
	fmt.Print(string(underscore)) //print les underscores avec les lettres aléatoires
	fmt.Print("\n")
	fmt.Print("Choissisez votre lettre: ") //Demande à la personne de Choisir une lettre

	for {
		reader := bufio.NewReader(os.Stdin)
		if err != nil {
			println(err)
		}
		input, _ := reader.ReadString('\n') //Lit ce que l'on écrit
		tabrun := []rune(input)             //tableau de rune de ce que l'on écrit
		isfalse = true
		// isfalse2 = true
		// tab2 := string(tabrun)
		// if tab2 == mot {
		// 	fmt.Print("Bravo")
		// }
		for x := 0; x < len(tabmot); x++ {
			// fmt.Println(mot, tabrun[:len(tabrun)-1])
			if mot == string(tabrun[:len(tabrun)-1]) {
				underscore[x] = tabmot[x]
				// isfalse2 = false
			} else if tabmot[x] == tabrun[0] { //compare l'index du mot a l'index de mon input
				underscore[x] = tabmot[x]
				isfalse = false //permet de ne pas rentrer dans la condition plus bas qui est true
			}
		}
		if essai == 0 { //condition pour finir le jeu si perdu
			Draw(attempts)
			fmt.Println("You're dead")
			break
		}
		fmt.Println(string(underscore)) //Affiche le mot avec les underscores modifié

		if mot == string(underscore) { //condition pour finir le jeu si gagné
			fmt.Println("Congratulation You found the word")
			break
		}
		if isfalse == true { //si le mot ou l'input entrée est fausse il rentre dans la condition
			fmt.Println("Dommage cette lettre n'est pas dans ce mot")
			fmt.Print("\n")
			fmt.Println("Il vous reste", essai, "essai")
			essai = essai - 1
			Draw(attempts)
			attempts = attempts + 1
			fmt.Println(string(underscore))
			// } else if isfalse2 == true { //si le mot ou l'input entrée est fausse il rentre dans la condition
			// 	fmt.Println("Dommage cette lettre n'est pas dans ce mot")
			// 	fmt.Print("\n")
			// 	fmt.Println("Il vous reste", essai, "essai")
			// 	essai = essai - 2
			// 	Draw(attempts)
			// 	attempts = attempts + 2
			// 	fmt.Println(string(underscore))
			// }
		}
		fmt.Print("Choisissez votre lettre: ")
	}
}

func Draw(attempts int) {
	count := 0

	file2, erreur := os.Open("hangman.txt") //ouvre le fichier des positions du pendu

	if erreur != nil {
		log.Fatal(erreur)
	}

	defer file2.Close()                 //fermer le fichier mot
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
