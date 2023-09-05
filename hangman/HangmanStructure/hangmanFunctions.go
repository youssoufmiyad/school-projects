package HangmanStructure

import (
	"math/rand"
	"time"
)

// Declaration of variables
var (
	continuer    bool = true
	randomNumber int
	numberReveal int
	min          int
	tabInt       []int
	counter      int
	tabLetters   []string
)

// This function gets a random word from the txt file and adds it to the structure 'HangmanData'
func GetRandomWordFromList(Hangman *HangmanData) {
	// Using a random number generator can potentially generate the same number multiple times
	// By using this, we can have a different output everytime
	rand.Seed(time.Now().UnixNano())
	// Explanation (Official Documentation) : "Top-level functions, such as Float64 and Int, use a default shared Source that produces
	// a deterministic sequence of values each time a program is run"

	randomNumber = rand.Intn(len(Hangman.GetRandomWordsFile()))

	// The randomNumber (from above) is used as an index to see which word from the list will be chosen
	for index, word := range Hangman.RandomWordsFile {
		if index == randomNumber {
			Hangman.SetWordToFind(word)
		}
	}
}

// This function returns an array of integers which will be used as index for the random word that was chosen
// The words revealed letters will be based on this array of integers
// EX : If the array of integers contains 1, then the index 1 of the word (which means the second letter) will be revealed
// All other letters will be represented by a _
func RevealLetters(Hangman *HangmanData) []int {
	// This is the formula that we need to use to get the amount of letters that we will need to reveal
	numberReveal = len(Hangman.GetWordToFind())/2 - 1

	// Sometimes it can happen that the formula gives a number smaller than 1
	// In that case (because we still need to reveal a letter) the variable will be set to 1
	if numberReveal < 1 {
		numberReveal = 1
	}

	// This is where we create the array that will contain random integers
	for i := 0; i < numberReveal; i++ {
		max := len(Hangman.GetWordToFind())
		tabInt = append(tabInt, (rand.Intn(max-min+1)+min)-1)
		if tabInt[i] < 0 {
			tabInt[i] = 0
		}
	}

	// We can't have two of the same number in our array so we call a function that will check that for us
	isUnique(tabInt, Hangman)
	return tabInt
}

// This function verifies if in the array of integers is composed of only unique numbers (no duplicates like [5 5 0])
// If it's not, the function will choose another number for that index
func isUnique(tabInt []int, Hangman *HangmanData) []int {
	max := len(Hangman.GetWordToFind())

	for continuer {
		for i := 0; i < len(tabInt); i++ {
			counter = 0
			for j := 0; j < len(tabInt); j++ {
				if i != j && tabInt[i] == tabInt[j] {
					counter++
					tabInt[i] = (rand.Intn(max-min+1) + min) - 1
				}
			}
		}
		if counter == 0 {
			continuer = false
		}
	}

	return tabInt
}

// This function verifies (using the array of integers from the function 'RevealLetters') which
// letter (based on the index) will need to change and which will stay normal (normal = revealed letters)
// EX : this function will get [0 3 2] which means that the first, forth and third letters will be revealed
func ChangeLetter(Hangman *HangmanData, tabInt []int) {
	wordToFind := Hangman.GetWordToFind()

	for index, letter := range wordToFind {
		counter = 0
		for i := 0; i < len(tabInt); i++ {
			if index != tabInt[i] {
				counter++
			}
		}
		// Taking the example from above ([0 3 2])
		// If the index is 1, the counter will be the same as the length of the array, which means that we don't want
		// this index to change, which is correct
		// On the other hand index 0, will not get the same length as the array which means that there is a 0 in the array
		// This means that we will need to change this index 0 into its actual letter and not _
		if counter == len(tabInt) {
			tabLetters = append(tabLetters, "_")
		} else {
			tabLetters = append(tabLetters, string(letter))
		}
	}

	Hangman.SetWord(tabLetters)
}
