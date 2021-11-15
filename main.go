package main

import "hangman/Hangman"

func main() {
	var game Hangman.Game
	Hangman.ReadFiles(&game)
}
