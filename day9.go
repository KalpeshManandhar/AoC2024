package main

import (
	"fmt"
	"os"
	// "strings"
)



func Day9(){
	fmt.Println("--- Day 9: Disk Fragmenter ---")

	file, err := os.Open("./inputs/day9.txt")
	
	if err != nil{
		panic(err)
	}
	defer file.Close()
	info, _ := file.Stat()
	
	bytes := make([]byte, info.Size())

	file.Read(bytes)

	sum := 0
	for i := range bytes{
		bytes[i] -= '0'
		sum += int(bytes[i])
	}
	

	blocks_p1 := make([]int, sum)
	blocks_p2 := make([]int, sum)

	
	pos := 0
	for i := range bytes{
		count := int(bytes[i])
		for j := 0; j < count; j++ {
			if (i%2 == 0){
				blocks_p1[pos + j] = i/2  
			}else{
				blocks_p1[pos + j] = -1  
			}
		}
		pos += count
	}
	
	copy(blocks_p2, blocks_p1)
	
	// for part 1
	{

		l:= 0
		r:= len(blocks_p1)-1
		
		for l < r{
			if blocks_p1[l] == -1 && blocks_p1[r] != -1{
				blocks_p1[l] = blocks_p1[r]
				blocks_p1[r] = -1
			} 
			if blocks_p1[r] == -1{
				r--
			}
			if blocks_p1[l] != -1{
				l++
			}
		}
	}

	// for part 2
	{
		r:= len(blocks_p2)-1
		
		for r >= 0{
			if (blocks_p2[r] != -1){

				rCount := 0
				id := blocks_p2[r]
				for r>=0 && blocks_p2[r] == id{
					r--
					rCount++
				}
				r++

				l:= 0				
				for l < r{
					if (blocks_p2[l] == -1){
						lCount := 0
						for l + lCount < len(blocks_p2) && blocks_p2[l + lCount] == -1{
							lCount++
						}
				
						if lCount >= rCount{
							for i := 0; i < rCount; i++ {
								blocks_p2[l+i] = blocks_p2[r+i];
								blocks_p2[r+i] = -1
							}
							break
						}
						l += lCount
					}
					
					l++
				}
			}
			r--

		}

	}
	
	calcSum := func(blocks []int) int{
		sum := 0

		for i, block := range blocks {
			if (block == -1){
				continue
			}
			sum += i * block
		}

		return sum
	}


	sum_p1 := calcSum(blocks_p1)
	sum_p2 := calcSum(blocks_p2)
	
	
	fmt.Println("Part 1: ", sum_p1)
	fmt.Println("Part 2: ", sum_p2)

	
}