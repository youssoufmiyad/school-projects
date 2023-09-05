package UserInput

import (
	"Hangman/HangmanStructure"
	"Hangman/RGB"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"unicode"
)

// Declaration of variables
var (
	userInput     string
	repeatLetters []string
)

// This function takes the input of a user and checks multiple things
// If the input is only letters and if the input is "stop" which will stop the program and save
func ScanUserInput(Hangman *HangmanStructure.HangmanData) string {
	fmt.Print("Choose a letter / word : ")
	fmt.Scanln(&userInput)
	fmt.Println()

	if !onlyLetters(userInput) {
		fmt.Println(RGB.RGB_Text(255, 0, 0, "ONLY LETTERS ALLOWED !"))
		ScanUserInput(Hangman)

	} else if userInput == "STOP" || userInput == "stop" || userInput == "Stop" {
		Hangman.SetGameFinished(false)
		Save(Hangman)

	} else {
		return userInput
	}

	return ""
}

// Function that changes the letters that the user has guessed correctly
func guessLetterChange(tabIndex []int, Hangman *HangmanStructure.HangmanData) bool {
	wordToFind := Hangman.GetWordToFind()

	// If the array is empty that means that the user has guessed a letter that is not present in the word
	if len(tabIndex) < 1 {
		return false
	}

	for i := 0; i < len(tabIndex); i++ {
		for index, letter := range wordToFind {
			if tabIndex[i] == index {
				Hangman.GetWord()[index] = string(letter)
			}
		}
	}

	return true
}

// This function verifies if the letter chosen by the user is part of the word that he has to find
// It returns an array of integers that indicate the index of the letters that we need to show because
// the user has guessed them
func IsLetterCorrect(letterVerify string, Hangman *HangmanStructure.HangmanData) bool {
	for i := 0; i < len(repeatLetters); i++ {
		if repeatLetters[i] == letterVerify {
			fmt.Println(RGB.RGB_Text(255, 0, 0, "This letter has already been guessed !"))
			return false
		}
	}

	repeatLetters = append(repeatLetters, letterVerify)

	wordToFind := Hangman.GetWordToFind()
	tabInt := []int{}

	for index, letter := range wordToFind {
		if string(letter) == letterVerify {
			tabInt = append(tabInt, index)
		}
	}

	return guessLetterChange(tabInt, Hangman)
}

// This function verifies if the word guessed by the user is the correct word
func IsWordCorrect(wordVerify string, Hangman *HangmanStructure.HangmanData) bool {
	wordToFind := Hangman.GetWordToFind()
	return wordVerify == wordToFind
}

// This function verifies if the word guessed by the user is the correct word
// If that is the case, the user has won
func Win(stringVerify string, Hangman *HangmanStructure.HangmanData) {
	if stringVerify == Hangman.GetWordToFind() {
		fmt.Println(RGB.RGB_Text(0, 255, 0, "YOU WON"))
		fmt.Printf("The word was : %v", Hangman.GetWordToFind())
		Hangman.SetGameFinished(true)
		Save(Hangman)
		os.Exit(0)
	}
}

// This function verifies if a string has only letters
func onlyLetters(stringVerify string) bool {
	for _, letter := range stringVerify {
		if !unicode.IsLetter(letter) {
			return false
		}
	}

	return true
}

// This function will save current Hangman data and put it into a json file
func Save(Hangman *HangmanStructure.HangmanData) {
	if Hangman.GetAttempts() == 0 && !Hangman.GetGameFinished() {
		fmt.Println(RGB.RGB_Text(255, 0, 0, "\nThe game has not been started yet which means that you can not save"))
		os.Exit(0)
	} else {
		file, err := json.MarshalIndent(Hangman, "", " ")
		_ = ioutil.WriteFile("save.json", file, 0644)
	
		if err != nil {
			log.Fatal(err)
		} else {
			os.Exit(0)
		}
	}
}
