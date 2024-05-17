package brfnCompiler

import (
	"reflect"
	"testing"
)

func TestCompile(t *testing.T) {
	program := ">+++++++++[<++++++++>-]<.>+++++++[<++++>-]<+.+++++++..+++.>>>++++++++[<++++>-]<.>>>++++++++++[<+++++++++>-]<---.<<<<.+++.------.--------.>>+.>++++++++++."

	factory := BrainfuckInstructionFactory{}
	result := Compile(program, factory)

	expected := "Hello World!\n"

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Compile() = %v, want %v", result, expected)
	}
}
