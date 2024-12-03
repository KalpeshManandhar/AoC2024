package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)


func Day3(){
	fmt.Println("--- Day 3: Mull It Over ---")

	file, err := os.Open("./inputs/day3.txt")
	defer file.Close()
	
	if err != nil{
		panic(err)
	}
	info, _ := file.Stat()
	
	bytes := make([]byte, info.Size())

	file.Read(bytes)
	
	regex := regexp.MustCompile(`mul\(([\d]+),([\d]+)\)|do\(\)|don't\(\)`)

	matches := regex.FindAllSubmatch(bytes, -1)
	
	sum := 0
	
	activeSum := 0
	active := true

	for i := 0;i < len(matches); i++{
		if (string(matches[i][0]) == "do()"){
			active = true
			continue
		} else if (string(matches[i][0]) == "don't()"){
			active = false
			continue
		}

		a, _ := strconv.Atoi(string(matches[i][1]))
		b, _ := strconv.Atoi(string(matches[i][2]))
	
		sum += (a * b)
		
		if (active){
			activeSum += (a*b)
		}

	}

	fmt.Println("Part 1: ", sum)
	fmt.Println("Part 2: ", activeSum)

}