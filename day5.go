package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)


func isValidSequence(nums []int, nodes map[int][]int) bool{
			
	current := -1
	for _, val := range nums{		
		if (current != -1){
			canGo := func() bool{
				for _, toNode:= range nodes[current]{
					if toNode == val{
						return true
					}
				}
				return false
			}()

			if (!canGo){
				return false
			}
		}

		current = val
	}


	return true
}





func Day5(){
	fmt.Println("--- Day 5: Print Queue ---")

	file, err := os.Open("./inputs/day5.txt")
	
	if err != nil{
		panic(err)
	}
	defer file.Close()
	info, _ := file.Stat()
	
	bytes := make([]byte, info.Size())

	file.Read(bytes)
	lines := strings.Split(string(bytes), "\r\n")

	nodes := map[int][]int{}
	
	p1_sum := 0
	p2_sum := 0

	for _, line := range lines{
		
		regex := regexp.MustCompile(`(\d+)\|(\d+)`)
		
		// create the graph
		if regex.Match([]byte(line)){	
			nums := strings.Split(line, "|")
			from, _ := strconv.Atoi(nums[0])
			to, _ := strconv.Atoi(nums[1])
			
			if _, ok := nodes[from]; !ok{
				nodes[from] = make([]int, 0)
			}
			if _, ok := nodes[to]; !ok{
				nodes[to] = make([]int, 0)
			}
			
			nodes[from] = append(nodes[from], to)
		
		}else{ // check the orderings
			numStrings := strings.Split(line, ",")
			nums := make([]int, len(numStrings))
			
			for i, numStr := range numStrings{
				val, _ := strconv.Atoi(numStr)
				nums[i] = val
			}

			// if is valid then for part 1
			if isValidSequence(nums, nodes){
				p1_sum += nums[len(nums)/2]
			}else{
			// else for part 2, reorder and count
				sort.Slice(nums, 
					func(i, j int) bool{
						for _, toNode:= range nodes[nums[i]]{
							if toNode == nums[j]{
								return true
							}
						}
						return false
					})
				
				p2_sum += nums[len(nums)/2]
			} 

			
		}
			
	}


	fmt.Println("Part 1: ", p1_sum)
	fmt.Println("Part 2: ", p2_sum)
	

}