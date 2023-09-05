package main

import (
	"Hangman/FileReader"
	"Hangman/HangmanStructure"
	"Hangman/Jose"
	"Hangman/Menu"
	"Hangman/RGB"
	"flag"
	"fmt"
	"os"
)

var (
	choice    string
	startWith string
)

// This function will be called the moment we run our program
// In this function, we initialize a flag that we can specify if we want to run the program with the save.json as a parameter
func init() {
	flag.StringVar(&startWith, "startWith", "", "Specify save file")
}

func main() {
	// This needs to be run for the flags to work as a command line
	flag.Parse()

	Hangman := new(HangmanStructure.HangmanData)
	Hangman.SetRandomWordsFile(FileReader.GetWordsFile())
	Jose.Image_jose(Hangman)
	HangmanStructure.GetRandomWordFromList(Hangman)
	HangmanStructure.ChangeLetter(Hangman, HangmanStructure.RevealLetters(Hangman))

	fmt.Printf("\x1bc")
	fmt.Println("WELCOME TO HANGMAN ! PRESS ENTER TO START")
	fmt.Scanln()
	fmt.Println(RGB.RGB_Text(255, 0, 0, "TYPING 'STOP' WILL SAVE YOUR PROGRESS AND EXIT THE GAME"))

	data := FileReader.GetSaveJsonData()
	Hangman.SetGameFinished(data.GameFinished)

	if _, err := os.Stat("save.json"); err == nil && !Hangman.GetGameFinished() || startWith == "save.json" {
		for {
			fmt.Println("1. Continue")
			fmt.Println("2. Exit")
			fmt.Scanln(&choice)
			switch choice {
			case "1":
				Menu.MenuSave()
			case "2":
				os.Exit(0)
			default:
				fmt.Println(RGB.RGB_Text(255, 0, 0, "You must choose a valid option !"))
			}
		}
	} else {
		for {
			fmt.Println("1. New Game")
			fmt.Println("2. Exit")
			fmt.Scanln(&choice)
			switch choice {
			case "1":
				Menu.MenuNewGame(Hangman)
			case "2":
				os.Exit(0)
			default:
				fmt.Println(RGB.RGB_Text(255, 0, 0, "You must choose a valid option !"))
			}
		}
	}
}
