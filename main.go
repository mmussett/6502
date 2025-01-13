package nmos6502

import (
	"fmt"
)

type CPU struct {
	A, X, Y, SP byte        // Registers: Accumulator, X, Y, Stack Pointer, Program Counter
	PC          uint16      // Program Counter
	Status      byte        // Processor Status
	Memory      [65536]byte // 64KB of memory
}

func NewCPU() *CPU {
	return &CPU{
		SP: 0xFF, // Stack Pointer starts at 0xFF
	}
}

func (cpu *CPU) Execute(opcode byte) (cycles byte, illegalInstruction bool) {
	switch opcode {

	case 0x00: // BRK Implied
		return cpu.opcode0x00()
	case 0x01: // ORA (oper,X)
		return cpu.opcode0x01()
	case 0x05: // ORA Zero Page
		return cpu.opcode0x05()
	case 0x06: // ASL Zero Page
		return cpu.opcode0x06()
	case 0x08: // PHP Implied
		return cpu.opcode0x08()
	case 0x09: // ORA Immediate
		return cpu.opcode0x09()
	case 0x0A: // ASL Accumulator
		return cpu.opcode0x0A()
	case 0x0D: // ORA Absolute
		return cpu.opcode0x0D()
	case 0x0E: // ASL Absolute
		return cpu.opcode0x0E()
	case 0x10: // BPL Relative
		return cpu.opcode0x10()
	case 0x11: // ORA (oper),Y
		return cpu.opcode0x11()
	case 0x15: // ORA Zero Page,X
		return cpu.opcode0x15()
	case 0x16: // ASL Zero Page,X
		return cpu.opcode0x16()
	case 0x18: // CLC Implied
		return cpu.opcode0x18()
	case 0x19: // ORA Absolute,Y
		return cpu.opcode0x19()
	case 0x1D: // ORA Absolute,X
		return cpu.opcode0x1D()
	case 0x1E: // ASL Absolute,X
		return cpu.opcode0x1E()
	case 0x20: // JSR Absolute
		return cpu.opcode0x20()
	case 0x21: // AND (oper,X)
		return cpu.opcode0x21()
	case 0x24: // BIT Zero Page
		return cpu.opcode0x24()
	case 0x25: // AND Zero Page
		return cpu.opcode0x25()
	case 0x26: // ROL Zero Page
		return cpu.opcode0x26()
	case 0x28: // PLP Implied
		return cpu.opcode0x28()
	case 0x29: // AND Immediate
		return cpu.opcode0x29()
	case 0x2A: // ROL Accumulator
		return cpu.opcode0x2A()
	case 0x2C: // BIT Absolute
		return cpu.opcode0x2C()
	case 0x2D: // AND Absolute
		return cpu.opcode0x2D()
	case 0x2E: // ROL Absolute
		return cpu.opcode0x2E()
	case 0x30: // BMI Relative
		return cpu.opcode0x30()
	case 0x31: // AND (oper),Y
		return cpu.opcode0x31()
	case 0x35: // AND Zero Page,X
		return cpu.opcode0x35()
	case 0x36: // ROL Zero Page,X
		return cpu.opcode0x36()
	case 0x38: // SEC Implied
		return cpu.opcode0x38()
	case 0x39: // AND Absolute,Y
		return cpu.opcode0x39()
	case 0x3D: // AND Absolute,X
		return cpu.opcode0x3D()
	case 0x3E: // ROL Absolute,X
		return cpu.opcode0x3E()
	case 0x40: // RTI Implied
		return cpu.opcode0x40()
	case 0x41: // EOR (oper,X)
		return cpu.opcode0x41()
	case 0x45: // EOR Zero Page
		return cpu.opcode0x45()
	case 0x46: // LSR Zero Page
		return cpu.opcode0x46()
	case 0x48: // PHA Implied
		return cpu.opcode0x48()
	case 0x49: // EOR Immediate
		return cpu.opcode0x49()
	case 0x4A: // LSR Accumulator
		return cpu.opcode0x4A()
	case 0x4C: // JMP Absolute
		return cpu.opcode0x4C()
	case 0x4D: // EOR Absolute
		return cpu.opcode0x4D()
	case 0x4E: // LSR Absolute
		return cpu.opcode0x4E()
	case 0x50: // BVC Relative
		return cpu.opcode0x50()
	case 0x51: // EOR (oper),Y
		return cpu.opcode0x51()
	case 0x55: // EOR Zero Page,X
		return cpu.opcode0x55()
	case 0x56: // LSR Zero Page,X
		return cpu.opcode0x56()
	case 0x58: // CLI Implied
		return cpu.opcode0x58()
	case 0x59: // EOR Absolute,Y
		return cpu.opcode0x59()
	case 0x5D: // EOR Absolute,X
		return cpu.opcode0x5D()
	case 0x5E: // LSR Absolute,X
		return cpu.opcode0x5E()
	case 0x60: // RTS Implied
		return cpu.opcode0x60()
	case 0x61: // ADC (oper,X)
		return cpu.opcode0x61()
	case 0x65: // ADC Zero Page
		return cpu.opcode0x65()
	case 0x66: // ROR Zero Page
		return cpu.opcode0x66()
	case 0x68: // PLA Implied
		return cpu.opcode0x68()
	case 0x69: // ADC Immediate
		return cpu.opcode0x69()
	case 0x6A: // ROR Accumulator
		return cpu.opcode0x6A()
	case 0x6C: // JMP Indirect
		return cpu.opcode0x6C()
	case 0x6D: // ADC Absolute
		return cpu.opcode0x6D()
	case 0x6E: // ROR Absolute
		return cpu.opcode0x6E()
	case 0x70: // BVS Relative
		return cpu.opcode0x70()
	case 0x71: // EOR (oper),Y
		return cpu.opcode0x71()
	case 0x75: // ADC Zero Page,X
		return cpu.opcode0x75()
	case 0x76: // LSR Zero Page,X
		return cpu.opcode0x76()
	case 0x78: // SEI Implied
		return cpu.opcode0x78()
	case 0x79: // EOR Absolute,Y
		return cpu.opcode0x79()
	case 0x7D: // EOR Absolute,X
		return cpu.opcode0x7D()
	case 0x7E: // LSR Absolute,X
		return cpu.opcode0x7E()
	case 0x81: // STA (oper,X)
		return cpu.opcode0x81()
	case 0x84: // STY Zero Page
		return cpu.opcode0x84()
	case 0x85: // STA Zero Page
		return cpu.opcode0x85()
	case 0x86: // STX Zero Page
		return cpu.opcode0x86()
	case 0x88: // DEY Implied
		return cpu.opcode0x88()
	case 0x8A: // TXA Implied
		return cpu.opcode0x8A()
	case 0x8C: // STY Absolute
		return cpu.opcode0x8C()
	case 0x8D: // STA Absolute
		return cpu.opcode0x8D()
	case 0x8E: // STX Absolute
		return cpu.opcode0x8E()
	case 0x90: // BCC Relative
		return cpu.opcode0x90()
	case 0x91: // STA (oper),Y
		return cpu.opcode0x91()
	case 0x94: // STY Zero Page,X
		return cpu.opcode0x94()
	case 0x95: // STA Zero Page,X
		return cpu.opcode0x95()
	case 0x96: // STX Zero Page,Y
		return cpu.opcode0x96()
	case 0x98: // TYA Implied
		return cpu.opcode0x98()
	case 0x99: // STA Absolute,Y
		return cpu.opcode0x99()
	case 0x9A: // TXS Implied
		return cpu.opcode0x9A()
	case 0x9D: // STA Absolute,X
		return cpu.opcode0x9D()
	case 0xA1: // LDA (oper,X)
		return cpu.opcode0xA1()
	case 0xA2: // LDX Immediate
		return cpu.opcode0xA2()
	case 0xA4: // LDY Zero Page
		return cpu.opcode0xA4()
	case 0xA5: // LDA Zero Page
		return cpu.opcode0xA5()
	case 0xA6: // LDX Zero Page
		return cpu.opcode0xA6()
	case 0xA8: // TAY Implied
		return cpu.opcode0xA8()
	case 0xAA: // TAX Implied
		return cpu.opcode0xAA()
	case 0xAC: // LDY Absolute
		return cpu.opcode0xAC()
	case 0xAD: // LDA Absolute
		return cpu.opcode0xAD()
	case 0xAE: // LDX Absolute
		return cpu.opcode0xAE()
	case 0xB1: // LDA (oper),Y
		return cpu.opcode0xB1()
	case 0xB4: // LDY Zero Page,X
		return cpu.opcode0xB4()
	case 0xB5: // LDA Zero Page,X
		return cpu.opcode0xB5()
	case 0xB6: // LDX Zero Page,Y
		return cpu.opcode0xB6()
	case 0xB8: // CLV Implied
		return cpu.opcode0xB8()
	case 0xBA: // TSX Implied
		return cpu.opcode0xBA()
	case 0xBC: // LDY Absolute,X
		return cpu.opcode0xBC()
	case 0xC0: // CPY Immediate
		return cpu.opcode0xC0()
	case 0xC1: // CMP (oper,X)
		return cpu.opcode0xC1()
	case 0xC4: // CPY Zero Page
		return cpu.opcode0xC4()
	case 0xC5: // CMP Zero Page
		return cpu.opcode0xC5()
	case 0xC6: // DEC Zero Page
		return cpu.opcode0xC6()
	case 0xC8: // INY Implied
		return cpu.opcode0xC8()
	case 0xCA: // DEX Implied
		return cpu.opcode0xCA()
	case 0xCC: // CPY Absolute
		return cpu.opcode0xCC()
	case 0xCD: // CMP Absolute
		return cpu.opcode0xCD()
	case 0xCE: // DEC Absolute
		return cpu.opcode0xCE()
	case 0xD0: // BNE Relative
		return cpu.opcode0xD0()
	case 0xD1: // CMP (oper),Y
		return cpu.opcode0xD1()
	case 0xD5: // CMP Zero Page,X
		return cpu.opcode0xD5()
	case 0xD6: // DEC Zero Page,X
		return cpu.opcode0xD6()
	case 0xD8: // CLD Implied
		return cpu.opcode0xD8()
	case 0xD9: // CMP Absolute,Y
		return cpu.opcode0xD9()
	case 0xDD: // CMP Absolute,X
		return cpu.opcode0xDD()
	case 0xDE: // DEC Absolute,X
		return cpu.opcode0xDE()
	case 0xE0: // CPX Immediate
		return cpu.opcode0xE0()
	case 0xE1: // SBC (oper,X)
		return cpu.opcode0xE1()
	case 0xE4: // CPX Zero Page
		return cpu.opcode0xE4()
	case 0xE5: // SBC Zero Page
		return cpu.opcode0xE5()
	case 0xE6: // DEC Zero Page
		return cpu.opcode0xE6()
	case 0xE8: // INX Implied
		return cpu.opcode0xE8()
	case 0xEA: // NOP Implied
		return cpu.opcode0xEA()
	case 0xEC: // CPX Absolute
		return cpu.opcode0xEC()
	case 0xED: // SBC Absolute
		return cpu.opcode0xED()
	case 0xEE: // DEC Absolute
		return cpu.opcode0xEE()
	case 0xF0: // BEQ Relative
		return cpu.opcode0xF0()
	case 0xF1: // SBC (oper),Y
		return cpu.opcode0xF1()
	case 0xF5: // SBC Zero Page,X
		return cpu.opcode0xF5()
	case 0xF6: // INC Zero Page,X
		return cpu.opcode0xF6()
	case 0xF8: // SED Implied
		return cpu.opcode0xF8()
	case 0xF9: // SBC Absolute,Y
		return cpu.opcode0xF9()
	case 0xFD: // SBC Absolute,X
		return cpu.opcode0xFD()
	case 0xFE: // INC Absolute,X
		return cpu.opcode0xFE()

	default:
		fmt.Printf("Unknown opcode: %02X\n", opcode)
		return 1, true
	}
}

