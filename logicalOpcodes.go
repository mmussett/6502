package nmos6502

/**
ORA - OR Memory with Accumulator

A OR M -> A
N	Z	C	I	D	V
+	+	-	-	-	-
addressing		assembler		opc	bytes	cycles
immediate		ORA #oper		09	2		2
zeropage		ORA oper		05	2		3
zeropage,X		ORA oper,X		15	2		4
absolute		ORA oper		0D	3		4
absolute,X		ORA oper,X		1D	3		4*
absolute,Y		ORA oper,Y		19	3		4*
(indirect,X)	ORA (oper,X)	01	2	6
(indirect),Y	ORA (oper),Y	11	2	5*
*/

func (cpu *CPU) opcode0x09() (byte, bool) {
	value := cpu.fetch()                  // Fetch the immediate value
	cpu.A |= value                        // OR the value with the accumulator
	cpu.updateZeroAndNegativeFlags(cpu.A) // Update the zero and negative flags
	return 2, false                       // ORA immediate takes 2 cycles
}

func (cpu *CPU) opcode0x05() (byte, bool) {
	address := uint16(cpu.fetch())        // Fetch the zero page address
	value := cpu.Memory[address]          // Get the value from memory
	cpu.A |= value                        // OR the value with the accumulator
	cpu.updateZeroAndNegativeFlags(cpu.A) // Update the zero and negative flags
	return 3, false                       // ORA zero page takes 3 cycles
}

func (cpu *CPU) opcode0x15() (byte, bool) {
	baseAddress := cpu.fetch()                      // Fetch the zero page base address
	address := uint16((baseAddress + cpu.X) & 0xFF) // Calculate the effective address with zero page wrap-around
	value := cpu.Memory[address]                    // Get the value from memory
	cpu.A |= value                                  // OR the value with the accumulator
	cpu.updateZeroAndNegativeFlags(cpu.A)           // Update the zero and negative flags
	return 4, false                                 // ORA zero page,X takes 4 cycles
}

func (cpu *CPU) opcode0x0D() (byte, bool) {
	low := cpu.fetch()                       // Fetch the low byte of the address
	high := cpu.fetch()                      // Fetch the high byte of the address
	address := uint16(high)<<8 | uint16(low) // Combine to form the full address
	value := cpu.Memory[address]             // Get the value from memory
	cpu.A |= value                           // OR the value with the accumulator
	cpu.updateZeroAndNegativeFlags(cpu.A)    // Update the zero and negative flags
	return 4, false                          // ORA absolute takes 4 cycles
}

func (cpu *CPU) opcode0x1D() (byte, bool) {
	low := cpu.fetch()                          // Fetch the low byte of the address
	high := cpu.fetch()                         // Fetch the high byte of the address
	address := uint16(high)<<8 | uint16(low)    // Combine to form the full address
	effectiveAddress := address + uint16(cpu.X) // Add the X register to the address
	value := cpu.Memory[effectiveAddress]       // Get the value from memory
	cpu.A |= value                              // OR the value with the accumulator
	cpu.updateZeroAndNegativeFlags(cpu.A)       // Update the zero and negative flags

	if (address & 0xFF00) != (effectiveAddress & 0xFF00) {
		// Page boundary crossed
		return 5, false
	} else {
		return 4, false
	}
}

func (cpu *CPU) opcode0x19() (byte, bool) {
	low := cpu.fetch()                          // Fetch the low byte of the address
	high := cpu.fetch()                         // Fetch the high byte of the address
	address := uint16(high)<<8 | uint16(low)    // Combine to form the full address
	effectiveAddress := address + uint16(cpu.Y) // Add the Y register to the address
	value := cpu.Memory[effectiveAddress]       // Get the value from memory
	cpu.A |= value                              // OR the value with the accumulator
	cpu.updateZeroAndNegativeFlags(cpu.A)       // Update the zero and negative flags

	if (effectiveAddress & 0xFF00) != (address & 0xFF00) {
		// Page boundary crossed
		return 5, false
	} else {
		return 4, false
	}
}

func (cpu *CPU) opcode0x01() (byte, bool) {
	baseAddress := cpu.fetch() // Fetch the base address
	// Calculate the effective address using (indirect,X) addressing
	effectiveAddressLow := cpu.Memory[(baseAddress+cpu.X)&0xFF]
	effectiveAddressHigh := cpu.Memory[(baseAddress+cpu.X+1)&0xFF]
	effectiveAddress := uint16(effectiveAddressHigh)<<8 | uint16(effectiveAddressLow)
	value := cpu.Memory[effectiveAddress] // Get the value from memory
	cpu.A |= value                        // OR the value with the accumulator
	cpu.updateZeroAndNegativeFlags(cpu.A) // Update the zero and negative flags
	return 6, false                       // ORA (indirect,X) takes 6 cycles
}

