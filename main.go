package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func make_A(a string, v uint16) string{
	res := "0"
	return res

}

func make_C(a string) string{
	res := "1"
	return res
}


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

	symbol_table := map[string]uint16{
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

	
	C_INS := map[string]uint8{
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
	pc := uint16(0)
	for _,line := range lines{
		ws ,_ := regexp.MatchString("//.*|^\\s$",line)
		label,_ := regexp.MatchString("^\\(.+\\)",line)
		
		if !ws && line!="" {
			first_pass = append(first_pass, strings.TrimSpace(line))
			if label{
				// fmt.Println(line)
				symbol_table[line[1:len(line)-2]] = pc
			}
			pc++
		}
	}
	
	next_var_availble := uint16(16)
	
	to_bin := make([]string,0)
	// second pass
	for _,line := range first_pass{
		// fmt.Printf("line: %v\n", line)
		// constant, _ := regexp.MatchString("^@\\d+",line)
		variable, _ := regexp.MatchString("^@[\\$_a-zA-Z\\.:]+[\\$_a-zA-Z0-9\\.:]*",line)
		
		if variable{
			// fmt.Println(line)
			symbol_table[line] = next_var_availble
			if next_var_availble+1 == symbol_table["SCREEN"] || next_var_availble+1 == symbol_table["KBD"]{
				
				next_var_availble++
			}
			y := line[1:]
			fmt.Println(y)
			next_var_availble++
		}






	}
	
	


}