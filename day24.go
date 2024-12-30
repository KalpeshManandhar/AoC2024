package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Gate struct{
	Type string 
	input [2]string 
	output string 
}





func Day24(){
	fmt.Println("--- Day 24: Crossed Wires ---")

	file, err := os.Open("./inputs/day24.txt")

	if err != nil{
		panic(err)
	}
	defer file.Close()
	
	info, _ := file.Stat()
	
	bytes := make([]byte, info.Size())

	file.Read(bytes)


	input_connections := strings.Split(string(bytes), "\r\n\r\n")

	inputValues := strings.Split(input_connections[0], "\r\n")

	logicValues := map[string]int{}

	x := uint64(0)
	y := uint64(0)
	outputToGateMap := map[string]Gate{}
	inputToGateMap := map[string][]Gate{}
	
	for _, input := range inputValues{
		regex := regexp.MustCompile(`((?:x|y)\d+): (1|0)`)
		
		match := regex.FindAllStringSubmatch(input, -1)
		val, _ := strconv.Atoi(match[0][2])
		logicValues[match[0][1]] = val

		bitPos, _ := strconv.Atoi(match[0][1][1:])
		if (match[0][1][0] == 'x'){
			x |= uint64(val) << bitPos
		} else if match[0][1][0] == 'y'{
			y |= uint64(val) << bitPos

		}
	}


	connections := strings.Split(input_connections[1], "\r\n")

	for _, connection := range connections{
		regex := regexp.MustCompile(`(\w+) (AND|OR|XOR) (\w+) -> (\w+)`)

		match := regex.FindAllStringSubmatch(connection, -1)
		gate := Gate{}
		gate.input[0] = match[0][1]
		gate.Type = match[0][2]
		gate.input[1] = match[0][3]
		gate.output = match[0][4]


		if _, ok := inputToGateMap[gate.input[0]]; !ok{
			inputToGateMap[gate.input[0]] = make([]Gate, 0)
		}
		if _, ok := inputToGateMap[gate.input[1]]; !ok{
			inputToGateMap[gate.input[1]] = make([]Gate, 0)
		}

		inputToGateMap[gate.input[0]] = append(inputToGateMap[gate.input[0]], gate)
		inputToGateMap[gate.input[1]] = append(inputToGateMap[gate.input[1]], gate)

		outputToGateMap[gate.output] = gate
	}



	// part 1
	var computeGate func(outputWire string)
	computeGate = func(outputWire string){
		gate := outputToGateMap[outputWire]

		if _, ok := logicValues[gate.input[0]]; !ok{
			computeGate(gate.input[0])
		}
		if _, ok := logicValues[gate.input[1]]; !ok{
			computeGate(gate.input[1])
		}
		
		a := logicValues[gate.input[0]]
		b := logicValues[gate.input[1]]

		c := 0
		switch gate.Type {
		case "AND":
			c = a & b
		case "OR":
			c = a | b
		case "XOR":
			c = (a ^ b)&0x1
		}
		logicValues[gate.output] = c

	}


	z := uint64(0)
	for _, gate := range outputToGateMap{
		if (gate.output[0] == 'z'){
			computeGate(gate.output)
			
			value := logicValues[gate.output]
			bitPos, _ := strconv.Atoi(gate.output[1:])
			z |= uint64(value) << bitPos
		}
	}
		
	
	// part 2
	zCarry := "z" + fmt.Sprintf("%02d", len(inputValues)/2)
	
	swappedMaybe := map[string]bool{}
	
	// I lost my patience with the technicalities of checking the structure of the full adder. So looked up a solution.
	for _, gate := range outputToGateMap{
		if gate.output[0] == 'z' && gate.Type != "XOR" && gate.output != zCarry{
			swappedMaybe[gate.output] = true
		}
		
		
		if (gate.Type == "XOR" && 
			gate.input[0][0] != 'x' && gate.input[0][0] != 'y' && 
			gate.input[1][0] != 'x' && gate.input[1][0] != 'y' && 
			gate.output[0] != 'z') {
				
			swappedMaybe[gate.output] = true
		}

		if gate.Type == "AND" && 
			gate.input[0] != "x00"  && gate.input[1] != "x00"{
			
			for _, g := range outputToGateMap{
				if (gate.output == g.input[0] || gate.output == g.input[1]) && g.Type != "OR"{
					swappedMaybe[gate.output] = true
				}
			}
			
		}


		if gate.Type == "XOR"{
			
			for _, g := range outputToGateMap{
				if (gate.output == g.input[0] || gate.output == g.input[1]) && g.Type == "OR"{
					swappedMaybe[gate.output] = true
				}
			}
			
		}

	}

	fmt.Println(swappedMaybe)
	
	if len(swappedMaybe) != 8{
		panic("Oh no")
	}

	swappedWires := make([]string, 0)
	for wire := range swappedMaybe{
		swappedWires = append(swappedWires, wire)
	}

	sort.Strings(swappedWires)
	
	concatenated_p2 := strings.Join(swappedWires, ",")

	fmt.Println("Part 1: ", z)
	fmt.Println("Part 2: ", concatenated_p2)

}