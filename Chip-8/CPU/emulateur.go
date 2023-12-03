package CPU

import (
	"os"
	"time"
)

// Structure du CPU
type CHIP8 struct {
	Memory [4096]byte
	//registers
	V              [16]byte
	index_register uint16

	stack [16]uint16
	//stack pointer
	SP uint8

	address byte

	//program counter
	PC          uint16
	sound_timer uint8
	delay_timer uint8

	Keyboard [16]byte
	Screen   [64 * 32]byte
	Clock    time.Ticker
	ROM      string
}

// Initialisation de la structure
func ChipInit() *CHIP8 {
	c := &CHIP8{
		// vide la memoire
		Memory: [4096]byte{},

		// vide les emplacement registers
		V: [16]byte{},

		index_register: 0,

		stack: [16]uint16{},
		SP:    0,

		// pas encore sur de l'usage
		address: 0,

		// cf design specification : "All of the supported programs will start at memory location 0x200."
		PC: 0x200,

		// rempli l'écran de 0s (remplace les potentiels 1 et clear donc l'ecran)
		Screen: [64 * 32]byte{},

		// 60 hz
		Clock: *time.NewTicker(time.Second / 39),
	}

	return c
}

// Fonction enregistrant la ROM dans la mémoire
func (c *CHIP8) ReadRoms() {
	file, err := os.ReadFile("ROMS/" + c.ROM + ".ch8")
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(file); i++ {
		// cf specification technique : "The first 512 bytes, from 0x000 to 0x1FF, are where the original interpreter was located, and should not be used by programs."
		c.Memory[i+512] = file[i]
	}
}

// Fonction parcourant les instructions dans la ROM
func (c *CHIP8) Play() {
	println("Program counter : ", c.PC)
	opcode := uint16(c.Memory[c.PC])<<8 | uint16(c.Memory[c.PC+1])

	Opcodes(opcode, c)
	select {
	case <-c.Clock.C:
		// Update timers
		if c.delay_timer > 0 {
			c.delay_timer--
		}

		if c.sound_timer > 0 {
			c.sound_timer--
		}
	default:
		// Skip the timers
	}

}

// Structure de Debug
type CPUState struct {
	Opcode string
	// name of the opcode
	OCName string
}

func (s CPUState) toString() string {
	return "Opcode : " + s.Opcode + "\nInstruction : " + s.OCName
}
