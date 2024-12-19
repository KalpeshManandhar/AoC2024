package main

import (
	"fmt"
	"os"
	"strings"
)


var targetCache = map[string] int{}


func tryCreate(target string, available [26][]string) int{
	if (target == ""){
		return 1
	}
	
	if val, ok := targetCache[target]; ok{
		return val
	}


	ways := 0
	startingLetter := target[0] - 'a'
	for _, pattern := range available[startingLetter] {
		if (len(pattern) > len(target)){
			continue
		}

		if (pattern != target[0: len(pattern)]){
			continue
		}
		
		nWays := tryCreate(target[len(pattern):], available)
		if (nWays != -1){
			ways += nWays
		}
	}
	
	if ways == 0{
		ways = -1
	}
	targetCache[target] = ways
	return ways
}



func Day19(){
	fmt.Println("--- Day 19: Linen Layout ---")

	file, err := os.Open("./inputs/day19.txt")
	defer file.Close()
	
	if err != nil{
		panic(err)
	}
	info, _ := file.Stat()
	
	bytes := make([]byte, info.Size())

	file.Read(bytes)

	available_required := strings.Split(string(bytes), "\r\n\r\n")

	availablePatterns := strings.Split(available_required[0], ", ")
	requiredPatterns := strings.Split(available_required[1], "\r\n")

	var begin [26][]string

	for _, pattern := range availablePatterns{
		startingLetter := pattern[0] - 'a'
		if (begin[startingLetter] == nil){
			begin[startingLetter] = make([]string, 0)
		}

		begin[startingLetter] = append(begin[startingLetter], pattern)
	}


	count_p1 := 0
	sum_p2 := 0
	for _, required := range requiredPatterns{
		nWays := tryCreate(required, begin)
		
		if nWays != -1{
			count_p1++
			sum_p2 += nWays
		}
	}
	
	fmt.Println("Part 1: ", count_p1)
	fmt.Println("Part 2: ", sum_p2)

}