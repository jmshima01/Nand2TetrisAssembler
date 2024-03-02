package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"strconv"
)

type C_INSTURCTION struct{
	a string;
	comp string;
	dest string;
	jmp string;

}

func get_availible_addr(ram map[string]uint16, s string) uint16{
	_,ok := ram[s]
	if ok{
		return ram[s]
	}
	vals := make([]uint16,0)
	for _,v := range ram{
		vals = append(vals, v)
	}
	
	if len(vals) == 65536{
		fmt.Println("Error: memory full overwrite @ addr 16")
		fmt.Printf("ram: %v\n", ram)
		return uint16(16)
	}

	init := uint16(16)
	found := true
	for{
		found = false
		for _,v := range vals{
			if v == init{
				found = true
			}
		}
		if !found{
			break
		}
		init++

	}
	return init
}

func bin_A(a string, m map[string]uint16, isVar bool) string{
	res := "0"
	val := int64(0)
	if isVar{ val = int64(m[a]) } else { val,_ = strconv.ParseInt(a,10,64)}
	
	b := strconv.FormatInt(val,2)
	return res+fmt.Sprintf("%015s", b)
}

func (self C_INSTURCTION) bin_C() string {
	res := "111"
	return res + self.a + self.comp + self.dest + self.jmp 
}

func make_C(line string) C_INSTURCTION{
	l := strings.Split(line, "=")
	des := ""
	if len(l) > 1{
		des = l[0]
	}
	dest := "000"
	comp := "101010"
	jmp := "000"
	c := ""
	a := ""
	j := ""
	switch des{
		case "M":
			dest = "001"
		case "D":
			dest = "010"
		case "DM":
			dest = "011"
		case "A":
			dest = "100"
		case "AM":
			dest = "101"
		case "AD":
			dest = "110"
		case "ADM":
			dest = "111"
		default:
			dest = "000"
	}
	if des==""{
		t := strings.Split(line,";")
		if len(t) == 1{
			comp = t[0] 
		} else {
			comp = t[0]
			jmp = t[1]
		}
	} else{
		t := strings.Split(l[1],";")
		if len(t) == 1{
			comp = t[0] 
		} else {
			comp = t[0]
			jmp = t[1]
		}

	}
	switch comp{
		case "0":
			c="101010"
			a="0"
		case "1":
			c="111111"
			a="0"
		case "-1":
			c="111010"
			a="0"
		case "D":
			c="001100"
			a="0"
		case "A":
			c="110000"
			a="0"
		case "!D":
			c="001101"
			a="0"
		case "!A":
			c="110001"
			a="0"
		case "-D":
			c="001111"
			a="0"
		case "-A":
			c="110011"
			a="0"
		case "D+1":
			c="011111"
			a="0"
		case "A+1":
			c="110111"
			a="0"
		case "D-1":
			c="001110"
			a="0"
		case "A-1":
			c="110010"
			a="0"
		case "D+A":
			c="000010"
			a="0"
		case "D-A":
			c="010011"
			a="0"
		case "A-D":
			c="000111"
			a="0"
		case "D&A":
			c="000000"
			a="0"
		case "D|A":
			c="010101"
			a="0"
		
		case "M":
			c="110000"
			a="1"
		case "!M":
			c="110001"
			a="1"
		case "-M":
			c="110011"
			a="1"
		case "M+1":
			c="110111"
			a="1"
		case "M-1":
			c="110010"
			a="1"
		case "D+M":
			c="000010"
			a="1"
		case "D-M":
			c="010011"
			a="1"
		case "M-D":
			c="000111"
			a="1"
		case "D&M":
			c="000000"
			a="1"
		case "D|M":
			c="010101"
			a="1"
		default:
			c="101010"
	}

	switch jmp{
		case "":
			j = "000"
		case "JGT":
			j = "001"
		case "JEQ":
			j = "010"
		case "JGE":
			j = "011"
		case "JLT":
			j = "100"
		case "JNE":
			j = "101"
		case "JLE":
			j = "110"
		case "JMP":
			j = "111"
		default:
			j = "000"
	}

	return C_INSTURCTION{a:a,comp:c,dest: dest,jmp: j}
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
				continue
				
				
			}
			pc++
		}
	}

	for i,v := range first_pass{
		fmt.Println(v,"    ",i)
	}
	fmt.Println("===============")
	// addr_availble := get_availible_addr(symbol_table)
	
	// to_bin := make([]string,0)
	// second pass
	curr_A := bin_A("R0",symbol_table,true)
	assembly := make([]string,0)
	for _,line := range first_pass{
		// fmt.Printf("line: %v\n", line)
		constant, _ := regexp.MatchString("^@\\d+",line)
		variable, _ := regexp.MatchString("^@[\\$_a-zA-Z\\.:]+[\\$_a-zA-Z0-9\\.:]*",line)
		c_inst, _ := regexp.MatchString("(^(D|M|A|DM|AM|ADM|AD)=(0|1|-1|D|A|!D|!A|-D|-A|D+1|A+1|D-1|A-1|D+A|D-A|A-D|D&A|D\\|A|M|!M|-M|M+1|M-1|D+M|M-D|D-M|D&M|D\\|M);(JNE|JEQ|JMP|JLT|JLE|JGT|JGE))|((0|1|-1|D|A|!D|!A|-D|-A|D+1|A+1|D-1|A-1|D+A|D-A|A-D|D&A|D\\|A|M|!M|-M|M+1|M-1|D+M|M-D|D-M|D&M|D\\|M);(JNE|JEQ|JMP|JLT|JLE|JGT|JGE))|(^(D|M|A|DM|AM|ADM|AD)=(0|1|-1|D|A|!D|!A|-D|-A|D+1|A+1|D-1|A-1|D+A|D-A|A-D|D&A|D\\|A|M|!M|-M|M+1|M-1|D+M|M-D|D-M|D&M|D\\|M))",line)
		label, _ := regexp.MatchString("^\\([\\$_a-zA-Z\\.:]+[\\$_a-zA-Z0-9\\.:]*\\)",line)
		
		if variable{
			symbol_table[line[1:]] = get_availible_addr(symbol_table,line[1:])
			curr_A = bin_A(line[1:],symbol_table,true)
			assembly= append(assembly, curr_A)
			fmt.Println("var: ", curr_A," : ",line)
		} else if constant{
			
			curr_A = bin_A(line[1:],symbol_table,false)
			assembly = append(assembly, curr_A)
			fmt.Println("const: ", curr_A," : ",line)

		} else if c_inst{
			// s = make_C(line)
			assembly = append(assembly, make_C(line).bin_C())
			fmt.Println(line,":",make_C(line).bin_C())

		} else if label{
			continue

		} else{
			fmt.Println("Synatx error:",line)
			os.Exit(1)
		}
	}

	outdata := make([]byte, 0)
	l := len(assembly)-1
	fmt.Printf("\n======\n\n")
	for i,v := range assembly{
		fmt.Println(v)
		for _,x := range v{
			outdata = append(outdata, byte(x))
		}
		if i < l { outdata = append(outdata, '\n') }
	}
	fmt.Println()
	fmt.Println(symbol_table)

	// name := strings.Split(args[1],".")[0]+".hack"
	name:="out1.hack"
	err = os.WriteFile(name,outdata,0644)
	if err!=nil{
		fmt.Println("Error: writing to file", err)
	}
}