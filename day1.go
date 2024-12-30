package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)


func Day1() {
	fmt.Println("--- Day 1: Historian Hysteria ---")

	file, err := os.Open("./inputs/day1.txt")
	
	if err != nil{
		panic(err)
	}
	defer file.Close()

	info, _ := file.Stat()
	
	bytes := make([]byte, info.Size())

	file.Read(bytes)
	
	lines := strings.Split(string(bytes), "\r\n")
	
	var numbers [2][]int

	numbers[0] = make([]int, len(lines))
	numbers[1] = make([]int, len(lines))


	for i, line := range lines{
		nums := strings.Split(line, "   ")

		numbers[0][i], _ = strconv.Atoi(nums[0])
		numbers[1][i], _ = strconv.Atoi(nums[1])

	}
	
	// sort the numbers 
	sort.Slice(numbers[0], func(i, j int) bool {return numbers[0][i] < numbers[0][j]})
	sort.Slice(numbers[1], func(i, j int) bool {return numbers[1][i] < numbers[1][j]})

	
	// get diff sum for part 1
	sum := 0
	for i := 0; i<len(lines); i++{
		sum += int(math.Abs(float64(numbers[0][i] - numbers[1][i])))
	}
	
	fmt.Println("Part 1:", sum)
	
	// get sum of (num x occurences) for pt 2 
	similarScore := 0
	for i := 0; i<len(lines); i++{
		occurences := 0

		for j:=0; j<len(lines) && numbers[1][j] <= numbers[0][i]; j++{
			if numbers[0][i] == numbers[1][j]{
				occurences++
			}
		} 
			
		similarScore += (numbers[0][i] * occurences)
	}
		
	fmt.Println("Part 2:", similarScore)

}
