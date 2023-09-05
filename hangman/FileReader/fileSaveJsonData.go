package FileReader

import (
	"Hangman/HangmanStructure"
	"encoding/json"
	"io/ioutil"
)

// Function that opens the save file and
func GetSaveJsonData() HangmanStructure.HangmanData {
	// We read the file 'save.json'
	file, _ := ioutil.ReadFile("save.json")

	// Create a variable 'HangmanData' so that we can store our data into it
	var HangmanData HangmanStructure.HangmanData

	// We take the data from the JSON file and put in into 'Hangman'
	// This data will be an array of bytes
	_ = json.Unmarshal([]byte(file), &HangmanData)

	// Return of the 'HangmanData' containing all of the data from the JSON file
	return HangmanData
}
