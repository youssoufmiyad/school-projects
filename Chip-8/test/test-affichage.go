package game

import (
	"chip-8/CPU"
	"fmt"
	"image/color"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  = 64
	screenHeight = 32
)

var chip *CPU.CHIP8

type Game struct {
	keys []ebiten.Key
}

// La fonction qui affiche les pixels à l'écran
func (g *Game) Draw(screen *ebiten.Image) {
	for y := 0; y < screenHeight; y++ {
		for x := 0; x < screenWidth; x++ {
			idx := y*screenWidth + x
			if chip.Screen[idx] == 1 {
				screen.Set(x, y, color.White)
			} else {
				screen.Set(x, y, color.Black)
			}
		}
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

// Fonction principale où vous initialiseriez Ebiten et appelleriez drawScreen
func (g *Game) Update() error {
	fmt.Print("actual : ", ebiten.ActualFPS())
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	keyboardHandler(g.keys)

	chip.Play()
	chip.Keyboard = [16]byte{}
	return nil
}

// Lance l'émulateur
func Run(c *CPU.CHIP8) {
	chip = c
	ebiten.SetTPS(450)
	// Initialisation de la fenêtre Ebiten
	if err := ebiten.RunGame(&Game{}); err != nil {
		fmt.Println(err)
	}

}

// gestion des input au clavier
func keyboardHandler(keys []ebiten.Key) {
	// keys = append(keys, ebiten.Key1, ebiten.Key2, ebiten.Key3, ebiten.Key4, ebiten.KeyA, ebiten.KeyZ, ebiten.KeyE, ebiten.KeyR, ebiten.KeyS, ebiten.KeyD, ebiten.KeyF, ebiten.KeyW, ebiten.KeyX, ebiten.KeyC, ebiten.KeyV)
	for _, k := range keys {
		println("key : " + k.String())
		switch k.String() {
		case "Digit1":
			chip.Keyboard[1] = 1
		case "Digit2":
			chip.Keyboard[2] = 1
		case "Digit3":
			chip.Keyboard[3] = 1
		case "Digit4":
			chip.Keyboard[12] = 1
		case "Q":
			chip.Keyboard[4] = 1
		case "W":
			chip.Keyboard[5] = 1
		case "E":
			chip.Keyboard[6] = 1
		case "R":
			chip.Keyboard[13] = 1
		case "A":
			chip.Keyboard[7] = 1
		case "S":
			chip.Keyboard[8] = 1
		case "D":
			chip.Keyboard[9] = 1
		case "F":
			chip.Keyboard[14] = 1
		case "Z":
			chip.Keyboard[10] = 1
		case "X":
			chip.Keyboard[0] = 1
		case "C":
			chip.Keyboard[11] = 1
		case "V":
			chip.Keyboard[15] = 1
		case "Escape":
			panic("Fin du programe")
		}
	}
}
