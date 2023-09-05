package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"hangman/gamestruct"
	"hangman/outils"
	"io/ioutil"
	"log"
	"os"
)

type Lettrutiliser struct {
	lettre [30]string
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error syntax : (go run main.go <file.txt>)")
		os.Exit(0)
	}
	var a Lettrutiliser
	index := 0
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error: File not found")
		os.Exit(0)
	}
	b, _ := ioutil.ReadAll(file)
	word := outils.Motrandom(b)
	nletter := len(word)/2 - 1
	hword := outils.Nouveaumot(word, nletter)
	hp := 10
	tested := ""
	fmt.Print("Bienvenue sur le jeu hangman." + "\n" + "\n" + "\n" + "\n" + "\n")
	fmt.Print("Bonne chance BG ta , 10 essaie." + "\n" + "\n" + "\n" + "\n" + "\n")
	fmt.Print("Veuillez tapper < Entrer>")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println("veuillez rentrer une valeur")
		fmt.Println("lettres testées: ", a.lettre)
		uti := scanner.Text()
		fmt.Println(hword)
		if uti == "stop" {
			os.Create("save.txt")
			u, _ := json.Marshal(gamestruct.Gamesave{Wordtested: tested, Health: hp, Fullword: word, Hiddenword: hword, Tab: a.lettre})
			os.WriteFile("save.json", u, 5)
			fmt.Println("game stopped")
			os.Exit(0)
		}
		tested += uti
		tested += ","
		if uti == "open save" {
			fileContent, err := os.Open("save.json")
			if err != nil {
				log.Fatal(err)
				return
			}

			fmt.Println("The File is opened successfully...")

			defer fileContent.Close()
			byteResult, _ := ioutil.ReadAll(fileContent)
			var temp gamestruct.Gamesave
			json.Unmarshal(byteResult, &temp)
			hp = temp.Health + 2
			tested = temp.Wordtested
			word = temp.Fullword
			hword = temp.Hiddenword
			a.lettre = temp.Tab
			fmt.Println(temp)
		}
		if uti == "" {
			fmt.Println("veuillez rentrer une valeur svp")
		}
		if uti == word {
			fmt.Println("vous avez gagné!")
			os.Exit(0)
		}
		if len(uti) > 1 && uti != word {
			fmt.Print("Mauvais mot, -2 vies!")
			hp -= 2
			outils.Pendu(hp)
			if hp <= 0 {
				fmt.Println("Vous avez perdu")
				fmt.Println("Le mot était: ", word)
				os.Exit(0)
			} else {
				fmt.Println("Nombre de vies restantes: ", hp)
			}
		}
		if len(uti) == 1 {
			if outils.Dejatester(uti, a.lettre) {
				fmt.Println("lettre déjà testée")
			} else if outils.Tcheking(uti, word) {
				a.lettre[index] = uti
				index++
				fmt.Println("bien joué")
				hword = outils.Bonnelettre(uti, word, hword)
				fmt.Println(hword)
				if hword == word || hword == word+" " {
					fmt.Println("vous avez gagné")
					os.Exit(0)
				}
			} else {
				a.lettre[index] = uti
				index++
				if hp > 1 {
					fmt.Println("Mauvaise lettre, -1 vie")
					hp--
					outils.Pendu(hp)
					fmt.Println("nombre de vies restantes", hp)
					fmt.Println("lettres testées: ", a.lettre)
					fmt.Println(hword)
				} else {
					hp--
					outils.Pendu(hp)
					fmt.Println("perdu")
					fmt.Println("le mot était: ", word)
					os.Exit(0)
				}
			}
		}
	}
}
