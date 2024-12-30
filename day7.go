package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)


func checkIfValid_AddMul(target int, nums []int) bool {
	if len(nums) == 1{
		return target == nums[0]
	}

	if (target % nums[len(nums)-1] == 0){
		mulRes := target / nums[len(nums)-1]
		if (checkIfValid_AddMul(mulRes, nums[:len(nums)-1])){
			return true
		}
	}
	
	addRes := target - nums[len(nums)-1]
	
	return checkIfValid_AddMul(addRes, nums[:len(nums)-1])
}

func checkIfValid_AddMulConcat(target int, nums []int) bool {
	if len(nums) == 1{
		return target == nums[0]
	}

	concatRes := target
	num := nums[len(nums)-1]

	concatFlag := true
	for (num > 0){
		remT := concatRes % 10
		remN := num % 10
		if (remN != remT){
			concatFlag = false
		}
		concatRes /= 10 
		num /= 10 
	}

	if concatFlag{
		if (checkIfValid_AddMulConcat(concatRes, nums[:len(nums)-1])){
			return true
		}
	}


	if (target % nums[len(nums)-1] == 0){
		mulRes := target / nums[len(nums)-1]
		if (checkIfValid_AddMulConcat(mulRes, nums[:len(nums)-1])){
			return true
		}
	}
	
	addRes := target - nums[len(nums)-1]

	return checkIfValid_AddMulConcat(addRes, nums[:len(nums)-1])
}



func Day7(){
	fmt.Println("--- Day 7: Bridge Repair ---")

	file, err := os.Open("./inputs/day7.txt")
	
	if err != nil{
		panic(err)
	}
	defer file.Close()
	info, _ := file.Stat()
	
	bytes := make([]byte, info.Size())

	file.Read(bytes)
	lines := strings.Split(string(bytes), "\r\n")

	sum_p1 := 0
	sum_p2 := 0

	for _, line := range lines{
		regex := regexp.MustCompile(`(\d+): (.+)`)
		matches := regex.FindAllStringSubmatch(line, -1)
		
		for _, match := range matches{
			target, _ := strconv.Atoi(match[1])
			numStrs := strings.Split(match[2], " ")
			nums := make([]int, len(numStrs))
			
			for i, numStr := range numStrs{
				nums[i], _ = strconv.Atoi(numStr)
			}

			if (checkIfValid_AddMul(target, nums)){
				sum_p1 += target
			}
			if (checkIfValid_AddMulConcat(target, nums)){
				sum_p2 += target
			}
		}
		
	}
	fmt.Println("Part 1: ", sum_p1)
	fmt.Println("Part 2: ", sum_p2)


}