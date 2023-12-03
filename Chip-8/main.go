package main

import (
	"chip-8/CPU"
	game "chip-8/test"
	"fmt"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

var c *CPU.CHIP8
var ROM string
var play bool

func main() {
	c = CPU.ChipInit()
	// Dans le cas ou l'utilisateur connait les code correspondant aux ROMS
	if len(os.Args) > 1 {
		ROM = os.Args[1]
		play = true
	} else {
		// Dans l'autre cas
		fmt.Print("Choose a ROM : ")
		fmt.Println("1-chip8-logo 2-ibm-logo 3-corax+ 4-flags 5-quirks 6-keypad 7-breakout other-exit")
		fmt.Scan(&ROM)
	}
	switch ROM {
	case "1":
		c.ROM = "1-chip8-logo"
		play = true
	case "2":
		c.ROM = "2-ibm-logo"
		play = true
	case "3":
		c.ROM = "3-corax+"
		play = true
	case "4":
		c.ROM = "4-flags"
		play = true
	case "5":
		c.ROM = "5-quirks"
		play = true
	case "6":
		c.ROM = "6-keypad"
		play = true
	case "7":
		c.ROM = "Breakout (Brix hack) [David Winter, 1997]"
		play = true
	default:
		play = false
	}

	if play {
		c.ReadRoms()
		ebiten.SetWindowTitle("Chip8")
		game.Run(c)
	}

}
