package nmos6502

const (
	Carry        = 0x01
	Zero         = 0x02
	IntDisable   = 0x04
	DecimalMode  = 0x08
	BreakCommand = 0x10
	Overflow     = 0x40
	Negative     = 0x80
)

func (cpu *CPU) setCarryFlag() {
	cpu.Status |= Carry
}

func (cpu *CPU) clearCarryFlag() {
	cpu.Status &^= Carry
}

func (cpu *CPU) setZeroFlag() {
	cpu.Status |= Zero
}

func (cpu *CPU) clearZeroFlag() {
	cpu.Status &^= Zero
}

func (cpu *CPU) clearIntDisableFlag() {
	cpu.Status &^= IntDisable
}

func (cpu *CPU) setIntDisableFlag() {
	cpu.Status |= IntDisable
}

func (cpu *CPU) setDecimalModeFlag() {
	cpu.Status |= DecimalMode
}

func (cpu *CPU) clearDecimalModeFlag() {
	cpu.Status &^= DecimalMode
}

func (cpu *CPU) clearBreakCommandFlag() {
	cpu.Status &^= BreakCommand
}

func (cpu *CPU) setBreakCommandFlag() {
	cpu.Status |= BreakCommand
}

func (cpu *CPU) clearOverflowFlag() {
	cpu.Status &^= Overflow
}

func (cpu *CPU) setOverflowFlag() {
	cpu.Status |= Overflow
}

func (cpu *CPU) clearNegativeFlag() {
	cpu.Status &^= Negative
}

func (cpu *CPU) setNegativeFlag() {
	cpu.Status |= Negative
}

func (cpu *CPU) updateZeroAndNegativeFlags(value byte) {
	if value == 0 {
		cpu.Status |= 0x02 // Set zero flag
	} else {
		cpu.Status &^= 0x02 // Clear zero flag
	}

	if value&0x80 != 0 {
		cpu.Status |= 0x80 // Set negative flag
	} else {
		cpu.Status &^= 0x80 // Clear negative flag
	}
}
