package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main(){
	args := os.Args
	fmt.Println(args)
	data, err := os.ReadFile(args[1])
    if err != nil {
        fmt.Println("Error reading file:", err)
        os.Exit(1)
    }

    f := string(data)

    
    lines := strings.Split(f, "\n")

    // for _, line := range lines {
    //     fmt.Println(line)
    // }

	symbol_table := map[string]int{
		"R0": 0,
		"R1": 1,
		"R2": 2,
		"R3": 3,
		"R4": 4,
		"R5": 5,
		"R6": 6,
		"R7": 7,
		"R8": 8,
		"R9": 9,
		"R10": 10,
		"R11": 11,
		"R12": 12,
		"R13": 13,
		"R14": 14,
		"R15": 15,
		"SCREEN":16384,
		"KBD":24576,
		"SP":0,
		"LCL":1,
		"ARG":2,
		"THIS":3,
		"THAT":4,
	}

	
	C_INS := map[string]int{
		"NONE": 0b000,
		"JGT" : 0b001,
		"JEQ" : 0b010,
		"JGE" : 0b011,
		"JLT" : 0b100,
		"JNE" : 0b101,
		"JLE" : 0b110,
		"JMP" : 0b111}
	
	C_INS["None"] = 0b000
	symbol_table["R0"] = 0

	first_pass := make([]string,0)
	pc := 0
	for _,line := range lines{
		ws ,_ := regexp.MatchString("//.*|^\\s$",line)
		label,_ := regexp.MatchString("^\\(.+\\)",line)
		
		if !ws && line!="" {
			first_pass = append(first_pass, line)
			if label{
				// fmt.Println(line)
				symbol_table[line[1:len(line)-2]] = pc
			}
			pc++
		}
	}

	// for _,line := range first_pass{
	// 	constant, _ = regexp.MustCompile()
	// }
	
	// defer fmt.Println(symbol_table)


}