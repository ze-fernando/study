name ?= main

all: $(name).o
	ld -o $(name) $(name).o
	@rm -rf *.o


%.o: %.asm
	@rm -rf $(name) 2>/dev/null
	nasm -g -f elf64 $<
