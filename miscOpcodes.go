package nmos6502

/*
BIT - Test Bits in Memory with Accumulator

bits 7 and 6 of operand are transfered to bit 7 and 6 of SR (N,V);
the zero-flag is set according to the result of the operand AND
the accumulator (set, if the result is zero, unset otherwise).
This allows a quick check of a few bits at once without affecting
any of the registers, other than the status register (SR).

A AND M -> Z, M7 -> N, M6 -> V
N	Z	C	I	D	V
M7	+	-	-	-	M6
addressing	assembler	opc	bytes	cycles
zeropage	BIT oper	24	2		3
absolute	BIT oper	2C	3		4
*/

func (cpu *CPU) opcode0x24() (byte, bool) { // BIT Zero Page
	address := uint16(cpu.fetch())
	value := cpu.Memory[address]
	cpu.BIT(value)
	return 3, false // BIT zero page takes 3 cycles
}

func (cpu *CPU) opcode0x2C() (byte, bool) { // BIT Absolute
	low := cpu.fetch()
	high := cpu.fetch()
	address := uint16(high)<<8 | uint16(low)
	value := cpu.Memory[address]
	cpu.BIT(value)
	return 4, false // BIT absolute takes 4 cycles
}

func (cpu *CPU) BIT(value byte) {
	// Set the zero flag based on the result of A AND M
	if cpu.A&value == 0 {
		cpu.setZeroFlag()
	} else {
		cpu.clearZeroFlag()
	}

	// Set the negative flag to bit 7 of the memory operand
	if value&0x80 != 0 {
		cpu.setNegativeFlag()
	} else {
		cpu.clearNegativeFlag()
	}

	// Set the overflow flag to bit 6 of the memory operand
	if value&0x40 != 0 {
		cpu.setOverflowFlag()
	} else {
		cpu.clearOverflowFlag()
	}
}
