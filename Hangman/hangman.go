package Hangman

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

type Game struct { //la structure qui nous servira tout au long du code
	tab        []string //Prends tout les mots du fichier words.txt
	word       []string // utilisé pour entrée un mot complet lors d'une partie
	input      string
	mot        string
	underscore []rune //utilisé pour modifié les bonnes lettre avec les underscores
	letter     []rune
	tabmot     []rune // découpe notre mot en rune pour comparer avec la rune de input
	tabrun     []rune // rune de input
	attempts   int    // nombre de tours utilisé en compteur
	essai      int    // nombre d'essai
	win        int    // utilisé en compteur pour compté les parties gangés
	loose      int    // utilisé en compteur pour compté les parties perdu
	isfalse    bool   // booléen utilisé pour passer des condition des bonne lettre au mauvaise entréé
	gagné      bool
	alreadyuse bool
	err        error
}

func ReadFiles(game *Game) { //fonction qui nous permet de lire le fichier de mot
	file, err := os.Open("words.txt") // ouvre le document

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close() // ferme le document

	scanner := bufio.NewScanner(file) //recupere le mot

	for scanner.Scan() {
		game.tab = append(game.tab, scanner.Text())
		if err := scanner.Err(); err != nil {
			log.Fatalln(err)
		}
	}
	game.essai = 10
	game.isfalse = true
	RandomString(game)
}

func RandomString(game *Game) { //selectionner un mot aleatoirement dans word .txt
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := len(game.tab)
	i := rand.Intn((max - min))
	game.mot = game.tab[i]
	game.tabmot = []rune(game.mot)
	Underscore(game)
}

func Underscore(game *Game) { //remplacer le mot par des underscores
	for u := 0; u < len(game.tabmot); u++ {
		game.underscore = append(game.underscore, '_')
	}
	PrintRandomLetter(game)
}

func PrintRandomLetter(game *Game) {
	for g := 0; g < (len(game.mot)/2 - 1); g++ { //affichage la moitié -1 des lettres du mot
		A := rand.Intn(len(game.mot)) //choisi aléatoirement les lettres à afficher du mot
		game.underscore[A] = game.tabmot[A]
	}
	Start(game)
}

func Start(game *Game) { // message de debut de partie
	fmt.Println("Bonne chance vous avez 10 essais")
	fmt.Println(string(game.underscore))
	fmt.Print("Choisissez votre lettre: ")
	Input(game)
}

func Input(game *Game) { //recupere l'entree de l'utilisateur
	reader := bufio.NewReader(os.Stdin)
	if game.err != nil {
		println(game.err)
	}

	game.input, _ = reader.ReadString('\n') //Lit ce que l'on écrit
	game.tabrun = []rune(game.input)        //tableau de rune de ce que l'on écrit
	game.isfalse = true
	game.gagné = false
	game.alreadyuse = false
	VerifyInput(game)
}

func VerifyInput(game *Game) { //verifie si l'input est la meme que le mot
	inputRune := game.tabrun[0]
	for b := 0; b < len(game.letter); b++ { //verifie si l'imput n'as pas deja ete utilise dans la partie
		if inputRune == game.letter[b] {
			fmt.Println("Cette lettre est déjà entrée tu fais quoi là !")
			game.alreadyuse = true
		}
	}
	for x := 0; x < len(game.word); x++ { //verifie si l'imput n'as pas deja ete utilise dans la partie
		if game.input == game.word[x] {
			fmt.Println("Ce mot est déjà entrée tu fais quoi là !")
			game.alreadyuse = true
		}
	}
	for x := 0; x < len(game.tabmot); x++ {
		if game.mot == string(game.tabrun[:len(game.tabrun)-1]) { //condition qui vérifie si le mot correspond à ce que l'on marque moins le \n
			game.underscore[x] = game.tabmot[x]
			game.isfalse = false
			game.gagné = true
			Win(game)
		} else if len(game.input) <= 2 {
			if game.tabmot[x] == game.tabrun[0] { //compare l'index du mot a l'index de mon input
				game.underscore[x] = game.tabmot[x]
				game.isfalse = false
			}
		}
		if game.mot == string(game.underscore) {
			game.gagné = true
			Win(game)
		}
	}
	if game.isfalse == true {
		False(game)
	}
	Win(game)
}

