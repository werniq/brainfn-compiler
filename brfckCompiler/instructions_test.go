package brfnCompiler

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

func setup() {
	pointer = 0
	memory = make([]int, 30000)
	output = strings.Builder{}
	code = ""
	index = 0
	indexMoves = make(map[int]int)
}

func TestDecrementPointerInstruction(t *testing.T) {
	setup()
	instr := DecrementPointerInstruction{}
	pointer = 1
	instr.Execute()
	if pointer != 0 {
		t.Errorf("Expected pointer to be 0, got %d", pointer)
	}
}

func TestIncrementByteInstruction(t *testing.T) {
	setup()
	instr := IncrementByteInstruction{}
	instr.Execute()

	if memory[pointer] != 1 {
		t.Errorf("Expected memory at pointer to be 1, got %d", memory[pointer])
	}

	memory[pointer] = 255
	instr.Execute()
	if memory[pointer] != 0 {
		t.Errorf("Expected memory at pointer to be 0, got %d", memory[pointer])
	}
}

func TestIncrementPointerInstruction(t *testing.T) {
	setup()
	instr := IncrementPointerInstruction{}
	instr.Execute()
	if pointer != 1 {
		t.Errorf("Expected pointer to be 1, got %d", pointer)
	}
}

func TestDecrementByteInstruction(t *testing.T) {
	setup()

	instr := DecrementByteInstruction{}
	instr.Execute()

	if memory[pointer] != 255 {
		t.Errorf("Expected memory at pointer to be 255, got %d", memory[pointer])
	}

	memory[pointer] = 1
	instr.Execute()
	if memory[pointer] != 0 {
		t.Errorf("Expected memory at pointer to be 0, got %d", memory[pointer])
	}
}

func TestOutputInstruction(t *testing.T) {
	setup()
	instr := OutputInstruction{}
	memory[pointer] = 'A'
	instr.Execute()

	if output.String() != "A" {
		t.Errorf("Expected output to be 'A', got '%s'", output.String())
	}
}

func TestLoopInstruction(t *testing.T) {
	setup()
	code = "[]"

	indexMoves = map[int]int{0: 1, 1: 0}
	memory[pointer] = 0
	index = 0

	instr := LoopInstruction{}
	instr.Execute()

	if index != 1 {
		t.Errorf("Expected index to be 1, got %d", index)
	}
	memory[pointer] = 1
	index = 1
	instr.Execute()
	if index != 0 {
		t.Errorf("Expected index to be 0, got %d", index)
	}
}

func TestLoopInstruction_prepareMoves(t *testing.T) {
	setup()
	tests := []struct {
		name    string
		l       LoopInstruction
		program string
		want    map[int]int
		wantErr error
	}{
		{
			name:    "Test 1",
			l:       LoopInstruction{},
			program: "[[[]]]",
			want:    map[int]int{0: 5, 1: 4, 2: 3, 3: 2, 4: 1, 5: 0},
			wantErr: nil,
		},
		{
			name:    "Test 2",
			l:       LoopInstruction{},
			program: "[[]",
			want:    nil,
			wantErr: errors.New("Brainfuck: mismatched parentheses (at indexes: [0])"),
		},
		{
			name:    "Test 3",
			l:       LoopInstruction{},
			program: "[]]",
			want:    nil,
			wantErr: errors.New("Brainfuck: mismatched parentheses (at indexes: 2)"),
		},
	}

	for _, tt := range tests {
		indexMoves = nil
		t.Run(tt.name, func(t *testing.T) {
			err := prepareMoves(tt.program)
			if !reflect.DeepEqual(indexMoves, tt.want) {
				t.Errorf("LoopInstruction.prepareMoves() = %v, want %v", indexMoves, tt.want)
			}
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("LoopInstruction.prepareMoves() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
