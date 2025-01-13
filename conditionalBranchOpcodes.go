package nmos6502

/*
BCC - Branch on Carry Clear

branch on C = 0
N	Z	C	I	D	V
-	-	-	-	-	-

addressing	assembler	opc	bytes	cycles
relative	BCC oper	90	2		2**
*/

func (cpu *CPU) opcode0x90() (byte, bool) { // BCC Relative
	offset := int8(cpu.fetch()) // Fetch the relative offset as a signed byte

	if cpu.Status&Carry == 0 { // Check if the carry flag is clear
		oldPC := cpu.PC
		cpu.PC += uint16(offset) // Add the offset to the program counter

		// Additional cycle if branch occurs to a new page
		if (oldPC & 0xFF00) != (cpu.PC & 0xFF00) {
			return 3, false // BCC takes 3 cycles if a page boundary is crossed
		}
		return 2, false // BCC takes 2 cycles if no page boundary is crossed
	}

	return 2, false // BCC takes 2 cycles if no branch is taken
}

/*
BCS - Branch on Carry Set

branch on C = 1
N	Z	C	I	D	V
-	-	-	-	-	-

addressing	assembler	opc	bytes	cycles
relative	BCS oper	B0	2	2**

*/

func (cpu *CPU) opcode0xB0() (byte, bool) { // BCS Relative
	offset := int8(cpu.fetch()) // Fetch the relative offset as a signed byte

	if cpu.Status&Carry != 0 { // Check if the carry flag is set
		oldPC := cpu.PC
		cpu.PC += uint16(offset) // Add the offset to the program counter

		// Additional cycle if branch occurs to a new page
		if (oldPC & 0xFF00) != (cpu.PC & 0xFF00) {
			return 3, false // BCS takes 3 cycles if a page boundary is crossed
		}
		return 2, false // BCS takes 2 cycles if no page boundary is crossed
	}

	return 2, false // BCS takes 2 cycles if no branch is taken
}

/*
BEQ - Branch on Result Zero

branch on Z = 1
N	Z	C	I	D	V
-	-	-	-	-	-

addressing	assembler	opc	bytes	cycles
relative	BEQ oper	F0	2	2**
*/

func (cpu *CPU) opcode0xF0() (byte, bool) { // BEQ Relative
	offset := int8(cpu.fetch()) // Fetch the relative offset as a signed byte

	if cpu.Status&Zero != 0 { // Check if the zero flag is set
		oldPC := cpu.PC
		cpu.PC += uint16(offset) // Add the offset to the program counter

		// Additional cycle if branch occurs to a new page
		if (oldPC & 0xFF00) != (cpu.PC & 0xFF00) {
			return 3, false // BEQ takes 3 cycles if a page boundary is crossed
		}
		return 2, false // BEQ takes 2 cycles if no page boundary is crossed
	}

	return 2, false // BEQ takes 2 cycles if no branch is taken
}

/*
BMI - Branch on Result Minus

branch on N = 1
N	Z	C	I	D	V
-	-	-	-	-	-
addressing	assembler	opc	bytes	cycles
relative	BMI oper	30	2		2**
*/

func (cpu *CPU) opcode0x30() (byte, bool) { // BMI Relative
	offset := int8(cpu.fetch()) // Fetch the relative offset as a signed byte

	if cpu.Status&Negative != 0 { // Check if the negative flag is set
		oldPC := cpu.PC
		cpu.PC += uint16(offset) // Add the offset to the program counter

		// Additional cycle if branch occurs to a new page
		if (oldPC & 0xFF00) != (cpu.PC & 0xFF00) {
			return 3, false // BMI takes 3 cycles if a page boundary is crossed
		}
		return 2, false // BMI takes 2 cycles if no page boundary is crossed
	}

	return 2, false // BMI takes 2 cycles if no branch is taken
}

/*
BNE - Branch on Result not Zero

branch on Z = 0
N	Z	C	I	D	V
-	-	-	-	-	-
addressing	assembler	opc	bytes	cycles
relative	BNE oper	D0	2		2**
*/
func (cpu *CPU) opcode0xD0() (byte, bool) { // BNE Relative
	offset := int8(cpu.fetch()) // Fetch the relative offset as a signed byte

	if cpu.Status&Zero == 0 { // Check if the zero flag is clear
		oldPC := cpu.PC
		cpu.PC += uint16(offset) // Add the offset to the program counter

		// Additional cycle if branch occurs to a new page
		if (oldPC & 0xFF00) != (cpu.PC & 0xFF00) {
			return 3, false // BNE takes 3 cycles if a page boundary is crossed
		}
		return 2, false // BNE takes 2 cycles if no page boundary is crossed
	}

	return 2, false // BNE takes 2 cycles if no branch is taken
}

/*
BPL - Branch on Result Plus

branch on N = 0
N	Z	C	I	D	V
-	-	-	-	-	-
addressing	assembler	opc	bytes	cycles
relative	BPL oper	10	2		2**
*/
func (cpu *CPU) opcode0x10() (byte, bool) { // BPL Relative
	offset := int8(cpu.fetch()) // Fetch the relative offset as a signed byte

	if cpu.Status&Negative == 0 { // Check if the negative flag is clear
		oldPC := cpu.PC
		cpu.PC += uint16(offset) // Add the offset to the program counter

		// Additional cycle if branch occurs to a new page
		if (oldPC & 0xFF00) != (cpu.PC & 0xFF00) {
			return 3, false // BPL takes 3 cycles if a page boundary is crossed
		}
		return 2, false // BPL takes 2 cycles if no page boundary is crossed
	}

	return 2, false // BPL takes 2 cycles if no branch is taken
}

/*
BVC - Branch on Overflow Clear

branch on V = 0
N	Z	C	I	D	V
-	-	-	-	-	-
addressing	assembler	opc	bytes	cycles
relative	BVC oper	50	2		2**
*/
func (cpu *CPU) opcode0x50() (byte, bool) { // BVC Relative
	offset := int8(cpu.fetch()) // Fetch the relative offset as a signed byte

	if cpu.Status&Overflow == 0 { // Check if the overflow flag is clear
		oldPC := cpu.PC
		cpu.PC += uint16(offset) // Add the offset to the program counter

		// Additional cycle if branch occurs to a new page
		if (oldPC & 0xFF00) != (cpu.PC & 0xFF00) {
			return 3, false // BVC takes 3 cycles if a page boundary is crossed
		}
		return 2, false // BVC takes 2 cycles if no page boundary is crossed
	}

	return 2, false // BVC takes 2 cycles if no branch is taken
}

/*
BVS - Branch on Overflow Set

branch on V = 1
N	Z	C	I	D	V
-	-	-	-	-	-
addressing	assembler	opc	bytes	cycles
relative	BVS oper	70	2		2**
*/

func (cpu *CPU) opcode0x70() (byte, bool) { // BVS Relative
	offset := int8(cpu.fetch()) // Fetch the relative offset as a signed byte

	if cpu.Status&Overflow != 0 { // Check if the overflow flag is clear
		oldPC := cpu.PC
		cpu.PC += uint16(offset) // Add the offset to the program counter

		// Additional cycle if branch occurs to a new page
		if (oldPC & 0xFF00) != (cpu.PC & 0xFF00) {
			return 3, false // BVS takes 3 cycles if a page boundary is crossed
		}
		return 2, false // BVS takes 2 cycles if no page boundary is crossed
	}

	return 2, false // BVS takes 2 cycles if no branch is taken
}
