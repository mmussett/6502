package nmos6502

/*
ASL - Shift Left One Bit (Memory or Accumulator)

C <- [76543210] <- 0
N	Z	C	I	D	V
+	+	+	-	-	-
addressing		assembler	opc	bytes	cycles
accumulator		ASL A		0A	1		2
zeropage		ASL oper	06	2		5
zeropage,X		ASL oper,X	16	2		6
absolute		ASL oper	0E	3		6
absolute,X		ASL oper,X	1E	3		7
*/

func (cpu *CPU) shiftLeft(value byte) byte {
	// Set the carry flag if the leftmost bit is 1
	if value&0x80 != 0 {
		cpu.setCarryFlag()
	} else {
		cpu.clearCarryFlag()
	}

	// Shift the value left by one bit
	result := value << 1

	// Update zero and negative flags based on the result
	cpu.updateZeroAndNegativeFlags(result)

	return result
}

func (cpu *CPU) opcode0x06() (byte, bool) { // ASL Zero Page
	address := uint16(cpu.fetch())             // Fetch the zero-page address
	value := cpu.Memory[address]               // Retrieve the value from memory
	cpu.Memory[address] = cpu.shiftLeft(value) // Shift the value left and store it back
	return 5, false                            // ASL zero page takes 5 cycles
}

func (cpu *CPU) opcode0x16() (byte, bool) { // ASL Zero Page,X
	baseAddress := cpu.fetch()                      // Fetch the zero-page base address
	address := uint16((baseAddress + cpu.X) & 0xFF) // Calculate the effective address with zero-page wrap-around
	value := cpu.Memory[address]                    // Retrieve the value from memory
	cpu.Memory[address] = cpu.shiftLeft(value)      // Shift the value left and store it back
	return 6, false                                 // ASL zero page,X takes 6 cycles
}
func (cpu *CPU) opcode0x0E() (byte, bool) { // ASL Absolute
	low := cpu.fetch()                         // Fetch the low byte of the address
	high := cpu.fetch()                        // Fetch the high byte of the address
	address := uint16(high)<<8 | uint16(low)   // Combine to form the full address
	value := cpu.Memory[address]               // Retrieve the value from memory
	cpu.Memory[address] = cpu.shiftLeft(value) // Shift the value left and store it back
	return 6, false                            // ASL absolute takes 6 cycles
}

func (cpu *CPU) opcode0x1E() (byte, bool) { // ASL Absolute,X
	low := cpu.fetch()                                  // Fetch the low byte of the address
	high := cpu.fetch()                                 // Fetch the high byte of the address
	address := uint16(high)<<8 | uint16(low)            // Combine to form the full address
	effectiveAddress := address + uint16(cpu.X)         // Add the X register to the address
	value := cpu.Memory[effectiveAddress]               // Retrieve the value from memory
	cpu.Memory[effectiveAddress] = cpu.shiftLeft(value) // Shift the value left and store it back
	return 7, false                                     // ASL absolute,X takes 7 cycles
}

func (cpu *CPU) opcode0x0A() (byte, bool) {
	cpu.A = cpu.shiftLeft(cpu.A)
	return 1, false
}

/*
LSR - Shift One Bit Right (Memory or Accumulator)

0 -> [76543210] -> C
N	Z	C	I	D	V
0	+	+	-	-	-
addressing	assembler	opc	bytes	cycles
accumulator	LSR A		4A	1		2
zeropage	LSR oper	46	2		5
zeropage,X	LSR oper,X	56	2		6
absolute	LSR oper	4E	3		6
absolute,X	LSR oper,X	5E	3		7
*/

func (cpu *CPU) opcode0x4A() (byte, bool) { // LSR Accumulator
	cpu.A = cpu.shiftRight(cpu.A)
	return 2, false // LSR accumulator takes 2 cycles
}

func (cpu *CPU) opcode0x46() (byte, bool) { // LSR Zero Page
	address := uint16(cpu.fetch())
	value := cpu.Memory[address]
	cpu.Memory[address] = cpu.shiftRight(value)
	return 5, false // LSR zeropage takes 5 cycles
}

func (cpu *CPU) opcode0x56() (byte, bool) { // LSR Zero Page,X
	baseAddress := cpu.fetch()
	address := uint16(baseAddress + cpu.X)
	value := cpu.Memory[address]
	cpu.Memory[address] = cpu.shiftRight(value)
	return 6, false // LSR zeropage,X takes 6 cycles
}

func (cpu *CPU) opcode0x4E() (byte, bool) { // LSR Absolute
	low := cpu.fetch()
	high := cpu.fetch()
	address := uint16(high)<<8 | uint16(low)
	value := cpu.Memory[address]
	cpu.Memory[address] = cpu.shiftRight(value)
	return 6, false // LSR absolute takes 6 cycles
}

func (cpu *CPU) opcode0x5E() (byte, bool) { // LSR Absolute,X
	low := cpu.fetch()
	high := cpu.fetch()
	address := uint16(high)<<8 | uint16(low)
	effectiveAddress := address + uint16(cpu.X)
	value := cpu.Memory[effectiveAddress]
	cpu.Memory[effectiveAddress] = cpu.shiftRight(value)
	return 7, false // LSR absolute,X takes 7 cycles
}

func (cpu *CPU) shiftRight(value byte) byte {
	// Set the carry flag if the rightmost bit is 1
	if value&0x01 != 0 {
		cpu.setCarryFlag()
	} else {
		cpu.clearCarryFlag()
	}

	// Shift the value right by one bit
	result := value >> 1

	// Update zero and negative flags based on the result
	cpu.updateZeroAndNegativeFlags(result)

	return result
}

