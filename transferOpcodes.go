package nmos6502

// Load, store, inter-register transfer
//
// LDA load accumulator

// LDX load X
// LDY load Y
// STA store accumulator
// STX store X
// STY store Y
// TAX transfer accumulator to X
// TAY transfer accumulator to Y
// TSX transfer stack pointer to X
// TXA transfer X to accumulator
// TXS transfer X to stack pointer
// TYA transfer Y to accumulator

/**
 * LDA - Load Accumulator with Memory
 *
 *      A,Z,N = M
 *
 *      addressing    assembler    opc  bytes  cyles
 *      --------------------------------------------
 *      immediate     LDA #oper     A9   2     2
 *      zeropage      LDA oper      A5   2     3
 *      zeropage,X    LDA oper,X    B5   2     4
 *      absolute      LDA oper      AD   3     4
 *      absolute,X    LDA oper,X    BD   3     4
 *      absolute,Y    LDA oper,Y    B9   3     4
 *      (indirect,X)  LDA (oper,X)  A1   2     6
 *      (indirect),Y  LDA (oper),Y  B1   2     5+
 */

func (cpu *CPU) opcode0xA9() (byte, bool) {
	value := cpu.fetch()
	cpu.A = value
	cpu.updateZeroAndNegativeFlags(cpu.A)
	return 2, false
}

func (cpu *CPU) opcode0xA5() (byte, bool) { // LDA Zero Page
	address := uint16(cpu.fetch())
	value := cpu.Memory[address]
	cpu.A = value
	cpu.updateZeroAndNegativeFlags(cpu.A)
	return 3, false
}
func (cpu *CPU) opcode0xB5() (byte, bool) { // LDA Zero Page,X
	baseAddress := cpu.fetch()
	address := uint16((baseAddress + cpu.X) & 0xFF) // Zero page wrap-around
	value := cpu.Memory[address]
	cpu.A = value
	cpu.updateZeroAndNegativeFlags(cpu.A)
	return 4, false
}
func (cpu *CPU) opcode0xAD() (byte, bool) { // LDA Absolute
	low := cpu.fetch()
	high := cpu.fetch()
	address := uint16(high)<<8 | uint16(low)
	value := cpu.Memory[address]
	cpu.A = value
	cpu.updateZeroAndNegativeFlags(cpu.A)
	return 4, false
}
func (cpu *CPU) opcode0xBD() (byte, bool) { // LDA Absolute,X
	low := cpu.fetch()
	high := cpu.fetch()
	address := uint16(high)<<8 | uint16(low)
	value := cpu.Memory[address+uint16(cpu.X)]
	cpu.A = value
	cpu.updateZeroAndNegativeFlags(cpu.A)
	return 4, false
}
func (cpu *CPU) opcode0xB9() (byte, bool) { // LDA Absolute,Y
	low := cpu.fetch()
	high := cpu.fetch()
	address := uint16(high)<<8 | uint16(low)
	value := cpu.Memory[address+uint16(cpu.Y)]
	cpu.A = value
	cpu.updateZeroAndNegativeFlags(cpu.A)
	return 4, false
}
func (cpu *CPU) opcode0xA1() (byte, bool) { // LDA (oper,X)
	baseAddress := cpu.fetch()
	// Calculate the effective address
	effectiveAddressLow := cpu.Memory[(baseAddress+cpu.X)&0xFF]
	effectiveAddressHigh := cpu.Memory[(baseAddress+cpu.X+1)&0xFF]
	effectiveAddress := uint16(effectiveAddressHigh)<<8 | uint16(effectiveAddressLow)
	value := cpu.Memory[effectiveAddress]
	cpu.A = value
	cpu.updateZeroAndNegativeFlags(cpu.A)
	return 6, false
}
func (cpu *CPU) opcode0xB1() (byte, bool) { // LDA (oper),Y
	baseAddress := cpu.fetch()
	effectiveAddressLow := cpu.Memory[baseAddress]
	effectiveAddressHigh := cpu.Memory[(baseAddress+1)&0xFF]
	effectiveAddress := uint16(effectiveAddressHigh)<<8 | uint16(effectiveAddressLow)
	address := effectiveAddress + uint16(cpu.Y)
	value := cpu.Memory[address]
	cpu.A = value
	cpu.updateZeroAndNegativeFlags(cpu.A)

	if (address & 0xFF00) != (effectiveAddress & 0xFF00) {
		return 6, false
	} else {
		return 5, false
	}

}

