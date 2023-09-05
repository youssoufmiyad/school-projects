package Jose

import (
	"Hangman/HangmanStructure"
	"bufio"
	"log"
	"os"
)

// Declaration of variables
var (
	test_str string
	compteur int
	position []string = make([]string, 10)
)

// This function reads the file 'hangman.txt' and based on the position given (nb int) it will show the "hangman"
func Image_jose(Hangman *HangmanStructure.HangmanData) {
	// We open the file 'words.txt' so that we can read it
	file, err := os.Open("hangman.txt")

	// In case there is an error, we will get a log with the details of the error
	// The error is retrieved when opening the file by putting it in a variable
	if err != nil {
		log.Fatal(err)
	}

	// We create a scanner so that we can read the file, line by line
	scanner := bufio.NewScanner(file)

	// This will loop until there is nothing to scan (read) anymore
	for scanner.Scan() {
		// We verify if the current line is not empty. If it's not, we will add the current line to a string using scanner.Text()
		if scanner.Text() != "" {
			test_str += "\n" + scanner.Text()
		} else {
			// If the line is now empty, that mean we have finished reading the current position of the "hangman"
			// We'll add what we have read (test_str) and add it to the 'positino' ([]string)
			// Next, we reset the string and add one to our counter so that when we add another string
			// It will not overlap the current one
			position[compteur] = test_str
			test_str = ""
			if compteur < 10 {
				compteur++
			}
		}

	}

	// In case there is an error with the scanner, we will get a log with the details of the error
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	Hangman.SetPositionHangman(position)
}
