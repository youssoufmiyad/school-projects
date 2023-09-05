package Jose

import (
	"Hangman/HangmanStructure"
	"Hangman/RGB"
	"Hangman/UserInput"
	"fmt"
	"os"
)

// This function shows the number of attempts left and can also verify if the user has lost
func Essai(nb int, Hangman *HangmanStructure.HangmanData) {
	if nb >= 10 {
		fmt.Println(RGB.RGB_Text(255, 0, 0, "GAME OVER"))
		Hangman.SetGameFinished(true)
		UserInput.Save(Hangman)
		os.Exit(0)
	} else {
		fmt.Printf("You have %v attempts left \n\n", 10-nb)
		Hangman.SetAttempts(10 - nb)
	}
}

// This function shows the number of attempts left and can also verify if the user has lost
// When we save, the number that this function will receive will be different
// This function receives directly the number of attempts while the other function receive a counter that is
// used as a substractor
func EssaiSave(nb int, Hangman *HangmanStructure.HangmanData) {
	if nb <= 0 {
		fmt.Println(RGB.RGB_Text(255, 0, 0, "GAME OVER"))
		Hangman.SetGameFinished(true)
		UserInput.Save(Hangman)
		os.Exit(0)
	} else {
		fmt.Printf("You have %v attempts left \n", nb)
		Hangman.SetAttempts(nb)
	}
}
