package FileReader

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Function that gets all words (line by line) from a text file and returns an array that contains those words
func GetWordsFile() []string {
	tabWords := []string{}

	// We open the file 'words.txt' so that we can read it
	file, err := os.Open("words.txt")

	// In case there is an error, we will get a log with the details of the error
	// The error is retrieved when opening the file by putting it in a variable
	if err != nil {
		log.Fatal(err)
	}

	// We create a scanner so that we can read the file, line by line
	scanner := bufio.NewScanner(file)

	// This will loop until there is nothing to scan (read) anymore
	for scanner.Scan() {
		tabWords = append(tabWords, scanner.Text())
	}

	// In case there is an error with the scanner, we will get a log with the details of the error
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	closeFile(file)

	return tabWords
}

// Function that closes the file that has been opened beforehands
func closeFile(file *os.File) {
	err := file.Close()

	if err != nil {
		//os.Stderr = Standard Error
		//Fprintf = a variation of Printf which allows us to use multiple arguments
		fmt.Fprintf(os.Stderr, "error : %v\n", err)
		//Exit the program
		os.Exit(1)
	}
}