/*
ROL - Rotate One Bit Left (Memory or Accumulator)

C <- [76543210] <- C
N	Z	C	I	D	V
+	+	+	-	-	-
addressing	assembler	opc	bytes	cycles
accumulator	ROL A		2A	1		2
zeropage	ROL oper	26	2		5
zeropage,X	ROL oper,X	36	2		6
absolute	ROL oper	2E	3		6
absolute,X	ROL oper,X	3E	3		7
*/

func (cpu *CPU) opcode0x2A() (byte, bool) { // ROL Accumulator
	cpu.A = cpu.rotateLeft(cpu.A)
	return 2, false // ROL accumulator takes 2 cycles
}

func (cpu *CPU) opcode0x26() (byte, bool) { // ROL Zero Page
	address := uint16(cpu.fetch())
	value := cpu.Memory[address]
	cpu.Memory[address] = cpu.rotateLeft(value)
	return 5, false // ROL zeropage takes 5 cycles
}

func (cpu *CPU) opcode0x36() (byte, bool) { // ROL Zero Page,X
	baseAddress := cpu.fetch()
	address := uint16(baseAddress + cpu.X)
	value := cpu.Memory[address]
	cpu.Memory[address] = cpu.rotateLeft(value)
	return 6, false // ROL zeropage,X takes 6 cycles
}

func (cpu *CPU) opcode0x2E() (byte, bool) { // ROL Absolute
	low := cpu.fetch()
	high := cpu.fetch()
	address := uint16(high)<<8 | uint16(low)
	value := cpu.Memory[address]
	cpu.Memory[address] = cpu.rotateLeft(value)
	return 6, false // ROL absolute takes 6 cycles
}

func (cpu *CPU) opcode0x3E() (byte, bool) { // ROL Absolute,X
	low := cpu.fetch()
	high := cpu.fetch()
	address := uint16(high)<<8 | uint16(low)
	effectiveAddress := address + uint16(cpu.X)
	value := cpu.Memory[effectiveAddress]
	cpu.Memory[effectiveAddress] = cpu.rotateLeft(value)
	return 7, false // ROL absolute,X takes 7 cycles
}

func (cpu *CPU) rotateLeft(value byte) byte {
	carryIn := byte(0)
	if cpu.Status&Carry != 0 {
		carryIn = 1
	}

	// Set the carry flag if the leftmost bit is 1
	if value&0x80 != 0 {
		cpu.setCarryFlag()
	} else {
		cpu.clearCarryFlag()
	}

	// Rotate left
	result := (value << 1) | carryIn

	// Update zero and negative flags based on the result
	cpu.updateZeroAndNegativeFlags(result)

	return result
}

/*
ROR - Rotate One Bit Right (Memory or Accumulator)

C -> [76543210] -> C
N	Z	C	I	D	V
+	+	+	-	-	-
addressing	assembler	opc	bytes	cycles
accumulator	ROR A		6A	1		2
zeropage	ROR oper	66	2		5
zeropage,X	ROR oper,X	76	2		6
absolute	ROR oper	6E	3		6
absolute,X	ROR oper,X	7E	3		7
*/

func (cpu *CPU) opcode0x6A() (byte, bool) { // ROR Accumulator
	cpu.A = cpu.rotateRight(cpu.A)
	return 2, false // ROR accumulator takes 2 cycles
}

func (cpu *CPU) opcode0x66() (byte, bool) { // ROR Zero Page
	address := uint16(cpu.fetch())
	value := cpu.Memory[address]
	cpu.Memory[address] = cpu.rotateRight(value)
	return 5, false // ROR zero page takes 5 cycles
}

func (cpu *CPU) opcode0x76() (byte, bool) { // ROR Zero Page,X
	baseAddress := cpu.fetch()
	address := uint16((baseAddress + cpu.X) & 0xFF)
	value := cpu.Memory[address]
	cpu.Memory[address] = cpu.rotateRight(value)
	return 6, false // ROR zero page,X takes 6 cycles
}

func (cpu *CPU) opcode0x6E() (byte, bool) { // ROR Absolute
	low := cpu.fetch()
	high := cpu.fetch()
	address := uint16(high)<<8 | uint16(low)
	value := cpu.Memory[address]
	cpu.Memory[address] = cpu.rotateRight(value)
	return 6, false // ROR absolute takes 6 cycles
}

func (cpu *CPU) opcode0x7E() (byte, bool) { // ROR Absolute,X
	low := cpu.fetch()
	high := cpu.fetch()
	address := uint16(high)<<8 | uint16(low)
	effectiveAddress := address + uint16(cpu.X)
	value := cpu.Memory[effectiveAddress]
	cpu.Memory[effectiveAddress] = cpu.rotateRight(value)
	return 7, false // ROR absolute,X takes 7 cycles
}

func (cpu *CPU) rotateRight(value byte) byte {
	carryIn := byte(0)
	if cpu.Status&Carry != 0 {
		carryIn = 0x80
	}

	// Set the carry flag if the rightmost bit is 1
	if value&0x01 != 0 {
		cpu.setCarryFlag()
	} else {
		cpu.clearCarryFlag()
	}

	// Rotate right
	result := (value >> 1) | carryIn

	// Update zero and negative flags based on the result
	cpu.updateZeroAndNegativeFlags(result)

	return result
}
