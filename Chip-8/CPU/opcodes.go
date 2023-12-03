package CPU

import (
	"fmt"
	"math/rand"
)

var s *CPUState

func Opcodes(code uint16, cpu *CHIP8) {
	s = &CPUState{}
	switch code & 0xF000 {
	case 0x0000:
		Opcodes_0(code, cpu)
	case 0x1000:
		Opcodes_1(code, cpu)
	case 0x2000:
		Opcodes_2(code, cpu)
	case 0x3000:
		Opcodes_3(code, cpu)
	case 0x4000:
		Opcodes_4(code, cpu)
	case 0x5000:
		Opcodes_5(code, cpu)
	case 0x6000:
		Opcodes_6(code, cpu)
	case 0x7000:
		Opcodes_7(code, cpu)
	case 0x8000:
		Opcodes_8(code, cpu)
	case 0x9000:
		Opcodes_9(code, cpu)
	case 0xA000:
		Opcodes_A(code, cpu)
	case 0xB000:
		Opcodes_B(code, cpu)
	case 0xC000:
		Opcodes_C(code, cpu)
	case 0xD000:
		Opcodes_D(code, cpu)
	case 0xE000:
		Opcodes_E(code, cpu)
	case 0xF000:
		Opcodes_F(code, cpu)
	}
}

// le nom des fonctions correspond à la liste de Opcodes
// géré par ces dernières (exemple : "Opcodes_8" gère tout les Opcodes en "0x8...").
// "nnn" étant les 3 dernier caractere du byte
// "kk" étant les 2 dernier caractere du byte
// "x" étant le 2e caractere du byte
// "y" étant le 3e caractere du byte

// ANCHOR - Opcodes 0x0...
func Opcodes_0(code uint16, cpu *CHIP8) {
	switch code {
	// en cas de "clears"
	// reset l'écran a son état initial (tout noir)
	case 0x00E0:
		s.Opcode = "0x00E0"
		s.OCName = "CLS"

		for i := range cpu.Screen {
			cpu.Screen[i] = 0
		}
		// passe à la prochaine instruction (chaque opcodes est inscrit sur 2 byte)
		cpu.PC += 2
	case 0x00EE:
		s.Opcode = "0x00EE"
		s.OCName = "RET"
		// The return address stack stores previous program counters when jumping into a new routine.
		cpu.address = byte(cpu.PC)
		cpu.PC = cpu.stack[cpu.SP] + 2
		cpu.SP--
	default:
		panic(fmt.Errorf("Invalid opcode : ", code, " doesnt exist"))
	}
	println(s.toString())
}

// ANCHOR - Opcodes 0x1...
// Jump to location nnn
func Opcodes_1(code uint16, cpu *CHIP8) {
	nnn := code & 0x0FFF
	cpu.PC = nnn

	s.Opcode = "0x1" + fmt.Sprint(nnn)
	s.OCName = "JP " + fmt.Sprint(nnn)

	println(s.toString())
}

// ANCHOR - Opcodes 0x2...
// Call subroutine at nnn
func Opcodes_2(code uint16, cpu *CHIP8) {
	nnn := code & 0x0FFF
	cpu.SP++
	cpu.stack[cpu.SP] = cpu.PC
	cpu.PC = nnn
	s.Opcode = "0x2" + fmt.Sprint(nnn)
	s.OCName = "CALL 0x" + fmt.Sprint(nnn)
	println(s.toString())
}

// ANCHOR - Opcodes 0x3...
// Skip next instruction if Vx = kk ()
func Opcodes_3(code uint16, cpu *CHIP8) {
	x := (code & 0x0F00) >> 8
	kk := byte(code & 0x00FF)
	if cpu.V[x] == kk {
		cpu.PC += 4
	} else {
		cpu.PC += 2
	}
	s.Opcode = "0x3" + fmt.Sprint(x) + fmt.Sprint(kk)
	s.OCName = "SE V" + fmt.Sprint(x) + " " + fmt.Sprint(kk)
	println(s.toString())
}

// ANCHOR - Opcodes 0x4...
// Skip next instruction if Vx != kk
func Opcodes_4(code uint16, cpu *CHIP8) {
	x := (code & 0x0F00) >> 8
	kk := byte(code & 0x00FF)

	if cpu.V[x] != kk {
		cpu.PC += 4
	} else {
		cpu.PC += 2
	}

	s.Opcode = "0x4" + fmt.Sprint(x) + fmt.Sprint(kk)
	s.OCName = "SNE V" + fmt.Sprint(x)
	println(s.toString())
}

