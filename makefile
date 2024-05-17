build:
	@echo "BUILDING BRAINFUNC COMPILER"
	go build -o bin/brfnc.exe

run: build
	@echo "RUNNING BRAINFUNC COMPILER"
	./bin/brfnc.exe

test:
	go test brfckCompiler/... -v