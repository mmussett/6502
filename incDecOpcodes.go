package nmos6502

/*
Decrements & Increments

DEC decrement (memory)
DEX decrement X
DEY decrement Y
INC increment (memory)
INX increment X
INY increment Y
*/

/*
DEC -Decrement Memory by One

M - 1 -> M
N	Z	C	I	D	V
+	+	-	-	-	-
addressing	assembler	opc	bytes	cycles
zeropage	DEC oper	C6	2		5
zeropage,X	DEC oper,X	D6	2		6
absolute	DEC oper	CE	3		6
absolute,X	DEC oper,X	DE	3		7
*/
func decrementMemory(cpu *CPU, addr uint16) {
	cpu.Memory[addr]--
}

func incrementMemory(cpu *CPU, addr uint16) {
	cpu.Memory[addr]++
}

func (cpu *CPU) opcode0xC6() (byte, bool) { // DEC Zero Page
	address := uint16(cpu.fetch()) // Fetch the zero page address
	cpu.Memory[address]--
	cpu.updateZeroAndNegativeFlags(cpu.Memory[address])
	return 5, false // DEC zero page takes 5 cycles
}

func (cpu *CPU) opcode0xD6() (byte, bool) { // DEC Zero Page,X
	baseAddress := cpu.fetch()                      // Fetch the zero page base address
	address := uint16((baseAddress + cpu.X) & 0xFF) // Calculate the effective address with zero page wrap-around
	cpu.Memory[address]--
	cpu.updateZeroAndNegativeFlags(cpu.Memory[address])
	return 6, false // DEC zero page,X takes 6 cycles
}

func (cpu *CPU) opcode0xCE() (byte, bool) { // DEC Absolute
	low := cpu.fetch()                       // Fetch the low byte of the address
	high := cpu.fetch()                      // Fetch the high byte of the address
	address := uint16(high)<<8 | uint16(low) // Combine to form the full address
	cpu.Memory[address]--
	cpu.updateZeroAndNegativeFlags(cpu.Memory[address])
	return 6, false // DEC absolute takes 6 cycles
}

func (cpu *CPU) opcode0xDE() (byte, bool) { // DEC Absolute,X
	low := cpu.fetch()                          // Fetch the low byte of the address
	high := cpu.fetch()                         // Fetch the high byte of the address
	address := uint16(high)<<8 | uint16(low)    // Combine to form the full address
	effectiveAddress := address + uint16(cpu.X) // Add the X register to the address
	cpu.Memory[address]--
	cpu.updateZeroAndNegativeFlags(cpu.Memory[effectiveAddress])
	return 7, false // DEC absolute,X takes 7 cycles
}

/*
DEX - Decrement Index X by One

X - 1 -> X
N	Z	C	I	D	V
+	+	-	-	-	-
addressing	assembler	opc	bytes	cycles
implied	DEX				CA	1		2
*/
func (cpu *CPU) opcode0xCA() (byte, bool) { // DEX Implied
	cpu.X--
	cpu.updateZeroAndNegativeFlags(cpu.X)
	return 2, false // DEX implied takes 2 cycles
}

/*
DEY
Decrement Index Y by One

Y - 1 -> Y
N	Z	C	I	D	V
+	+	-	-	-	-
addressing	assembler	opc	bytes	cycles
implied		DEY			88	1		2
*/

func (cpu *CPU) opcode0x88() (byte, bool) { // DEY Implied
	cpu.Y--
	cpu.updateZeroAndNegativeFlags(cpu.Y)
	return 2, false // DEY implied takes 2 cycles
}

/*
INC - Increment Memory by One

M + 1 -> M
N	Z	C	I	D	V
+	+	-	-	-	-
addressing	assembler	opc	bytes	cycles
zeropage	INC oper	E6	2		5
zeropage,X	INC oper,X	F6	2		6
absolute	INC oper	EE	3		6
absolute,X	INC oper,X	FE	3		7
*/

func (cpu *CPU) opcode0xE6() (byte, bool) { // INC Zero Page
	address := uint16(cpu.fetch())
	incrementMemory(cpu, address)
	cpu.updateZeroAndNegativeFlags(cpu.Memory[address])
	return 5, false // INC zeropage takes 5 cycles
}

func (cpu *CPU) opcode0xF6() (byte, bool) { // INC Zero Page,X
	baseAddress := cpu.fetch()
	address := uint16((baseAddress + cpu.X) & 0xFF)
	incrementMemory(cpu, address)
	cpu.updateZeroAndNegativeFlags(cpu.Memory[address])
	return 6, false // INC zeropage,X takes 6 cycles
}

func (cpu *CPU) opcode0xEE() (byte, bool) { // INC Absolute
	low := cpu.fetch()
	high := cpu.fetch()
	address := uint16(high)<<8 | uint16(low)
	incrementMemory(cpu, address)
	cpu.updateZeroAndNegativeFlags(cpu.Memory[address])
	return 6, false // INC absolute takes 6 cycles
}

func (cpu *CPU) opcode0xFE() (byte, bool) { // INC Absolute,X
	low := cpu.fetch()
	high := cpu.fetch()
	address := uint16(high)<<8 | uint16(low)
	effectiveAddress := address + uint16(cpu.X)
	incrementMemory(cpu, effectiveAddress)
	cpu.updateZeroAndNegativeFlags(cpu.Memory[effectiveAddress])
	return 7, false // INC absolute,X takes 7 cycles
}

/*
INX - Increment Index X by One

X + 1 -> X
N	Z	C	I	D	V
+	+	-	-	-	-
addressing	assembler	opc	bytes	cycles
implied		INX			E8	1		2
*/

func (cpu *CPU) opcode0xE8() (byte, bool) { // INX Implied
	cpu.X++
	cpu.updateZeroAndNegativeFlags(cpu.X)
	return 2, false // INX implied takes 2 cycles
}

/*
INY
Increment Index Y by One

Y + 1 -> Y
N	Z	C	I	D	V
+	+	-	-	-	-
addressing	assembler	opc	bytes	cycles
implied		INY			C8	1		2
*/
func (cpu *CPU) opcode0xC8() (byte, bool) { // INY Implied
	cpu.Y++
	cpu.updateZeroAndNegativeFlags(cpu.Y)
	return 2, false // INY implied takes 2 cycles
}
