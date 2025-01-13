package nmos6502

/*
CMP
Compare Memory with Accumulator

A - M
N	Z	C	I	D	V
+	+	+	-	-	-
addressing		assembler		opc	bytes	cycles
immediate		CMP #oper		C9	2	2
zeropage		CMP oper		C5	2	3
zeropage,X		CMP oper,X		D5	2	4
absolute		CMP oper		CD	3	4
absolute,X		CMP oper,X		DD	3	4*
absolute,Y		CMP oper,Y		D9	3	4*
(indirect,X)	CMP (oper,X)	C1	2	6
(indirect),Y	CMP (oper),Y	D1	2	5*
*/

func (cpu *CPU) opcode0xC9() (byte, bool) { // CMP Immediate
	value := cpu.fetch() // Fetch the immediate value

	result := cpu.A - value

	// Set or clear the carry flag
	if cpu.A >= value {
		cpu.setCarryFlag()
	} else {
		cpu.clearCarryFlag()
	}

	// Update zero and negative flags based on the result
	cpu.updateZeroAndNegativeFlags(result)

	return 2, false // CMP immediate takes 2 cycles
}

func (cpu *CPU) opcode0xC5() (byte, bool) { // CMP Zero Page
	address := uint16(cpu.fetch())
	value := cpu.Memory[address] // Read from memory at the fetched address

	result := cpu.A - value
	// Set or clear the carry flag
	if cpu.A >= value {
		cpu.setCarryFlag()
	} else {
		cpu.clearCarryFlag()
	}

	// Update zero and negative flags based on the result
	cpu.updateZeroAndNegativeFlags(result)
	return 3, false // CMP zeropage takes 3 cycles

}

func (cpu *CPU) opcode0xD5() (byte, bool) { // CMP Zero Page Indexed by X
	address := uint16(cpu.fetch()) + uint16(cpu.X)
	value := cpu.Memory[address]

	result := cpu.A - value

	// Set or clear the carry flag
	if cpu.A >= value {
		cpu.setCarryFlag()
	} else {
		cpu.clearCarryFlag()
	}

	// Update zero and negative flags based on the result
	cpu.updateZeroAndNegativeFlags(result)
	return 4, false // CMP zeropage indexed by X takes 4 cycles

}

func (cpu *CPU) opcode0xCD() (byte, bool) { // CMP Absolute

	low := cpu.fetch()
	high := cpu.fetch()
	address := uint16(high)<<8 | uint16(low)

	value := cpu.Memory[address]

	result := cpu.A - value

	// Set or clear the carry flag
	if cpu.A >= value {
		cpu.setCarryFlag()
	} else {
		cpu.clearCarryFlag()
	}

	// Update zero and negative flags based on the result
	cpu.updateZeroAndNegativeFlags(result)
	return 4, false // CMP absolute takes 4 cycles
}

func (cpu *CPU) opcode0xDD() (byte, bool) { // CMP Absolute Indexed by X

	low := cpu.fetch()
	high := cpu.fetch()
	effectiveAddress := uint16(high)<<8 | uint16(low)

	address := effectiveAddress + uint16(cpu.X)
	value := cpu.Memory[address]

	result := cpu.A - value
	cpu.updateZeroAndNegativeFlags(result)

	// Set or clear the carry flag
	if cpu.A >= value {
		cpu.setCarryFlag()
	} else {
		cpu.clearCarryFlag()
	}

	// CMP absolute indexed by X takes 4 cycles if page boundary is crossed; otherwise it takes 3 cycles.
	if (address & 0xFF00) != (effectiveAddress & 0xFF00) {
		return 4, false
	} else {
		return 3, false
	}

}

func (cpu *CPU) opcode0xD9() (byte, bool) { // CMP Absolute Indexed by Y
	low := cpu.fetch()
	high := cpu.fetch()
	effectiveAddress := uint16(high)<<8 | uint16(low)
	address := effectiveAddress + uint16(cpu.Y)
	value := cpu.Memory[address]

	result := cpu.A - value
	cpu.updateZeroAndNegativeFlags(result)

	// Set or clear the carry flag
	if cpu.A >= value {
		cpu.setCarryFlag()
	} else {
		cpu.clearCarryFlag()
	}

	// CMP absolute indexed by Y takes 4 cycles if page boundary is crossed; otherwise it takes 3 cycles.
	if (address & 0xFF00) != (effectiveAddress & 0xFF00) {
		return 4, false
	} else {
		return 3, false
	}
}

func (cpu *CPU) opcode0xC1() (byte, bool) { // CMP Indirect Indexed by X
	indexedAddress := uint16(cpu.fetch()) + uint16(cpu.X)
	low := cpu.Memory[indexedAddress]
	high := cpu.Memory[indexedAddress+1]
	address := uint16(high)<<8 | uint16(low)
	value := cpu.Memory[address]

	result := cpu.A - value

	// Set or clear the carry flag
	if cpu.A >= value {
		cpu.setCarryFlag()
	} else {
		cpu.clearCarryFlag()
	}

	// Update zero and negative flags based on the result
	cpu.updateZeroAndNegativeFlags(result)
	return 6, false // CMP indirect indexed by X takes 6 cycles
}

