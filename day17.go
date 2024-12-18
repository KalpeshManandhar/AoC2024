package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type CPU struct{
	A, B, C, ip int
	output []string
	instructions []string
} 

func getComboOperand(cpu CPU, operand string) int {
	switch (operand){
		case "0": return 0
		case "1": return 1
		case "2": return 2
		case "3": return 3
		case "4": return cpu.A
		case "5": return cpu.B
		case "6": return cpu.C
		case "7": 
		default:
			panic("Invalid combo operand")
	}
	return 0
}


func simulateUntilOutput(cpu CPU) int {
	for cpu.ip < len(cpu.instructions){
		switch cpu.instructions[cpu.ip] {
		case "0": // ADV - Divide A by 2^(combo operand) 
			operand := getComboOperand(cpu, cpu.instructions[cpu.ip + 1])
			cpu.A = cpu.A / (0x1 << operand)
			
		case "1": // BXL - Bitwise XOR of B with literal operand
			operand, _ := strconv.Atoi(cpu.instructions[cpu.ip + 1])
			cpu.B = cpu.B ^ operand
			
		case "2": // BST - Store (combo operand)%8 into B
			operand := getComboOperand(cpu, cpu.instructions[cpu.ip + 1])
			cpu.B = operand % 8
			
		case "3": // JNZ - Jump to literal operand if A != 0
			operand, _ := strconv.Atoi(cpu.instructions[cpu.ip + 1])
			if cpu.A != 0{
				cpu.ip = operand
				continue
			}

		case "4": // BXC - Bitwise xor B and C, store into B
			cpu.B = cpu.B ^ cpu.C
	
		case "5": // OUT - Output the value (combo operand)%8
			operand := getComboOperand(cpu, cpu.instructions[cpu.ip + 1])
			res := operand % 8
			return res

		case "6": // BDV - Divide A by 2^(combo operand), store into B
			operand := getComboOperand(cpu, cpu.instructions[cpu.ip + 1])
			cpu.B = cpu.A / (0x1 << operand)

		case "7": // CDV - Divide A by 2^(combo operand), store into C
			operand := getComboOperand(cpu, cpu.instructions[cpu.ip + 1])
			cpu.C = cpu.A / (0x1 << operand)
		}

		cpu.ip += 2

	}

	return -1
}


/*
For each loop iteration, get a value Aend which is a valid value for register A at the end of that iteration
It assumes that the ADV instruction has combo operand of 0-3 aka no register values.
And also only one ADV instruction in the loop.
And jump goes back to ip=0.
And a lot more things.
So, the ADV instruction basically does this
Aend <- Astart/2^operand

So, the range of values for Astart is
[Aend * (2^operand), Aend * (2^operand) + (2^operand - 1)]
For each of these possible values of Astart, check whether it is valid by simulating until the output.
If the output is the same as the required opcode, then the Astart value works for this iteration of the loop,
and that value is the Aend value at the end of the previous iteration, so try assigning backwards to the previous iterations.
else just discard.
*/

func backwardsAssign(cpu CPU, Aend int, instructionNo int) int {
	if (instructionNo < 0){
		return Aend
	}

	var Astart_left int
	var Astart_right int
	
	cpu.ip = len(cpu.instructions) - 2
	for cpu.ip >= 0{
		// the only instruction we are concerned with is ADV as it is the only one that modifies A
		if (cpu.instructions[cpu.ip] == "0"){
			operand := getComboOperand(cpu, cpu.instructions[cpu.ip + 1])
			Astart_left = (1 << operand) * Aend 
			Astart_right = (1 << operand) * Aend + (1 << operand) - 1	
		}

		cpu.ip -= 2
	}
	
	for Astart:= Astart_left; Astart <= Astart_right; Astart++{
		cpu.A = Astart
		cpu.ip = 0
		output := simulateUntilOutput(cpu)
		
		expectedOutput, _ := strconv.Atoi(cpu.instructions[instructionNo])

		if (output != expectedOutput){
			continue
		}

		ret := backwardsAssign(cpu, Astart, instructionNo - 1)
		
		if ret != -1{
			return ret
		}
	}

	return -1
}





func Day17(){
	fmt.Println("--- Day 17: Chronospatial Computer ---")

	file, err := os.Open("./inputs/day17.txt")
	if err != nil{
		defer file.Close()
	}
	
	if err != nil{
		panic(err)
	}
	info, _ := file.Stat()
	
	bytes := make([]byte, info.Size())

	file.Read(bytes)

	regex := regexp.MustCompile(`Register A: (\d+)\s*Register B: (\d+)\s*Register C: (\d+)\s*Program: (.+)`)

	matches	:= regex.FindAllStringSubmatch(string(bytes), -1)
	
	cpu := CPU{}
	
	cpu.A, _ = strconv.Atoi(matches[0][1])
	cpu.B, _ = strconv.Atoi(matches[0][2])
	cpu.C, _ = strconv.Atoi(matches[0][3])
	cpu.ip = 0
	cpu.output = make([]string, 0)
	cpu.instructions = strings.Split(matches[0][4], ",")	

	
	// part 1
	for cpu.ip < len(cpu.instructions){
		switch cpu.instructions[cpu.ip] {
		case "0": // ADV - Divide A by 2^(combo operand) 
			operand := getComboOperand(cpu, cpu.instructions[cpu.ip + 1])
			cpu.A = cpu.A / (0x1 << operand)
			
		case "1": // BXL - Bitwise XOR of B with literal operand
			operand, _ := strconv.Atoi(cpu.instructions[cpu.ip + 1])
			cpu.B = cpu.B ^ operand
			
		case "2": // BST - Store (combo operand)%8 into B
			operand := getComboOperand(cpu, cpu.instructions[cpu.ip + 1])
			cpu.B = operand % 8
			
		case "3": // JNZ - Jump to literal operand if A != 0
			operand, _ := strconv.Atoi(cpu.instructions[cpu.ip + 1])
			if cpu.A != 0{
				cpu.ip = operand
				continue
			}

		case "4": // BXC - Bitwise xor B and C, store into B
			cpu.B = cpu.B ^ cpu.C
	
		case "5": // OUT - Output the value (combo operand)%8
			operand := getComboOperand(cpu, cpu.instructions[cpu.ip + 1])
			res := operand % 8
			cpu.output = append(cpu.output, fmt.Sprint(res))

		case "6": // BDV - Divide A by 2^(combo operand), store into B
			operand := getComboOperand(cpu, cpu.instructions[cpu.ip + 1])
			cpu.B = cpu.A / (0x1 << operand)

		case "7": // CDV - Divide A by 2^(combo operand), store into C
			operand := getComboOperand(cpu, cpu.instructions[cpu.ip + 1])
			cpu.C = cpu.A / (0x1 << operand)
		}

		cpu.ip += 2

	}
	
	outputString := strings.Join(cpu.output, ",")
	fmt.Println("Part 1: ", outputString)
	
	
	// part 2
	cpu.ip = 0
	cpu.output = make([]string, 0)
	
	aStartValForQuine := backwardsAssign(cpu, 0, len(cpu.instructions) - 1)

	fmt.Println("Part 2: ", aStartValForQuine)


}