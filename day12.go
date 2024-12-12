package main

import (
	"fmt"
	"os"
	"strings"
)

type Info struct{
	area int
	perimeter int
}


func Day12(){
	fmt.Println("--- Day 12: Garden Groups ---")

	file, err := os.Open("./inputs/day12.txt")
	defer file.Close()
	
	if err != nil{
		panic(err)
	}
	info, _ := file.Stat()
	
	bytes := make([]byte, info.Size())

	file.Read(bytes)
	grid := strings.Split(string(bytes), "\r\n")
	m, n := len(grid), len(grid[0])
	
	isInsideRange := func (x, y int) bool{
		return x >= 0 && x < n && y >= 0 && y < m
	}

	regionMap := make([]int, m * n)
	regionsCount := 0

	var visit func(i, j int, info Info, value byte) Info 
	visit = func(i, j int, info Info, value byte) Info{
		if (!isInsideRange(j, i)){
			return Info{area: info.area, perimeter: info.perimeter + 1} 
		}
		
		if (grid[i][j] != value){
			return Info{area: info.area, perimeter: info.perimeter + 1} 
		}

		if (regionMap[i * n + j] != 0){
			return info
		}

		
		regionMap[i * n + j] = regionsCount
		info.area++

		info = visit(i + 1, j, info, value)
		info = visit(i - 1, j, info, value)
		info = visit(i, j + 1, info, value)
		info = visit(i, j - 1, info, value)
		
		return info
	}


	countEdges := func(regionNo int) int{
		nEdges := 0
		
		for i := range m{
			up := -2
			down := -2

			for j := range n{
				if regionMap[i*n+j] != regionNo{
					continue
				}
				
				if (!isInsideRange(i-1, j) || regionMap[(i-1)*n+j] != regionNo){
					if (up < j-1){
						nEdges++
					}
					up = j
				}

				if (!isInsideRange(i+1, j) || regionMap[(i+1)*n+j] != regionNo){
					if (down < j-1){
						nEdges++
					}
					down = j
				}
			}
		}

		for j := range n{
			left := -2
			right := -2

			for i := range m{
				if regionMap[i*n+j] != regionNo{
					continue
				}
				
				if (!isInsideRange(i, j-1) || regionMap[i*n+(j-1)] != regionNo){
					if (left < i-1){
						nEdges++
					}
					left = i
				}

				if (!isInsideRange(i, j+1) || regionMap[i*n+(j+1)] != regionNo){
					if (right < i-1){
						nEdges++
					}
					right = i
				}
			}
		}


		return nEdges
	}

	
	sum_p1 := 0
	sum_p2 := 0

	for i := range m{
		for j := range n{
			if (regionMap[i*n+j] != 0){
				continue
			}
			
			regionsCount++
			info := visit(i, j, Info{0,0}, grid[i][j])
			nEdges := countEdges(regionsCount)

			sum_p1 += (info.area * info.perimeter)
			sum_p2 += (info.area * nEdges)
		}
	}

	fmt.Println("Part 1: ", sum_p1)
	fmt.Println("Part 2: ", sum_p2)
	

}