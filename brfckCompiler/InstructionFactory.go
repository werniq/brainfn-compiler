package brfnCompiler

// InstructionFactory defines the abstract factory interface
type InstructionFactory interface {
	CreateIncrementPointer() Instruction
	CreateDecrementPointer() Instruction
	CreateIncrementByte() Instruction
	CreateDecrementByte() Instruction
	CreateOutput() Instruction
	CreateLoop() Instruction
}

// BrainfuckInstructionFactory implements the InstructionFactory interface for Brainfuck compiler
type BrainfuckInstructionFactory struct{}

func (f BrainfuckInstructionFactory) CreateIncrementPointer() Instruction {
	return IncrementPointerInstruction{}
}

func (f BrainfuckInstructionFactory) CreateDecrementPointer() Instruction {
	return DecrementPointerInstruction{}
}

func (f BrainfuckInstructionFactory) CreateIncrementByte() Instruction {
	return IncrementByteInstruction{}
}

func (f BrainfuckInstructionFactory) CreateDecrementByte() Instruction {
	return DecrementByteInstruction{}
}

func (f BrainfuckInstructionFactory) CreateOutput() Instruction {
	return OutputInstruction{}
}

func (f BrainfuckInstructionFactory) CreateLoop() Instruction {
	l := LoopInstruction{}
	return l
}