/**
 * LDX - Load Index X With Memory
 *
 *      X,Z,N = M
 *
 *      addressing    assembler    opc  bytes  cyles
 *      --------------------------------------------
 *      immediate     LDX #oper     A2   2     2
 *      zeropage      LDX oper      A6   2     3
 *      zeropage,Y    LDX oper,Y    B6   2     4
 *      absolute      LDX oper      AE   3     4
 *      absolute,Y    LDX oper,Y    BE   3     4
 */

func (cpu *CPU) opcode0xA2() (byte, bool) { // LDX Immediate
	value := cpu.fetch()
	cpu.X = value
	cpu.updateZeroAndNegativeFlags(cpu.X)
	return 2, false
}
func (cpu *CPU) opcode0xA6() (byte, bool) { // LDX Zero Page
	address := uint16(cpu.fetch())
	value := cpu.Memory[address]
	cpu.X = value
	cpu.updateZeroAndNegativeFlags(cpu.X)
	return 3, false
}
func (cpu *CPU) opcode0xB6() (byte, bool) { // LDX Zero Page,Y
	baseAddress := uint16(cpu.fetch())
	address := baseAddress + uint16(cpu.Y)
	var value = cpu.Memory[address]
	cpu.X = value
	cpu.updateZeroAndNegativeFlags(cpu.X)
	return 4, false
}
func (cpu *CPU) opcode0xAE() (byte, bool) { // LDX Absolute
	low := cpu.fetch()
	high := cpu.fetch()
	address := uint16(high)<<8 | uint16(low)
	value := cpu.Memory[address]
	cpu.X = value
	cpu.updateZeroAndNegativeFlags(cpu.X)
	return 4, false
}
func (cpu *CPU) opcode0xBE() (byte, bool) { // LDX Absolute,Y
	low := cpu.fetch()
	high := cpu.fetch()
	address := uint16(high)<<8 | uint16(low)
	value := cpu.Memory[address+uint16(cpu.Y)]
	cpu.X = value
	cpu.updateZeroAndNegativeFlags(cpu.X)
	return 4, false
}

/**
 * LDY - Load Index Y With Memory
 *
 *      Y,Z,N = M
 *
 *      addressing    assembler    opc  bytes  cyles
 *      --------------------------------------------
 *      immediate     LDY #oper     A0   2     2
 *      zeropage      LDY oper      A4   2     3
 *      zeropage,X    LDY oper,X    B4   2     4
 *      absolute      LDY oper      AC   3     4
 *      absolute,X    LDY oper,X    BC   3     4
 */

func (cpu *CPU) opcode0xA0() (byte, bool) { // LDY Immediate
	value := cpu.fetch()
	cpu.Y = value
	cpu.updateZeroAndNegativeFlags(cpu.Y)
	return 2, false
}
func (cpu *CPU) opcode0xA4() (byte, bool) { // LDY Zero Page
	address := uint16(cpu.fetch())
	value := cpu.Memory[address]
	cpu.Y = value
	cpu.updateZeroAndNegativeFlags(cpu.Y)
	return 3, false
}
func (cpu *CPU) opcode0xB4() (byte, bool) { // LDY Zero Page,X
	baseAddress := uint16(cpu.fetch())
	address := baseAddress + uint16(cpu.X)
	var value = cpu.Memory[address]
	cpu.Y = value
	cpu.updateZeroAndNegativeFlags(cpu.Y)
	return 4, false
}
func (cpu *CPU) opcode0xAC() (byte, bool) { // LDY Absolute
	high := cpu.fetch()
	low := cpu.fetch()
	address := uint16(high)<<8 | uint16(low)
	value := cpu.Memory[address]
	cpu.Y = value
	cpu.updateZeroAndNegativeFlags(cpu.Y)
	return 4, false
}
func (cpu *CPU) opcode0xBC() (byte, bool) { // LDY Absolute,X
	high := cpu.fetch()
	low := cpu.fetch()
	address := uint16(high)<<8 | uint16(low)
	value := cpu.Memory[address+uint16(cpu.X)]
	cpu.Y = value
	cpu.updateZeroAndNegativeFlags(cpu.Y)
	return 4, false

}

