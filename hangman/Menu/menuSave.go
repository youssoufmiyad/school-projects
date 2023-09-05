package Menu

import (
	"Hangman/FileReader"
	"Hangman/Jose"
	"Hangman/RGB"
	"Hangman/UserInput"
	"fmt"
)

var (
	counterPositionHangman int
	counterAttempts        int
)

func MenuSave() {
	data := FileReader.GetSaveJsonData()
	Hangman := &data
	counterPositionHangman = Hangman.GetSavePositionHangman()
	counterAttempts = Hangman.GetAttempts()

	fmt.Printf("You have %v attempts left\n", Hangman.GetAttempts())

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
				counterPositionHangman += 2
				counterAttempts -= 2

				if counterAttempts <= 0 {
					Jose.EssaiSave(counterAttempts, Hangman)
				} else {
					fmt.Println(Hangman.GetPositionHangman()[counterPositionHangman])
					Hangman.SetSavePositionHangman(counterPositionHangman)
					Jose.EssaiSave(counterAttempts, Hangman)
				}
			} else {
				newString = Hangman.WordToFind
				UserInput.Win(newString, Hangman)
			}
		} else {
			if !UserInput.IsLetterCorrect(inputUser, Hangman) {
				counterPositionHangman++
				counterAttempts--
				fmt.Println(Hangman.GetPositionHangman()[counterPositionHangman])
				Hangman.SetSavePositionHangman(counterPositionHangman)
				Jose.EssaiSave(counterAttempts, Hangman)
			}
		}
	}
}
