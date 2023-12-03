package fonction

import (
	"encoding/json"
	"fmt"
	"groupie/Data"
	"net/http"
	"strings"
	"time"
)

var Artists []Data.Artists
var Dates Data.Dates
var Client *http.Client
var Location Data.Locations
var Relations Data.Relation

func GetArtists() {
	url := "https://groupietrackers.herokuapp.com/api/artists"
	err := GetJson(url, &Artists)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// Prends les dates
func GetDates() {
	url := "https://groupietrackers.herokuapp.com/api/dates"
	err := GetJson(url, &Dates)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func GetJson(url string, target interface{}) error {
	Client = &http.Client{Timeout: 10 * time.Second}
	resp, err := Client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}

func GetLocations() {
	url := "https://groupietrackers.herokuapp.com/api/locations"
	err := GetJson(url, &Location)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func GetRelations() {
	url := "https://groupietrackers.herokuapp.com/api/relation"
	err := GetJson(url, &Relations)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func WordToMin(word_tab []string) {
	mot := strings.Join(word_tab, "")
	runes := []rune(mot)
	var result []int
	for i := 0; i < len(runes); i++ {
		result = append(result, int(runes[i]))
	}
	//word[i] = string(intVar - 32)
	for i := 0; i < len(mot); i++ {
		if result[i] > 96 {
			result[i] = result[i] - 32
			word_tab[i] = string(result[i])
		}
	}
}