// ANCHOR - Opcodes 0x5...
// Skip next instruction if Vx = Vy
func Opcodes_5(code uint16, cpu *CHIP8) {
	x := (code & 0x0F00) >> 8
	y := (code & 0x00F0) >> 4

	if cpu.V[x] == cpu.V[y] {
		cpu.PC += 4
	} else {
		cpu.PC += 2
	}

	s.Opcode = "0x5" + fmt.Sprint(x) + fmt.Sprint(y) + "0"
	s.OCName = "SE V" + fmt.Sprint(x) + " V" + fmt.Sprint(y)
	println(s.toString())
}

// ANCHOR - Opcodes 0x6...
// Set Vx = kk
func Opcodes_6(code uint16, cpu *CHIP8) {
	x := (code & 0x0F00) >> 8
	kk := byte(code & 0x00FF)

	cpu.V[x] = kk

	s.Opcode = "0x6" + fmt.Sprint(x) + fmt.Sprint(kk)
	s.OCName = "LD V" + fmt.Sprint(x) + " " + fmt.Sprint(kk)
	println(s.toString())

	cpu.PC += 2
}

// ANCHOR - Opcodes 0x7...
// Set Vx = Vx + kk
func Opcodes_7(code uint16, cpu *CHIP8) {
	x := (code & 0x0F00) >> 8
	kk := byte(code & 0x00FF)

	cpu.V[x] += kk

	s.Opcode = "0x7" + fmt.Sprint(x) + fmt.Sprint(kk)
	s.OCName = "LD V" + fmt.Sprint(x)
	println(s.toString())

	cpu.PC += 2
}

// ANCHOR - Opcodes 0x8...
func Opcodes_8(code uint16, cpu *CHIP8) {
	x := (code & 0x0F00) >> 8
	y := (code & 0x00F0) >> 4

	switch code & 0x000F {
	case 0x0000:
		// Set Vx = Vy
		cpu.V[x] = cpu.V[y]
	case 0x0001:
		// Set Vx = Vx OR Vy
		// "|=" = bitwise OR operator
		cpu.V[x] |= cpu.V[y]
		cpu.V[0xF] = 0
		s.Opcode = "0x8" + fmt.Sprint(x) + fmt.Sprint(y) + "1"
		s.OCName = "OR V" + fmt.Sprint(x) + " V" + fmt.Sprint(y)
	case 0x0002:
		// Set Vx = Vx AND Vy
		// "&=" = bitwise AND operator
		cpu.V[x] &= cpu.V[y]
		cpu.V[0xF] = 0
		s.Opcode = "0x8" + fmt.Sprint(x) + fmt.Sprint(y) + "2"
		s.OCName = "AND V" + fmt.Sprint(x) + " V" + fmt.Sprint(y)
	case 0x0003:
		// Set Vx = Vx XOR Vy
		// "^=" = bitwise OR operator
		cpu.V[x] ^= cpu.V[y]
		cpu.V[0xF] = 0
		s.Opcode = "0x8" + fmt.Sprint(x) + fmt.Sprint(y) + "3"
		s.OCName = "XOR V" + fmt.Sprint(x) + " V" + fmt.Sprint(y)
	case 0x0004:
		// Set Vx = Vx + Vy
		cpu.V[x] += cpu.V[y]

		// set VF = carry
		// If the result is greater than 8 bits (i.e., ¿ 255,) VF is set to 1, otherwise 0.
		if cpu.V[x] > 0x000F {
			cpu.V[0xF] = 1
		} else {
			cpu.V[0xF] = 0
		}
		println("VF = ", cpu.V[0xF], " V16 = ", cpu.V[15])
		s.Opcode = "0x8" + fmt.Sprint(x) + fmt.Sprint(y) + "4"
		s.OCName = "ADD V" + fmt.Sprint(x) + " V" + fmt.Sprint(y)
	case 0x0005:
		println("V", x, " = ", cpu.V[x])
		println("V", y, " = ", cpu.V[y])
		println("X : ", x, "Y : ", y)
		// Set Vx = Vx - Vy, set VF = NOT borrow

		if cpu.V[y] > cpu.V[x] {
			cpu.V[0xF] = 0
		} else {
			cpu.V[0xF] = 1
		}
		cpu.V[x] -= cpu.V[y]
		println("V", x, " = ", cpu.V[x])

		s.Opcode = "0x8" + fmt.Sprint(x) + fmt.Sprint(y) + "5"
		s.OCName = "SUB V" + fmt.Sprint(x) + " V" + fmt.Sprint(y)
	case 0x0006:

		// Set Vx = Vx SHR 1

		least := cpu.V[x] & 0x01
		if least == 1 {
			cpu.V[0xF] = 1
		} else {
			cpu.V[0xF] = 0
		}

		cpu.V[x] = cpu.V[x] >> 1

		s.Opcode = "0x8" + fmt.Sprint(x) + fmt.Sprint(y) + "6"
		s.OCName = "SHR V" + fmt.Sprint(x) + " V" + fmt.Sprint(y)
	case 0x0007:
		// Set Vx = Vy - Vx, set VF = NOT borrow
		if cpu.V[y] > cpu.V[x] {
			cpu.V[0xF] = 1
		} else {
			cpu.V[0xF] = 0
		}

		cpu.V[x] = cpu.V[y] - cpu.V[x]
		s.Opcode = "0x8" + fmt.Sprint(x) + fmt.Sprint(y) + "7"
		s.OCName = "SUBN V" + fmt.Sprint(x) + " V" + fmt.Sprint(y)
	case 0x000E:
		// Set Vx = Vx SHL 1
		most := cpu.V[y] & 0x01

		if most == 1 {
			cpu.V[0xF] = 1
		} else {
			cpu.V[0xF] = 0
		}
		cpu.V[x] = cpu.V[x] << 1
		s.Opcode = "0x8" + fmt.Sprint(x) + fmt.Sprint(y) + "E"
		s.OCName = "SHL V" + fmt.Sprint(x) + " V" + fmt.Sprint(y)
	default:
		panic("Invalid opcode :")
	}

	println(s.toString())

	cpu.PC += 2
}