func (cpu *CPU) opcode0x11() (byte, bool) {
	baseAddress := cpu.fetch() // Fetch the base address
	effectiveAddressLow := cpu.Memory[baseAddress]
	effectiveAddressHigh := cpu.Memory[(baseAddress+1)&0xFF]
	effectiveAddress := uint16(effectiveAddressHigh)<<8 | uint16(effectiveAddressLow)
	finalAddress := effectiveAddress + uint16(cpu.Y)
	value := cpu.Memory[finalAddress]     // Add Y register to the effective address
	cpu.A |= value                        // OR the value with the accumulator
	cpu.updateZeroAndNegativeFlags(cpu.A) // Update the zero and negative flags
	if (effectiveAddress & 0xFF00) != (finalAddress & 0xFF00) {
		// Page boundary crossed
		return 6, false
	} else {
		return 5, false
	}

}

/*
AND Memory with Accumulator

A AND M -> A
N	Z	C	I	D	V
+	+	-	-	-	-
addressing		assembler		opc	bytes	cycles
immediate		AND #oper		29	2		2
zeropage		AND oper		25	2		3
zeropage,X		AND oper,X		35	2		4
absolute		AND oper		2D	3		4
absolute,X		AND oper,X		3D	3		4*
absolute,Y		AND oper,Y		39	3		4*
(indirect,X)	AND (oper,X)	21	2		6
(indirect),Y	AND (oper),Y	31	2		5*
*/

func (cpu *CPU) opcode0x29() (byte, bool) {
	value := cpu.fetch()
	cpu.A &= value
	cpu.updateZeroAndNegativeFlags(cpu.A)
	return 2, false
}

func (cpu *CPU) opcode0x25() (byte, bool) {
	address := uint16(cpu.fetch())
	value := cpu.Memory[address]
	cpu.A &= value
	cpu.updateZeroAndNegativeFlags(cpu.A)
	return 3, false
}

func (cpu *CPU) opcode0x35() (byte, bool) {
	baseAddress := cpu.fetch()
	address := uint16((baseAddress + cpu.X) & 0xFF)
	value := cpu.Memory[address]
	cpu.A &= value
	cpu.updateZeroAndNegativeFlags(cpu.A)
	return 4, false
}

func (cpu *CPU) opcode0x2D() (byte, bool) {
	low := cpu.fetch()
	high := cpu.fetch()
	address := uint16(high)<<8 | uint16(low)
	value := cpu.Memory[address]
	cpu.A &= value
	cpu.updateZeroAndNegativeFlags(cpu.A)
	return 4, false
}

func (cpu *CPU) opcode0x3D() (byte, bool) {
	low := cpu.fetch()
	high := cpu.fetch()
	address := uint16(high)<<8 | uint16(low)
	effectiveAddress := address + uint16(cpu.X)
	value := cpu.Memory[effectiveAddress]
	cpu.A &= value
	cpu.updateZeroAndNegativeFlags(cpu.A)
	if (address & 0xFF00) != (effectiveAddress & 0xFF00) {
		return 5, false
	} else {
		return 4, false
	}
}

func (cpu *CPU) opcode0x39() (byte, bool) {
	low := cpu.fetch()
	high := cpu.fetch()
	address := uint16(high)<<8 | uint16(low)
	effectiveAddress := address + uint16(cpu.Y)
	value := cpu.Memory[effectiveAddress]
	cpu.A &= value
	cpu.updateZeroAndNegativeFlags(cpu.A)
	if (effectiveAddress & 0xFF00) != (address & 0xFF00) {
		return 5, false
	} else {
		return 4, false
	}
}

func (cpu *CPU) opcode0x21() (byte, bool) {
	baseAddress := cpu.fetch()
	effectiveAddressLow := cpu.Memory[(baseAddress+cpu.X)&0xFF]
	effectiveAddressHigh := cpu.Memory[(baseAddress+cpu.X+1)&0xFF]
	effectiveAddress := uint16(effectiveAddressHigh)<<8 | uint16(effectiveAddressLow)
	value := cpu.Memory[effectiveAddress]
	cpu.A &= value
	cpu.updateZeroAndNegativeFlags(cpu.A)
	return 6, false
}

