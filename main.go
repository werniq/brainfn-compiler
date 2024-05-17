package main

import brfnCompiler "brfckCompiler/brfckCompiler"

func main() {
	program := ">+++++++++[<++++++++>-]<.>+++++++[<++++>-]<+.+++++++..+++.>>>++++++++[<++++>-]<.>>>++++++++++[<+++++++++>-]<---.<<<<.+++.------.--------.>>+.>++++++++++."

	factory := brfnCompiler.BrainfuckInstructionFactory{}
	brfnCompiler.Compile(program, factory)
}
