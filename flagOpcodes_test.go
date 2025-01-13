package nmos6502

import "testing"

func TestSEC(t *testing.T) {

	cpu := NewCPU()
	program := []byte{
		0x38, // SEC (set carry flag) instruction
		0xEA, // NOP (no operation) instruction
	} // Example program

	cpu.LoadProgram(program, 0x8000)
	cpu.Reset()
	cpu.RunStep()

	if (cpu.Status & Carry) == 0 {
		t.Errorf("Expected Carry Flag to be set")
	}
}
