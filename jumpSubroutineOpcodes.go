package nmos6502

/*
JMP - Jump to New Location

operand 1st byte -> PCL
operand 2nd byte -> PCH
N	Z	C	I	D	V
-	-	-	-	-	-
addressing	assembler	opc	bytes	cycles
absolute	JMP oper	4C	3	3
indirect	JMP (oper)	6C	3	5

*/

func (cpu *CPU) opcode0x4C() (byte, bool) { // JMP Absolute
	low := cpu.fetch()                       // Fetch the low byte of the address
	high := cpu.fetch()                      // Fetch the high byte of the address
	address := uint16(high)<<8 | uint16(low) // Combine to form the full address
	cpu.PC = address                         // Set the program counter to the new address
	return 3, false                          // JMP absolute takes 3 cycles
}

func (cpu *CPU) opcode0x6C() (byte, bool) { // JMP Indirect
	low := cpu.fetch()                       // Fetch the low byte of the address
	high := cpu.fetch()                      // Fetch the high byte of the address
	address := uint16(high)<<8 | uint16(low) // Combine to form the full address

	// Handle the 6502 bug where the high byte of the indirect address wraps around
	if low == 0xFF {
		// Simulate the page boundary bug
		targetLow := cpu.Memory[address]
		targetHigh := cpu.Memory[address&0xFF00] // Wrap around to the start of the page
		cpu.PC = uint16(targetHigh)<<8 | uint16(targetLow)
	} else {
		// Normal indirect addressing
		targetLow := cpu.Memory[address]
		targetHigh := cpu.Memory[address+1]
		cpu.PC = uint16(targetHigh)<<8 | uint16(targetLow)
	}

	return 5, false // JMP indirect takes 5 cycles
}

/*
JSR - Jump to New Location Saving Return Address

push (PC+2),
operand 1st byte -> PCL
operand 2nd byte -> PCH
N	Z	C	I	D	V
-	-	-	-	-	-
addressing	assembler	opc	bytes	cycles
absolute	JSR oper	20	3		6
*/

func (cpu *CPU) opcode0x20() (byte, bool) { // JSR Absolute
	low := cpu.fetch()
	high := cpu.fetch()
	address := uint16(high)<<8 | uint16(low)

	// Push return address onto stack
	stackAddress := cpu.SP + 1
	cpu.Memory[stackAddress] = uint8((cpu.PC >> 8) & 0xFF)
	stackAddress--
	cpu.Memory[stackAddress] = uint8(cpu.PC & 0xFF)

	cpu.PC = address
	return 6, false // JSR absolute takes 6 cycles
}

/*
RTS - Return from Subroutine

pull PC, PC+1 -> PC
N	Z	C	I	D	V
-	-	-	-	-	-
addressing	assembler	opc	bytes	cycles
implied	RTS	60	1	6

*/

func (cpu *CPU) opcode0x60() (byte, bool) { // RTS Implied
	// Pull the program counter from the stack
	low := cpu.popByte()
	high := cpu.popByte()
	cpu.PC = uint16(high)<<8 | uint16(low)

	// Increment the program counter by one
	cpu.PC++

	return 6, false // RTS takes 6 cycles
}
