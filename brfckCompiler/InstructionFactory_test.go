// compiler_test.go

package brfnCompiler

import (
	"testing"
)

func TestBrainfuckInstructionFactory_CreateIncrementPointer(t *testing.T) {
	factory := BrainfuckInstructionFactory{}
	instruction := factory.CreateIncrementPointer()

	if _, ok := instruction.(IncrementPointerInstruction); !ok {
		t.Errorf("Expected IncrementPointerInstruction, got %T", instruction)
	}
}

func TestBrainfuckInstructionFactory_CreateDecrementPointer(t *testing.T) {
	factory := BrainfuckInstructionFactory{}
	instruction := factory.CreateDecrementPointer()

	if _, ok := instruction.(DecrementPointerInstruction); !ok {
		t.Errorf("Expected DecrementPointerInstruction, got %T", instruction)
	}
}

func TestBrainfuckInstructionFactory_CreateIncrementByte(t *testing.T) {
	factory := BrainfuckInstructionFactory{}
	instruction := factory.CreateIncrementByte()

	if _, ok := instruction.(IncrementByteInstruction); !ok {
		t.Errorf("Expected IncrementByteInstruction, got %T", instruction)
	}
}

func TestBrainfuckInstructionFactory_CreateDecrementByte(t *testing.T) {
	factory := BrainfuckInstructionFactory{}
	instruction := factory.CreateDecrementByte()

	if _, ok := instruction.(DecrementByteInstruction); !ok {
		t.Errorf("Expected DecrementByteInstruction, got %T", instruction)
	}
}

func TestBrainfuckInstructionFactory_CreateOutput(t *testing.T) {
	factory := BrainfuckInstructionFactory{}
	instruction := factory.CreateOutput()

	if _, ok := instruction.(OutputInstruction); !ok {
		t.Errorf("Expected OutputInstruction, got %T", instruction)
	}
}

func TestBrainfuckInstructionFactory_CreateLoop(t *testing.T) {
	factory := BrainfuckInstructionFactory{}
	instruction := factory.CreateLoop()

	if _, ok := instruction.(LoopInstruction); !ok {
		t.Errorf("Expected OutputInstruction, got %T", instruction)
	}
}