// ANCHOR - Opcodes 0x9...
// Skip next instruction if Vx != Vy
func Opcodes_9(code uint16, cpu *CHIP8) {
	x := (code & 0x0F00) >> 8
	y := (code & 0x00F0) >> 4

	if cpu.V[x] != cpu.V[y] {
		cpu.PC += 4
	} else {
		cpu.PC += 2
	}

	s.Opcode = "0x9" + fmt.Sprint(x) + fmt.Sprint(y) + "0"
	s.OCName = "SNE V" + fmt.Sprint(x) + " V" + fmt.Sprint(y)
	println(s.toString())
}

// ANCHOR - Opcodes 0xA...
// Set I = nnn. "I" étant l'abréviation de "index"
func Opcodes_A(code uint16, cpu *CHIP8) {
	nnn := code & 0x0FFF

	cpu.index_register = nnn

	s.Opcode = "0xA" + fmt.Sprint(nnn)
	s.OCName = "LD I " + fmt.Sprint(cpu.index_register)
	println(s.toString())

	cpu.PC += 2
}

// ANCHOR - Opcodes 0xB...
// Jump to location nnn + V0
func Opcodes_B(code uint16, cpu *CHIP8) {
	nnn := code & 0x0FFF
	n1 := code & 0x0F00
	n2 := code & 0x00F0
	n3 := code & 0x000F
	var highest uint16
	if n1 > n2 && n1 > n3 {
		highest = n1
	} else if n2 > n1 && n2 > n3 {
		highest = n2
	} else if n3 > n1 && n3 > n2 {
		highest = n3
	}

	cpu.PC += (nnn) + uint16(cpu.V[highest])

	s.Opcode = "0xB" + fmt.Sprint(nnn)
	s.OCName = "JP V0"
	println(s.toString())
}

// ANCHOR - Opcodes 0xC...
// Set Vx = random byte AND kk
func Opcodes_C(code uint16, cpu *CHIP8) {
	x := (code & 0x0F00) >> 8
	kk := byte(code & 0x00FF)

	cpu.V[x] = byte(rand.Float32()*255) & kk

	s.Opcode = "0xC" + fmt.Sprint(x) + fmt.Sprint(kk)
	s.OCName = "RND V" + fmt.Sprint(x)
	println(s.toString())

	cpu.PC += 2
}

// ANCHOR - Opcodes 0xD...
// Display n-byte sprite starting at memory location I at (Vx, Vy), set VF = collision
func Opcodes_D(code uint16, cpu *CHIP8) {
	x := uint16(cpu.V[code&0x0F00>>8])
	y := uint16(cpu.V[code&0x00F0>>4])
	n := code & 0x000F
	var pixel uint16

	cpu.V[0xF] = 0

	for height := uint16(0); height < n; height++ {
		pixel = uint16(cpu.Memory[cpu.index_register+height])
		// CHIP-8 sprites are always eight pixels wide
		for width := uint16(0); width < 8; width++ {
			index := (x + width + ((y + height) * 64))
			if index <= uint16(len(cpu.Screen)) {

				if (pixel & (0x80 >> width)) != 0 {

					if cpu.Screen[index] == 1 {
						cpu.V[0xF] = 1
					}

					cpu.Screen[index] ^= 1
				}
			}
		}
	}

	s.Opcode = "0xD" + fmt.Sprint(x) + fmt.Sprint(y) + fmt.Sprint(n)
	s.OCName = "DRW V" + fmt.Sprint(x) + " V" + fmt.Sprint(y) + " " + fmt.Sprint(n)
	println(s.toString())

	cpu.PC += 2
}

