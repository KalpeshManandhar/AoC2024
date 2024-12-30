package main

import (
	"fmt"
	"os"
	"strings"
)

func Day25(){
	fmt.Println("--- Day 25: Code Chronicle ---")

	file, err := os.Open("./inputs/day25.txt")

	if err != nil{
		panic(err)
	}
	defer file.Close()
	
	info, _ := file.Stat()
	
	bytes := make([]byte, info.Size())

	file.Read(bytes)

	locksKeys := strings.Split(string(bytes), "\r\n\r\n")
	
	locks := make([][]int, 0)
	keys := make([][]int, 0)

	const GRID_HEIGHT = 5

	for _, lockOrKey := range locksKeys{
		grid := strings.Split(lockOrKey, "\r\n")
		counts := make([]int, len(grid[0]))

		for i := 1; i < len(grid)-1; i++{
			for j := range grid[i]{
				if grid[i][j] == '#'{
					counts[j]++ 
				}
			}
		}
		
		if grid[0][0] == '#'{
			locks = append(locks, counts)
		} else if grid[0][0] == '.'{
			keys = append(keys, counts)
		}
	} 
	
	fmt.Println("Locks: ")
	fmt.Println(locks)
	fmt.Println("Keys: ")
	fmt.Println(keys)

	
	count_p1 := 0
	for _, lock := range locks{
		for _, key := range keys{
			willFit := true

			for i := range len(key){
				if (lock[i] + key[i] > GRID_HEIGHT){
					willFit = false
					break
				}
			}

			if willFit{
				count_p1++
			}
		}
	}

	fmt.Println("Part 1: ", count_p1)
	fmt.Println("Part 2: Done!")

}