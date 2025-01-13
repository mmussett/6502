package nmos6502

/*
BRK Force Break

BRK initiates a software interrupt similar to a hardware
interrupt (IRQ). The return address pushed to the stack is
PC+2, providing an extra byte of spacing for a break mark
(identifying a reason for the break.)
The status register will be pushed to the stack with the break
flag set to 1. However, when retrieved during RTI or by a PLP
instruction, the break flag will be ignored.
The interrupt disable flag is not set automatically.

interrupt,
push PC+2, push SR
N	Z	C	I	D	V
-	-	-	1	-	-
addressing	assembler	opc	bytes	cycles
implied		BRK			00	1		7
*/
func (cpu *CPU) opcode0x00() (byte, bool) { // BRK - Force Interrupt
	// Push PC+2 onto the stack
	pc := cpu.PC + 1              // PC+2 is already handled by the fetch-execute cycle
	cpu.pushByte(byte(pc >> 8))   // Push high byte
	cpu.pushByte(byte(pc & 0xFF)) // Push low byte

	// Push the status register onto the stack with the break flag set
	cpu.setBreakCommandFlag()

	cpu.pushByte(cpu.Status)

	// Set the interrupt disable flag
	cpu.setIntDisableFlag()

	// Load the interrupt vector from $FFFE/$FFFF
	low := cpu.Memory[0xFFFE]
	high := cpu.Memory[0xFFFF]
	cpu.PC = uint16(high)<<8 | uint16(low)

	return 7, false // BRK takes 7 cycles
}

/*
RTI - Return from Interrupt

The status register is pulled with the break flag
and bit 5 ignored. Then PC is pulled from the stack.

pull SR, pull PC
N	Z	C	I	D	V
from stack
addressing	assembler	opc	bytes	cycles
implied		RTI			40	1		6

*/

func (cpu *CPU) opcode0x40() (byte, bool) {
	// Pull the status register from the stack and ignore the break flag
	cpu.Status = cpu.popByte()
	cpu.clearBreakCommandFlag()

	// Pull the program counter from the stack
	low := cpu.popByte()
	high := cpu.popByte()
	cpu.PC = uint16(high)<<8 | uint16(low)

	return 6, false

}
