package brfnCompiler

import "strings"

var (
	pointer    = 0
	memory     = make([]int, 3000)
	index      int
	code       string
	indexMoves map[int]int
	output     = strings.Builder{}
)

func Compile(program string, factory InstructionFactory) string {
	prepareMoves(program)
	code = program

	instructions := map[rune]Instruction{
		'>': factory.CreateIncrementPointer(),
		'<': factory.CreateDecrementPointer(),
		'+': factory.CreateIncrementByte(),
		'-': factory.CreateDecrementByte(),
		'.': factory.CreateOutput(),
		'[': factory.CreateLoop(),
		']': factory.CreateLoop(),
	}
	chars := []rune(program)

	for index = 0; index <= len(program)-1; index++ {
		instruction := instructions[chars[index]]
		instruction.Execute()
	}

	return output.String()
}