// ANCHOR - Opcodes 0xE...
// Keyboard inputs
func Opcodes_E(code uint16, cpu *CHIP8) {
	x := (code & 0x0F00) >> 8
	switch code & 0x00FF {
	case 0x009E:
		// Skip next instruction if key with the value of Vx is pressed.
		if cpu.Keyboard[cpu.V[x]] == 1 {
			println("KEY IS PRESSED")
			cpu.PC += 4
		} else {
			cpu.PC += 2
		}
		s.Opcode = "0xE" + fmt.Sprint(x) + "9E"
		s.OCName = "SKP V" + fmt.Sprint(x)
	case 0x00A1:
		// Skip next instruction if key with the value of Vx is not pressed
		if cpu.Keyboard[cpu.V[x]] == 0 {
			cpu.PC += 4
		} else {
			cpu.PC += 2
		}
		s.Opcode = "0xE" + fmt.Sprint(x) + "A1"
		s.OCName = "SKNP V" + fmt.Sprint(x)
	default:
		panic("Invalid opcode :")
	}
	println(s.toString())
}

// ANCHOR - Opcodes 0xF...
func Opcodes_F(code uint16, cpu *CHIP8) {
	x := (code & 0x0F00) >> 8
	switch code & 0x00FF {
	case 0x0007:
		// Set Vx = delay timer value
		cpu.V[x] = cpu.delay_timer
		s.Opcode = "0xF" + fmt.Sprint(x) + "07"
		s.OCName = "LD V" + fmt.Sprint(x) + " DT"
	case 0x000A:
		// Wait for a key press, store the value of the key in Vx
		for index, k := range cpu.Keyboard {
			if k != 0 {
				cpu.V[x] = byte(index)
				cpu.PC += 2
				break
			}
		}

		cpu.Keyboard[cpu.V[x]] = 0
		s.Opcode = "0xF" + fmt.Sprint(x) + "0A"
		s.OCName = "LD V" + fmt.Sprint(x) + " K"
	case 0x0015:
		// Set delay timer = Vx
		cpu.delay_timer = cpu.V[x]
		s.Opcode = "0xF" + fmt.Sprint(x) + "15"
		s.OCName = "LD DT V" + fmt.Sprint(x)
	case 0x0018:
		// Set sound timer = Vx
		cpu.sound_timer = cpu.V[x]
		s.Opcode = "0xF" + fmt.Sprint(x) + "18"
		s.OCName = "LD ST V" + fmt.Sprint(x)
	case 0x001E:
		// Set I = I + Vx
		cpu.index_register += uint16(cpu.V[x])
		s.Opcode = "0xF" + fmt.Sprint(x) + "1E"
		s.OCName = "ADD I V" + fmt.Sprint(x)
	case 0x0029:
		// Set I = location of sprite for digit Vx
		cpu.index_register = uint16(cpu.V[x] * 5)
		s.Opcode = "0xF" + fmt.Sprint(x) + "29"
		s.OCName = "LD F V" + fmt.Sprint(x)
	case 0x0033:
		// Store BCD representation of Vx in memory locations I, I+1, and I+2
		cpu.Memory[cpu.index_register] = cpu.V[x] / 100
		cpu.Memory[cpu.index_register+1] = (cpu.V[x] / 10) % 10
		cpu.Memory[cpu.index_register+2] = (cpu.V[x] % 100) % 10
		s.Opcode = "0xF" + fmt.Sprint(x) + "33"
		s.OCName = "LD B V" + fmt.Sprint(x)
	case 0x0055:
		// Stores V0 to VX in memory starting at address I. I is then set to I + x + 1.
		for i := uint16(0); i <= x; i++ {
			cpu.Memory[cpu.index_register] = cpu.V[i]
			// The save and load opcodes (Fx55 and Fx65) increment the index register.
			cpu.index_register++
		}
		s.Opcode = "0xF" + fmt.Sprint(x) + "55"
		s.OCName = "LD [" + fmt.Sprint(cpu.index_register) + "] V" + fmt.Sprint(x)
	case 0x0065:
		// Fills V0 to VX with values from memory starting at address I. I is then set to I + x + 1
		for i := uint16(0); i <= x; i++ {
			cpu.V[i] = cpu.Memory[cpu.index_register]
			// The save and load opcodes (Fx55 and Fx65) increment the index register.
			cpu.index_register++
		}
		s.Opcode = "0xF" + fmt.Sprint(x) + "65"
		s.OCName = "LD V" + fmt.Sprint(x) + "[" + fmt.Sprint(cpu.index_register) + "]"
	default:
		panic("Invalid opcode :")
	}
	println(s.toString())
	cpu.PC += 2
}
