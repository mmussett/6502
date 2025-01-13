package nmos6502

/*
ADC - Add Memory to Accumulator with Carry

A + M + C -> A, C
N	Z	C	I	D	V
+	+	+	-	-	+
addressing		assembler		opc	bytes	cycles
immediate		ADC #oper		69	2		2
zeropage		ADC oper		65	2		3
zeropage,X		ADC oper,X		75	2		4
absolute		ADC oper		6D	3		4
absolute,X		ADC oper,X		7D	3		4*
absolute,Y		ADC oper,Y		79	3		4*
(indirect,X)	ADC (oper,X)	61	2		6
(indirect),Y	ADC (oper),Y	71	2		5*
*/

func (cpu *CPU) opcode0x69() (byte, bool) { // ADC Immediate
	value := cpu.fetch()
	cpu.ADC(value)
	return 2, false // ADC immediate takes 2 cycles
}

func (cpu *CPU) opcode0x65() (byte, bool) { // ADC Zero Page
	address := uint16(cpu.fetch())
	value := cpu.Memory[address]
	cpu.ADC(value)
	return 3, false // ADC zero page takes 3 cycles
}

func (cpu *CPU) opcode0x75() (byte, bool) { // ADC Zero Page, X
	address := uint16(cpu.fetch() + cpu.X)
	value := cpu.Memory[address]
	cpu.ADC(value)
	return 4, false // ADC zero page, X takes 4 cycles
}

func (cpu *CPU) opcode0x6D() (byte, bool) { // ADC Absolute
	low := cpu.fetch()
	high := cpu.fetch()
	address := uint16(high)<<8 | uint16(low)
	value := cpu.Memory[address]
	cpu.ADC(value)
	return 4, false // ADC absolute takes 4 cycles
}

func (cpu *CPU) opcode0x7D() (byte, bool) { // ADC Absolute, X
	low := cpu.fetch()
	high := cpu.fetch()
	address := uint16(high)<<8 | uint16(low) + uint16(cpu.X)
	value := cpu.Memory[address]
	cpu.ADC(value)
	// Additional cycle if page boundary is crossed
	if (address & 0xFF00) != ((address - uint16(cpu.X)) & 0xFF00) {
		return 5, false
	}
	return 4, false
}

func (cpu *CPU) opcode0x79() (byte, bool) { // ADC Absolute, Y
	low := cpu.fetch()
	high := cpu.fetch()
	address := uint16(high)<<8 | uint16(low) + uint16(cpu.Y)
	value := cpu.Memory[address]
	cpu.ADC(value)
	// Additional cycle if page boundary is crossed
	if (address & 0xFF00) != ((address - uint16(cpu.Y)) & 0xFF00) {
		return 5, false
	}
	return 4, false
}

func (cpu *CPU) opcode0x61() (byte, bool) { // ADC (Indirect, X)
	base := uint16(cpu.fetch() + cpu.X)
	low := cpu.Memory[base]
	high := cpu.Memory[base+1]
	address := uint16(high)<<8 | uint16(low)
	value := cpu.Memory[address]
	cpu.ADC(value)
	return 6, false // ADC (indirect, X) takes 6 cycles
}

func (cpu *CPU) opcode0x71() (byte, bool) { // ADC (Indirect), Y
	base := uint16(cpu.fetch())
	low := cpu.Memory[base]
	high := cpu.Memory[base+1]
	address := uint16(high)<<8 | uint16(low) + uint16(cpu.Y)
	value := cpu.Memory[address]
	cpu.ADC(value)
	// Additional cycle if page boundary is crossed
	if (address & 0xFF00) != ((address - uint16(cpu.Y)) & 0xFF00) {
		return 6, false
	}
	return 5, false
}

func (cpu *CPU) ADC(value byte) {
	sum := uint16(cpu.A) + uint16(value) + uint16(cpu.Status&Carry)
	carry := sum > 0xFF

	if carry {
		cpu.setCarryFlag()
	} else {
		cpu.clearCarryFlag()
	}

	result := byte(sum)
	cpu.updateZeroAndNegativeFlags(result)

	// Overflow occurs if the sign bit is incorrect
	if (cpu.A^result)&(value^result)&0x80 != 0 {
		cpu.setOverflowFlag()
	} else {
		cpu.clearOverflowFlag()
	}

	cpu.A = result
}