/**
 * STA - Store Accumulator in Memory
 *
 *      M = A
 *
 *      addressing    assembler    opc  bytes  cyles
 *      --------------------------------------------
 *      zeropage      STA oper      85   2     3
 *      zeropage,X    STA oper,X    95   2     4
 *      absolute      STA oper      8D   3     4
 *      absolute,X    STA oper,X    9D   3     5
 *      absolute,Y    STA oper,Y    99   3	   5
 *      (indirect,X)  STA (oper,X)  81   2     6
 *      (indirect),Y  STA (oper),Y  91   2     6
 */

func (cpu *CPU) opcode0x85() (byte, bool) { // STA Zero Page
	address := uint16(cpu.fetch())
	cpu.Memory[address] = cpu.A
	return 3, false
}
func (cpu *CPU) opcode0x95() (byte, bool) { // STA Zero Page,X
	baseAddress := cpu.fetch()
	address := uint16(baseAddress + cpu.X)
	cpu.Memory[address] = cpu.A
	return 4, false
}
func (cpu *CPU) opcode0x8D() (byte, bool) { // STA Absolute
	low := cpu.fetch()
	high := cpu.fetch()
	address := uint16(high)<<8 | uint16(low)
	cpu.Memory[address] = cpu.A
	return 4, false
}
func (cpu *CPU) opcode0x9D() (byte, bool) { // STA Absolute,X
	low := cpu.fetch()
	high := cpu.fetch()
	address := uint16(high)<<8 | uint16(low)
	cpu.Memory[address+uint16(cpu.X)] = cpu.A
	return 5, false
}
func (cpu *CPU) opcode0x99() (byte, bool) { // STA Absolute,Y
	low := cpu.fetch()
	high := cpu.fetch()
	address := uint16(high)<<8 | uint16(low)
	cpu.Memory[address+uint16(cpu.Y)] = cpu.A
	return 5, false
}

func (cpu *CPU) opcode0x81() (byte, bool) { // STA (oper,X)
	baseAddress := cpu.fetch()
	effectiveAddressLow := cpu.Memory[(baseAddress+cpu.X)&0xFF]
	effectiveAddressHigh := cpu.Memory[(baseAddress+cpu.X+1)&0xFF]
	effectiveAddress := uint16(effectiveAddressHigh)<<8 | uint16(effectiveAddressLow)
	cpu.Memory[effectiveAddress] = cpu.A
	return 6, false
}

func (cpu *CPU) opcode0x91() (byte, bool) { // STA (oper),Y
	baseAddress := cpu.fetch()
	effectiveAddressLow := cpu.Memory[baseAddress]
	effectiveAddressHigh := cpu.Memory[(baseAddress+1)&0xFF]
	effectiveAddress := uint16(effectiveAddressHigh)<<8 | uint16(effectiveAddressLow)
	cpu.Memory[effectiveAddress+uint16(cpu.Y)] = cpu.A
	return 6, false

}

/*
STX -Store Index X in Memory

X -> M
N	Z	C	I	D	V
-	-	-	-	-	-
addressing	assembler	opc	bytes	cycles
zeropage	STX oper	86	2		3
zeropage,Y	STX oper,Y	96	2		4
absolute	STX oper	8E	3		4
*/

func (cpu *CPU) opcode0x86() (byte, bool) { // STX Zero Page
	address := uint16(cpu.fetch())
	cpu.Memory[address] = cpu.X
	return 3, false
}

func (cpu *CPU) opcode0x96() (byte, bool) { // STX Zero Page,Y
	baseAddress := cpu.fetch()
	address := uint16(baseAddress + cpu.Y)
	cpu.Memory[address] = cpu.X
	return 4, false
}