func (cpu *CPU) fetch() byte {
	value := cpu.Memory[cpu.PC]
	cpu.PC++
	return value
}

func (cpu *CPU) LoadProgram(program []byte, startAddress uint16) {
	copy(cpu.Memory[startAddress:], program)
}

func (cpu *CPU) StoreData(data byte, address uint16) {
	cpu.Memory[address] = data
}

func (cpu *CPU) Run() {
	for {
		opcode := cpu.fetch()
		cpu.Execute(opcode)
	}
}

func (cpu *CPU) RunStep() {
	opcode := cpu.fetch()
	cpu.Execute(opcode)
}

func (cpu *CPU) Reset() {
	cpu.SP = 0xFF   // Stack Pointer starts at 0xFF
	cpu.PC = 0x8000 // Reset vector points to address $FFFC/$FFFD in the ROM cartridge
	cpu.X = 0       // Initialize X register to 0
	cpu.Y = 0       // Initialize Y register to 0
	cpu.A = 0       // Initialize accumulator to 0

	cpu.clearDecimalModeFlag()
	cpu.setIntDisableFlag()

}

func main() {
	cpu := NewCPU()
	program := []byte{
		0xA9, 0x01, // LDA Immediate with value 0x01
		0x65, 0x10, // ADC Zero Page with address 0x10
		0x6D, 0x34, 0x12, // ADC $1234
	} // Example program

	cpu.StoreData(0x01, 0x10) // Store a value of 0x01 at address 0x10
	cpu.StoreData(0x01, 0x1234)

	cpu.LoadProgram(program, 0x8000)
	cpu.Reset()
	cpu.Run()
}
