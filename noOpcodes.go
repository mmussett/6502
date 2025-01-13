package nmos6502

/*
NOPs (including DOP, TOP)
Instructions effecting in 'no operations' in various address modes. Operands are ignored.

N	Z	C	I	D	V
-	-	-	-	-	-
opc	addressing	bytes	cycles
1A	implied		1	2
3A	implied		1	2
5A	implied		1	2
7A	implied		1	2
DA	implied		1	2
FA	implied		1	2
80	immediate	2	2
82	immediate	2	2
89	immediate	2	2
C2	immediate	2	2
E2	immediate	2	2
04	zeropage	2	3
44	zeropage	2	3
64	zeropage	2	3
14	zeropage,X	2	4
34	zeropage,X	2	4
54	zeropage,X	2	4
74	zeropage,X	2	4
D4	zeropage,X	2	4
F4	zeropage,X	2	4
0C	absolute	3	4
1C	absolut,X	3	4*
3C	absolut,X	3	4*
5C	absolut,X	3	4*
7C	absolut,X	3	4*
DC	absolut,X	3	4*
FC	absolut,X	3	4*

*/

func (cpu *CPU) opcode0x1A() (byte, bool) { // TAY
	return 2, false
}

func (cpu *CPU) opcode0x3A() (byte, bool) { // TRB
	return 2, false
}

func (cpu *CPU) opcode0x5A() (byte, bool) { // PHA
	return 2, false
}

func (cpu *CPU) opcode0x7A() (byte, bool) { // PLA
	return 2, false
}

func (cpu *CPU) opcode0xDA() (byte, bool) { // PHX
	return 2, false
}

func (cpu *CPU) opcode0xFA() (byte, bool) { // PLX
	return 2, false
}

func (cpu *CPU) opcode0x80() (byte, bool) { // BRA
	return 2, true
}

func (cpu *CPU) opcode0x82() (byte, bool) { // BRA
	return 2, true
}

func (cpu *CPU) opcode0x89() (byte, bool) { // BRA
	return 2, true
}

func (cpu *CPU) opcode0xC2() (byte, bool) { // BRA
	return 2, true
}

func (cpu *CPU) opcode0xE2() (byte, bool) { // BRA
	return 2, true
}

func (cpu *CPU) opcode0x04() (byte, bool) { // NOP
	return 3, false
}

func (cpu *CPU) opcode0x44() (byte, bool) { // NOP
	return 3, false
}

func (cpu *CPU) opcode0x64() (byte, bool) { // NOP
	return 3, false
}

func (cpu *CPU) opcode0x14() (byte, bool) { // NOP
	return 4, false
}

func (cpu *CPU) opcode0x34() (byte, bool) { // NOP
	return 4, false
}

func (cpu *CPU) opcode0x54() (byte, bool) { // NOP
	return 4, false
}

func (cpu *CPU) opcode0x74() (byte, bool) { // NOP
	return 4, false
}

func (cpu *CPU) opcode0xD4() (byte, bool) { // NOP
	return 4, false
}

func (cpu *CPU) opcode0xF4() (byte, bool) { // NOP
	return 4, false
}

func (cpu *CPU) opcode0x0C() (byte, bool) { // NOP
	return 4, false
}

func (cpu *CPU) opcode0x1C() (byte, bool) { // NOP
	return 4, false
}

func (cpu *CPU) opcode0x3C() (byte, bool) { // NOP
	return 4, false
}

func (cpu *CPU) opcode0x5C() (byte, bool) { // NOP
	return 4, false
}

func (cpu *CPU) opcode0x7C() (byte, bool) { // NOP
	return 4, false
}

func (cpu *CPU) opcode0xDC() (byte, bool) { // NOP
	return 4, false
}

func (cpu *CPU) opcode0xFC() (byte, bool) { // NOP
	return 4, false
}

/*
NOP
No Operation

---
N	Z	C	I	D	V
-	-	-	-	-	-
addressing	assembler	opc	bytes	cycles
implied		NOP			EA	1		2
*/
func (cpu *CPU) opcode0xEA() (byte, bool) {
	return 2, false
}