func (cpu *CPU) opcode0x8E() (byte, bool) { // STX Absolute
	low := cpu.fetch()
	high := cpu.fetch()
	address := uint16(high)<<8 | uint16(low)
	cpu.Memory[address] = cpu.X

	return 4, false
}

/*
STY
Sore Index Y in Memory

Y -> M
N	Z	C	I	D	V
-	-	-	-	-	-
addressing	assembler	opc	bytes	cycles
zeropage	STY oper	84	2	3
zeropage,X	STY oper,X	94	2	4
absolute	STY oper	8C	3	4
*/

func (cpu *CPU) opcode0x84() (byte, bool) { // STY Zero Page
	address := uint16(cpu.fetch())
	cpu.Memory[address] = cpu.Y
	return 3, false
}

func (cpu *CPU) opcode0x94() (byte, bool) { // STY Zero Page,X
	baseAddress := cpu.fetch()
	address := uint16(baseAddress + cpu.X)
	cpu.Memory[address] = cpu.Y
	return 4, false
}

func (cpu *CPU) opcode0x8C() (byte, bool) { // STY Absolute
	low := cpu.fetch()
	high := cpu.fetch()
	address := uint16(high)<<8 | uint16(low)
	cpu.Memory[address] = cpu.Y
	return 4, false
}

/**
 * TAX - Transfer Accumulator To Index X
 *
 *      X = A
 *      N,Z = A
 *
 *      addressing    assembler    opc  bytes  cyles
 *      --------------------------------------------
 *      implied       TAX           AA   1     2
 */
func (cpu *CPU) opcode0xAA() (byte, bool) { // TAX
	cpu.X = cpu.A
	cpu.updateZeroAndNegativeFlags(cpu.X)
	return 2, false
}

/**
 * TXA - Transfer Index X To Accumulator
 *
 *      A = X
 *      N,Z = A
 *
 *      addressing    assembler    opc  bytes  cyles
 *      --------------------------------------------
 *      implied       TXA           8A   1     2
 */
func (cpu *CPU) opcode0x8A() (byte, bool) { // TXA
	cpu.A = cpu.X
	cpu.updateZeroAndNegativeFlags(cpu.A)
	return 2, false
}

/**
 * TYA - Transfer Index Y To Accumulator
 *
 *      A = Y
 *      N,Z = A
 *
 *      addressing    assembler    opc  bytes  cyles
 *      --------------------------------------------
 *      implied       TYA           98   1     2
 */
func (cpu *CPU) opcode0x98() (byte, bool) { // TYA
	cpu.A = cpu.Y
	cpu.updateZeroAndNegativeFlags(cpu.A)
	return 2, false
}

/**
 * TAY - Transfer Accumulator To Index Y
 *
 *      Y = A
 *      N,Z = A
 *
 *      addressing    assembler    opc  bytes  cyles
 *      --------------------------------------------
 *      implied       TAY           A8   1     2
 */
func (cpu *CPU) opcode0xA8() (byte, bool) { // TAY
	cpu.Y = cpu.A
	cpu.updateZeroAndNegativeFlags(cpu.Y)
	return 2, false
}

/**
 * TSX - Transfer Stack Pointer To Index X
 *
 *      X = SP
 *      N,Z = X
 *
 *      addressing    assembler    opc  bytes  cyles
 *      --------------------------------------------
 *      implied       TSX           BA   1     2
 */
func (cpu *CPU) opcode0xBA() (byte, bool) { // TSX
	cpu.X = cpu.SP
	cpu.updateZeroAndNegativeFlags(cpu.X)
	return 2, false
}

/**
 * TXS - Transfer Index X To Stack Register
 *
 *      SP = X
 *
 *      addressing    assembler    opc  bytes  cyles
 *      --------------------------------------------
 *      implied       TXS           9A   1     2
 */
func (cpu *CPU) opcode0x9A() (byte, bool) { // TXS
	cpu.SP = cpu.X
	return 2, false
}
