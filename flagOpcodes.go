package nmos6502

/*
Flag Instructions

CLC clear carry
CLD clear decimal (BCD arithmetics disabled)
CLI clear interrupt disable
CLV clear overflow
SEC set carry
SED set decimal (BCD arithmetics enabled)
SEI set interrupt disable
*/

func (cpu *CPU) opcode0x38() (byte, bool) { // SEC - Set Carry Flag
	cpu.setCarryFlag()
	return 1, false
}

func (cpu *CPU) opcode0x78() (byte, bool) { // SEI - Set Interrupt Disable
	cpu.setIntDisableFlag()
	return 1, false
}

func (cpu *CPU) opcode0xF8() (byte, bool) { // SED - Set Decimal Mode
	cpu.setDecimalModeFlag()
	return 1, false
}

/*
Clear Carry Flag

0 -> C
N	Z	C	I	D	V
-	-	0	-	-	-
addressing	assembler	opc	bytes	cycles
implied	CLC	18	1	2
*/
func (cpu *CPU) opcode0x18() (byte, bool) { // CLC - Clear Carry Flag
	cpu.clearCarryFlag()
	return 1, false
}

func (cpu *CPU) opcode0x58() (byte, bool) { // CLI - Clear Interrupt Disable
	cpu.clearIntDisableFlag()
	return 1, false
}

func (cpu *CPU) opcode0xD8() (byte, bool) { // CLD - Clear Decimal Mode
	cpu.clearDecimalModeFlag()
	return 1, false
}

func (cpu *CPU) opcode0xB8() (byte, bool) { // CLV - Clear Overflow Flag
	cpu.clearOverflowFlag()
	return 1, false
}