/*
SBC - Subtract Memory from Accumulator with Borrow

A - M - C̅ -> A
N	Z	C	I	D	V
+	+	+	-	-	+
addressing		assembler		opc	bytes	cycles
immediate		SBC #oper		E9	2		2
zeropage		SBC oper		E5	2		3
zeropage,X		SBC oper,X		F5	2		4
absolute		SBC oper		ED	3		4
absolute,X		SBC oper,X		FD	3		4*
absolute,Y		SBC oper,Y		F9	3		4*
(indirect,X)	SBC (oper,X)	E1	2		6
(indirect),Y	SBC (oper),Y	F1	2		5*
*/

func (cpu *CPU) opcode0xE9() (byte, bool) { // SBC Immediate
	value := cpu.fetch()
	cpu.SBC(value)
	return 2, false // SBC immediate takes 2 cycles
}

func (cpu *CPU) opcode0xE5() (byte, bool) { // SBC Zero Page
	address := uint16(cpu.fetch())
	value := cpu.Memory[address]
	cpu.SBC(value)
	return 3, false // SBC zero page takes 3 cycles
}

func (cpu *CPU) opcode0xF5() (byte, bool) { // SBC Zero Page, X
	address := uint16(cpu.fetch() + cpu.X)
	value := cpu.Memory[address]
	cpu.SBC(value)
	return 4, false // SBC zero page, X takes 4 cycles
}

func (cpu *CPU) opcode0xED() (byte, bool) { // SBC Absolute
	low := cpu.fetch()
	high := cpu.fetch()
	address := uint16(high)<<8 | uint16(low)
	value := cpu.Memory[address]
	cpu.SBC(value)
	return 4, false // SBC absolute takes 4 cycles
}

func (cpu *CPU) opcode0xFD() (byte, bool) { // SBC Absolute, X
	low := cpu.fetch()
	high := cpu.fetch()
	address := uint16(high)<<8 | uint16(low) + uint16(cpu.X)
	value := cpu.Memory[address]
	cpu.SBC(value)
	// Additional cycle if page boundary is crossed
	if (address & 0xFF00) != ((address - uint16(cpu.X)) & 0xFF00) {
		return 5, false
	}
	return 4, false
}

func (cpu *CPU) opcode0xF9() (byte, bool) { // SBC Absolute, Y
	low := cpu.fetch()
	high := cpu.fetch()
	address := uint16(high)<<8 | uint16(low) + uint16(cpu.Y)
	value := cpu.Memory[address]
	cpu.SBC(value)
	// Additional cycle if page boundary is crossed
	if (address & 0xFF00) != ((address - uint16(cpu.Y)) & 0xFF00) {
		return 5, false
	}
	return 4, false
}

func (cpu *CPU) opcode0xE1() (byte, bool) { // SBC (Indirect, X)
	base := uint16(cpu.fetch() + cpu.X)
	low := cpu.Memory[base]
	high := cpu.Memory[base+1]
	address := uint16(high)<<8 | uint16(low)
	value := cpu.Memory[address]
	cpu.SBC(value)
	return 6, false // SBC (indirect, X) takes 6 cycles
}

func (cpu *CPU) opcode0xF1() (byte, bool) { // SBC (Indirect), Y
	base := uint16(cpu.fetch())
	low := cpu.Memory[base]
	high := cpu.Memory[base+1]
	address := uint16(high)<<8 | uint16(low) + uint16(cpu.Y)
	value := cpu.Memory[address]
	cpu.SBC(value)
	// Additional cycle if page boundary is crossed
	if (address & 0xFF00) != ((address - uint16(cpu.Y)) & 0xFF00) {
		return 6, false
	}
	return 5, false
}

func (cpu *CPU) SBC(value byte) {
	value = ^value // Invert the value for subtraction
	sum := uint16(cpu.A) + uint16(value) + uint16(cpu.Status&Carry)
	carry := sum > 0xFF

	if carry {
		cpu.setCarryFlag()
	} else {
		cpu.clearCarryFlag()
	}

	result := byte(sum)
	cpu.updateZeroAndNegativeFlags(result)

	// Overflow occurs if the sign bit is incorrect
	if (cpu.A^result)&(^value^result)&0x80 != 0 {
		cpu.setOverflowFlag()
	} else {
		cpu.clearOverflowFlag()
	}

	cpu.A = result
}
