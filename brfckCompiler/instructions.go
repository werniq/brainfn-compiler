package brfnCompiler

import (
	"errors"
	"fmt"
)

type Instruction interface {
	Execute()
}

type IncrementPointerInstruction struct{}

func (i IncrementPointerInstruction) Execute() {
	pointer += 1
}

type DecrementPointerInstruction struct{}

func (i DecrementPointerInstruction) Execute() {
	pointer -= 1
}

type IncrementByteInstruction struct{}

func (i IncrementByteInstruction) Execute() {
	memory[pointer] = memory[pointer] + 1
	if memory[pointer] == 256 {
		memory[pointer] = 0
	}
}

type DecrementByteInstruction struct{}

func (i DecrementByteInstruction) Execute() {
	memory[pointer] = memory[pointer] - 1
	if memory[pointer] == -1 {
		memory[pointer] = 255
	}
}

type OutputInstruction struct{}

func (i OutputInstruction) Execute() {
	fmt.Printf("%c", rune(memory[pointer]))
	output.WriteRune(rune(memory[pointer]))
}

type LoopInstruction struct{}

func (l LoopInstruction) Execute() {

	chars := []rune(code)

	if chars[index] == '[' {
		if val := memory[pointer]; val == 0 {
			index = indexMoves[index]
		}
	} else if chars[index] == ']' {
		if val := memory[pointer]; val != 0 {
			index = indexMoves[index]
		}
	}
}

func prepareMoves(program string) error {
	lBraces := make([]int, 0)
	res := make(map[int]int)

	for i, command := range program {
		if command == '[' {
			lBraces = append(lBraces, i)
		} else if command == ']' {
			if len(lBraces) == 0 {
				return errors.New(fmt.Sprintf("Brainfuck: mismatched parentheses (at indexes: %d)", i))
			}

			lBraceIndex := lBraces[len(lBraces)-1]
			lBraces = lBraces[:len(lBraces)-1]
			res[lBraceIndex] = i
			res[i] = lBraceIndex
		}
	}

	if len(lBraces) > 0 {
		return errors.New(fmt.Sprintf("Brainfuck: mismatched parentheses (at indexes: %v)", lBraces))
	}
	indexMoves = res

	return nil
}
