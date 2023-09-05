package Menu

import (
	"Hangman/HangmanStructure"
	"Hangman/Jose"
	"Hangman/RGB"
	"Hangman/UserInput"
	"fmt"
)

var (
	newString string
	counter   int
	inputUser string
)

func MenuNewGame(Hangman *HangmanStructure.HangmanData) {
	for {
		// To be able to more easily modify the word that will show the letters that the user has guessed
		// We made it into an array. This means that before showing it, we need to "assemble it" into an actual string
		newString = ""
		for _, letter := range Hangman.GetWord() {
			newString += letter
		}

		// We check if the user has won
		UserInput.Win(newString, Hangman)

		fmt.Printf(RGB.RGB_Text(0, 0, 255, "The word is : %v \n"), newString)

		inputUser = UserInput.ScanUserInput(Hangman)

		// If the length is superior than 1, that will be considered as a 'word'
		if len(inputUser) > 1 {
			if !UserInput.IsWordCorrect(inputUser, Hangman) {
				if counter >= 10 {
					Jose.Essai(counter, Hangman)
				} else {
					fmt.Println(Hangman.GetPositionHangman()[counter])
					Hangman.SetSavePositionHangman(counter)
					counter += 2
					Jose.Essai(counter, Hangman)
				}
			} else {
				newString = Hangman.WordToFind
				UserInput.Win(newString, Hangman)
			}
		} else {
			if !UserInput.IsLetterCorrect(inputUser, Hangman) {
				if counter >= 10 {
					Jose.Essai(counter, Hangman)
				} else {
					fmt.Println(Hangman.GetPositionHangman()[counter])
					Hangman.SetSavePositionHangman(counter)
					counter++
					Jose.Essai(counter, Hangman)
				}
			}
		}
	}
}
