package nmos6502

const (
	StackBase = 0x100 // Base address for stack operations
)

/*
Stack Instructions

These instructions transfer the accumulator or status register (flags) to and from the stack. The processor stack is a last-in-first-out (LIFO) stack of 256 bytes length, implemented at addresses $0100 - $01FF. The stack grows down as new values are pushed onto it with the current insertion point maintained in the stack pointer register.
(When a byte is pushed onto the stack, it will be stored in the address indicated by the value currently in the stack pointer, which will be then decremented by 1. Conversely, when a value is pulled from the stack, the stack pointer is incremented. The stack pointer is accessible by the TSX and TXS instructions.)

PHA push accumulator
PHP push processor status register (with break flag set)
PLA pull accumulator
PLP pull processor status register
*/
func (cpu *CPU) pushByte(b byte) {
	cpu.SP--
	if cpu.SP < 0 {
		cpu.SP = 0xFF
	}
	cpu.Memory[StackBase+uint16(cpu.SP)] = b
}

func (cpu *CPU) popByte() byte {
	b := cpu.Memory[StackBase+uint16(cpu.SP)]
	cpu.SP++
	if cpu.SP > 0xFF {
		cpu.SP = 0x00
	}
	return b
}

/*
PHA - Push Accumulator on Stack

push A
N	Z	C	I	D	V
-	-	-	-	-	-

addressing	assembler	opc	bytes	cycles
implied		PHA			48	1		3
*/
func (cpu *CPU) opcode0x48() (byte, bool) { // PHA - Push Accumulator
	cpu.pushByte(cpu.A)
	return 3, false
}

/*
PHP - Push Processor Status on Stack

The status register will be pushed with the break
flag and bit 5 set to 1.

push SR
N	Z	C	I	D	V
-	-	-	-	-	-
addressing	assembler	opc	bytes	cycles
implied		PHP			08	1		3
*/
func (cpu *CPU) opcode0x08() (byte, bool) { // PHP - Push Processor Status Register
	// Set the break flag (bit 4) before pushing
	statusWithBreak := cpu.Status | 0x10
	cpu.pushByte(statusWithBreak)
	return 3, false
}

/*
PLA - Pull Accumulator from Stack

pull A
N	Z	C	I	D	V
+	+	-	-	-	-
addressing	assembler	opc	bytes	cycles
implied		PLA			68	1		4
*/
func (cpu *CPU) opcode0x68() (byte, bool) { // PLA - Pull Accumulator
	cpu.A = cpu.popByte()
	cpu.updateZeroAndNegativeFlags(cpu.A)
	return 4, false
}

/*
PLP - Pull Processor Status from Stack

The status register will be pulled with the break
flag and bit 5 ignored.

pull SR
N	Z	C	I	D	V
from stack
addressing	assembler	opc	bytes	cycles
implied		PLP			28	1		4
*/
func (cpu *CPU) opcode0x28() (byte, bool) { // PLP - Pull Processor Status Register
	cpu.Status = cpu.popByte()
	// Ensure the break flag is not set in the status register after pulling
	cpu.Status &^= BreakCommand
	return 4, false
}