func (cpu *CPU) opcode0xD1() (byte, bool) { // CMP Indirect Indexed by Y
	indexedAddress := uint16(cpu.fetch())
	low := cpu.Memory[indexedAddress]
	high := cpu.Memory[indexedAddress+1]
	effectiveAddress := uint16(high)<<8 | uint16(low)
	address := effectiveAddress + uint16(cpu.Y)
	value := cpu.Memory[address]

	result := cpu.A - value
	cpu.updateZeroAndNegativeFlags(result)

	// Set or clear the carry flag
	if cpu.A >= value {
		cpu.setCarryFlag()
	} else {
		cpu.clearCarryFlag()
	}

	// CMP Indirect Indexed by Y takes 5 cycles if page boundary is crossed; otherwise it takes 4 cycles.
	if (address & 0xFF00) != (effectiveAddress & 0xFF00) {
		return 5, false
	} else {
		return 4, false
	}
}

/*
CPX
Compare Memory and Index X

X - M
N	Z	C	I	D	V
+	+	+	-	-	-
addressing	assembler	opc	bytes	cycles
immediate	CPX #oper	E0	2		2
zeropage	CPX oper	E4	2		3
absolute	CPX oper	EC	3		4
*/
func (cpu *CPU) opcode0xE0() (byte, bool) { // CPX Immediate
	value := cpu.fetch() // Fetch the immediate value
	result := cpu.X - value

	// Set or clear the carry flag
	if cpu.X >= value {
		cpu.setCarryFlag()
	} else {
		cpu.clearCarryFlag()
	}

	// Update zero and negative flags based on the result
	cpu.updateZeroAndNegativeFlags(result)

	return 2, false // CPX immediate takes 2 cycles
}

func (cpu *CPU) opcode0xE4() (byte, bool) { // CPX Zero Page
	address := uint16(cpu.fetch())
	value := cpu.Memory[address]
	result := cpu.X - value

	// Set or clear the carry flag
	if cpu.X >= value {
		cpu.setCarryFlag()
	} else {
		cpu.clearCarryFlag()
	}

	// Update zero and negative flags based on the result
	cpu.updateZeroAndNegativeFlags(result)

	return 3, false // CPX zero page takes 3 cycles
}

func (cpu *CPU) opcode0xEC() (byte, bool) { // CPX Absolute
	low := cpu.fetch()
	high := cpu.fetch()
	address := uint16(high)<<8 | uint16(low)
	value := cpu.Memory[address]
	result := cpu.X - value

	// Set or clear the carry flag
	if cpu.X >= value {
		cpu.setCarryFlag()
	} else {
		cpu.clearCarryFlag()
	}

	// Update zero and negative flags based on the result
	cpu.updateZeroAndNegativeFlags(result)

	return 4, false // CPX absolute takes 4 cycles
}

/*
CPY - Compare Memory and Index Y

Y - M
N	Z	C	I	D	V
+	+	+	-	-	-
addressing	assembler	opc	bytes	cycles
immediate	CPY #oper	C0	2		2
zeropage	CPY oper	C4	2		3
absolute	CPY oper	CC	3		4
*/

func (cpu *CPU) opcode0xC0() (byte, bool) { // CPY Immediate
	value := cpu.fetch() // Fetch the immediate value
	result := cpu.Y - value

	// Set or clear the carry flag
	if cpu.Y >= value {
		cpu.setCarryFlag()
	} else {
		cpu.clearCarryFlag()
	}

	// Update zero and negative flags based on the result
	cpu.updateZeroAndNegativeFlags(result)

	return 2, false // CPY immediate takes 2 cycles
}

func (cpu *CPU) opcode0xC4() (byte, bool) { // CPY Zero Page
	address := uint16(cpu.fetch())
	value := cpu.Memory[address]
	result := cpu.Y - value

	// Set or clear the carry flag
	if cpu.Y >= value {
		cpu.setCarryFlag()
	} else {
		cpu.clearCarryFlag()
	}

	// Update zero and negative flags based on the result
	cpu.updateZeroAndNegativeFlags(result)

	return 3, false // CPY zero page takes 3 cycles
}

func (cpu *CPU) opcode0xCC() (byte, bool) { // CPY Absolute
	low := cpu.fetch()
	high := cpu.fetch()
	address := uint16(high)<<8 | uint16(low)
	value := cpu.Memory[address]
	result := cpu.Y - value

	// Set or clear the carry flag
	if cpu.Y >= value {
		cpu.setCarryFlag()
	} else {
		cpu.clearCarryFlag()
	}

	// Update zero and negative flags based on the result
	cpu.updateZeroAndNegativeFlags(result)

	return 4, false // CPY absolute takes 4 cycles
}
