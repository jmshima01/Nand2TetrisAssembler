a. James Shima
b. few days
c. This is the code base for project 06 for CSCI410.
It is writen in Go (1.22.0) and has been tested against 
the provided Assembler.sh and .asm programs. There were a lot of bugs
that I had to find such as inc non label vars only based on themselves and
the reserved regs not the ROM labels.

How to run:
`make`
`./Assembler <path/to/file.asm>`
the resulting `<file.hack>` will be written to the same dir as the file path given as the arg 
(same behavior as the provided Assembler.sh)

I also made a testing bash script called `test.sh` that will make, chmod, and diff my results to
the expected `.hack` files provided in the `tests`

My make file also has other options if that is of any help e.g. clean, run, test, etc.