func OneLetter(game *Game) { //affiche les lettres utilise et retire le nombre d'essai
	var reset string
	fmt.Println("Dommage cette lettre n'est pas dans ce mot", "\n")
	game.essai = game.essai - 1
	fmt.Println("Il vous reste", game.essai, "essai(s)")
	Draw(game.attempts)
	game.attempts = game.attempts + 1
	game.letter = append(game.letter, game.tabrun[0], 32) //afficher les mauvaises lettres
	fmt.Println(string(game.underscore))
	fmt.Println("les mauvais mots entrée sont :", game.word)
	fmt.Println("les mauvaises lettres entrée sont :", string(game.letter))
	game.input = reset
	Loose(game)
}

func OneWord(game *Game) { //affiche les mots utilise et retire le nombre d'essai
	var reset string
	fmt.Println("Dommage ce mot ne correspond pas", "\n")
	game.essai = game.essai - 2
	fmt.Println("Il vous reste", game.essai, "essai(s)")
	game.attempts = game.attempts + 2
	Draw(game.attempts - 1)
	game.word = append(game.word, game.input) //afficher les mauvais mots
	fmt.Println(string(game.underscore))
	fmt.Println("les mauvais mots entrée sont :", game.word)
	fmt.Println("les mauvaises lettres entrée sont :", string(game.letter))
	game.input = reset
	Loose(game)
}

func LetterByLetter(game *Game) {
	fmt.Println(string(game.underscore))
	fmt.Print("Choisissez votre lettre: ")
	Input(game)
}

func Loose(game *Game) { //condition pour finir le jeu si perdu
	var reset int
	if game.essai <= 0 {
		Draw(game.attempts)
		fmt.Println("You're dead")
		game.loose++
		fmt.Println("Vous avez perdu", game.loose, "fois")
		fmt.Println("vous avez gagné", game.win, "fois")
		game.attempts = reset
		Restart(game)
	}
	fmt.Print("Choisissez votre lettre: ")
	Input(game)
}

func Win(game *Game) { //condition pour finir le jeu si gagné
	var reset int
	if game.gagné == true {
		fmt.Println("Congratulation You found the word")
		game.win++
		fmt.Println("Vous avez gagné", game.win, "fois")
		fmt.Println("Vous avez perdu", game.loose, "fois")
		game.attempts = reset
		Restart(game)
	} else {
		LetterByLetter(game)
	}
}

func False(game *Game) { //condition d'erreur
	for game.isfalse == true && game.alreadyuse == false {
		if len(game.input) <= 2 {
			OneLetter(game)
		} else {
			OneWord(game)
		}
		if !game.alreadyuse {
			game.letter = append(game.letter, game.tabrun[0], 32)
			game.word = append(game.word, string(game.input))
		}
	}
}

func Draw(attempts int) { //affichage du pendu
	count := 0

	file2, erreur := os.Open("hangman.txt") //ouvre le fichier des positions du pendu
	if erreur != nil {
		log.Fatal(erreur)
	}

	defer file2.Close()                 //fermer le fichier mot
	position := bufio.NewScanner(file2) //Lit mon fichier words
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

func Restart(game *Game) { // gère la reprise de la partie
	var resetrune []rune
	var resetstring []string
	reader := bufio.NewReader(os.Stdin)
	for {
		print("Voulez-vous rejouer ? (oui/non) ")
		restart, _ := reader.ReadString('\n')
		if restart != "oui\n" && restart != "non\n" {
			println("Erreur veuillez utilisé oui ou non")
		} else {
			switch restart {
			case "oui\n":
				game.tabmot = resetrune     //clear pour le choix d'un nouveau mot
				game.underscore = resetrune //clear pour le nouveau mot
				game.letter = resetrune     //clear pour les mauvaises lettres
				game.word = resetstring     //clear pour les mauvais mots
				ReadFiles(game)
				os.Exit(0)
			case "non\n":
				os.Exit(1)
			}
		}
	}
}