func (cpu *CPU) opcode0x31() (byte, bool) {
	baseAddress := cpu.fetch()
	effectiveAddressLow := cpu.Memory[baseAddress]
	effectiveAddressHigh := cpu.Memory[(baseAddress+1)&0xFF]
	effectiveAddress := uint16(effectiveAddressHigh)<<8 | uint16(effectiveAddressLow)
	finalAddress := effectiveAddress + uint16(cpu.Y)
	value := cpu.Memory[finalAddress]
	cpu.A &= value
	cpu.updateZeroAndNegativeFlags(cpu.A)
	if (effectiveAddress & 0xFF00) != (finalAddress & 0xFF00) {
		return 6, false
	} else {
		return 5, false
	}

}

/*
EOR
Exclusive-OR Memory with Accumulator

A EOR M -> A
N	Z	C	I	D	V
+	+	-	-	-	-
addressing		assembler		opc	bytes	cycles
immediate		EOR #oper		49	2		2
zeropage		EOR oper		45	2		3
zeropage,X		EOR oper,X		55	2		4
absolute		EOR oper		4D	3		4
absolute,X		EOR oper,X		5D	3		4*
absolute,Y		EOR oper,Y		59	3		4*
(indirect,X)	EOR (oper,X)	41	2		6
(indirect),Y	EOR (oper),Y	51	2		5*
*/

func (cpu *CPU) opcode0x49() (byte, bool) {
	value := cpu.fetch()
	cpu.A ^= value
	cpu.updateZeroAndNegativeFlags(cpu.A)
	return 2, false
}

func (cpu *CPU) opcode0x45() (byte, bool) {
	address := uint16(cpu.fetch())
	value := cpu.Memory[address]
	cpu.A ^= value
	cpu.updateZeroAndNegativeFlags(cpu.A)
	return 3, false
}

func (cpu *CPU) opcode0x55() (byte, bool) {
	baseAddress := cpu.fetch()
	address := uint16((baseAddress + cpu.X) & 0xFF)
	value := cpu.Memory[address]
	cpu.A ^= value
	cpu.updateZeroAndNegativeFlags(cpu.A)
	return 4, false
}

func (cpu *CPU) opcode0x4D() (byte, bool) {
	low := cpu.fetch()
	high := cpu.fetch()
	address := uint16(high)<<8 | uint16(low)
	value := cpu.Memory[address]
	cpu.A ^= value
	cpu.updateZeroAndNegativeFlags(cpu.A)
	return 4, false
}

func (cpu *CPU) opcode0x5D() (byte, bool) {
	low := cpu.fetch()
	high := cpu.fetch()
	address := uint16(high)<<8 | uint16(low)
	effectiveAddress := address + uint16(cpu.X)
	value := cpu.Memory[effectiveAddress]
	cpu.A ^= value
	cpu.updateZeroAndNegativeFlags(cpu.A)
	if (address & 0xFF00) != (effectiveAddress & 0xFF00) {
		return 5, false
	} else {
		return 4, false
	}
}

func (cpu *CPU) opcode0x59() (byte, bool) {
	low := cpu.fetch()
	high := cpu.fetch()
	address := uint16(high)<<8 | uint16(low)
	effectiveAddress := address + uint16(cpu.Y)
	value := cpu.Memory[effectiveAddress]
	cpu.A ^= value
	cpu.updateZeroAndNegativeFlags(cpu.A)
	if (effectiveAddress & 0xFF00) != (address & 0xFF00) {
		return 5, false
	} else {
		return 4, false
	}
}

func (cpu *CPU) opcode0x41() (byte, bool) {
	baseAddress := cpu.fetch()
	effectiveAddressLow := cpu.Memory[(baseAddress+cpu.X)&0xFF]
	effectiveAddressHigh := cpu.Memory[(baseAddress+cpu.X+1)&0xFF]
	effectiveAddress := uint16(effectiveAddressHigh)<<8 | uint16(effectiveAddressLow)
	value := cpu.Memory[effectiveAddress]
	cpu.A ^= value
	cpu.updateZeroAndNegativeFlags(cpu.A)
	return 6, false

}

func (cpu *CPU) opcode0x51() (byte, bool) {
	baseAddress := cpu.fetch()
	effectiveAddressLow := cpu.Memory[baseAddress]
	effectiveAddressHigh := cpu.Memory[(baseAddress+1)&0xFF]
	effectiveAddress := uint16(effectiveAddressHigh)<<8 | uint16(effectiveAddressLow)
	finalAddress := effectiveAddress + uint16(cpu.Y)
	value := cpu.Memory[finalAddress]
	cpu.A ^= value
	cpu.updateZeroAndNegativeFlags(cpu.A)
	if (effectiveAddress & 0xFF00) != (finalAddress & 0xFF00) {
		return 6, false
	} else {
		return 5, false
	}
}
