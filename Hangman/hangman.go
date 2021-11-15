package Hangman

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

type Game struct {
	tab        []string
	underscore []rune
	attempts   int
	essai      int
	p          []rune
	win        int
	loose      int
	isfalse    bool
}

func Pendu() {
	var game Game

	// var (
	// 	tab        []string
	// 	underscore []rune
	// 	attempts   int
	// 	essai      int
	// 	p          []rune
	// 	win        int
	// 	loose      int
	// )

	game.essai = 9
	game.isfalse = true

	file, err := os.Open("words.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		game.tab = append(game.tab, scanner.Text())
		if err := scanner.Err(); err != nil {
			log.Fatalln(err)
		}
	}

	rand.Seed(time.Now().UnixNano())
	min := 0
	max := len(game.tab)
	i := rand.Intn((max - min))

	mot := game.tab[i] // Mot aléatoire du fichier mots dans un tableau
	fmt.Println(mot)
	tabmot := []rune(mot)

	for u := 0; u < len(tabmot); u++ {
		game.underscore = append(game.underscore, '_')
	}

	for g := 0; g < (len(mot)/2 - 1); g++ { //affichage la moitié -1 des lettres du mot
		A := rand.Intn(len(mot)) //choisi aléatoirement les lettres à afficher du mot
		game.underscore[A] = tabmot[A]
	}
	fmt.Println("Bonne chance vous avez 10 essais")
	fmt.Println(string(game.underscore))
	fmt.Print("Choisissez votre lettre: ")

	for {
		reader := bufio.NewReader(os.Stdin)
		if err != nil {
			println(err)
		}

		input, _ := reader.ReadString('\n') //Lit ce que l'on écrit
		tabrun := []rune(input)             //tableau de rune de ce que l'on écrit
		game.isfalse = true

		for x := 0; x < len(tabmot); x++ {
			// fmt.Println(mot, tabrun[:len(tabrun)-1])
			if mot == string(tabrun[:len(tabrun)-1]) { //condition qui vérifie si le mot correspond à ce que l'on marque moins le \n
				game.underscore[x] = tabmot[x]
				game.isfalse = false
			} else if tabmot[x] == tabrun[0] { //compare l'index du mot a l'index de mon input
				game.underscore[x] = tabmot[x]
				game.isfalse = false
			}
		}

		if game.essai <= 0 { //condition pour finir le jeu si perdu
			Draw(game.attempts)
			fmt.Println("You're dead")
			game.loose++
			fmt.Println("Vous avez perdu", game.loose, "fois")
			Restart()
			break
		}

		fmt.Println(string(game.underscore)) //Affiche le mot avec les underscores modifié
		if mot == string(game.underscore) {  //condition pour finir le jeu si gagné
			fmt.Println("Congratulation You found the word")
			game.win++
			fmt.Println("Vous avez gagné", game.win, "fois")
			Restart()
			break
		}

		if game.isfalse == true { //si le mot ou l'input entrée est fausse il rentre dans la condition
			fmt.Println("Dommage cette lettre n'est pas dans ce mot")
			fmt.Print("\n")
			fmt.Println("Il vous reste", game.essai, "essai(s)")
			game.essai--
			Draw(game.attempts)
			game.attempts++
			game.p = append(game.p, tabrun[0], 32)
			fmt.Println(string(game.underscore))
			fmt.Println("les mauvaises lettres entrée sont :", string(game.p))
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

func Restart() {
	reader := bufio.NewReader(os.Stdin)

	for {
		print("Voulez-vous rejouer ? (oui/non) ")
		restart, _ := reader.ReadString('\n')

		if restart != "oui\n" && restart != "non\n" {
			println("Erreur veuillez utilisé oui ou non")
		} else {
			switch restart {
			case "oui\n":
				Pendu()
				os.Exit(0)
			case "non\n":
				os.Exit(1)
			}
		}
	}
}

// else {
// 	fmt.Println("Il vous reste", game.essai, "essai")
// 	game.essai = game.essai - 2
// 	Draw(game.attempts)
// 	game.attempts = game.attempts + 2
// 	game.p = append(game.p, tabrun[0], 32)
// 	fmt.Println(string(game.underscore))
// 	fmt.Println("les mauvaises lettres entrée sont :", string(game.p))
// 	game.isfalse = false
// }
