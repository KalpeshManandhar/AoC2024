package main

import (
	"fmt"
	"os"
	"strings"
)






func Day10(){
	fmt.Println("--- Day 10: Hoof It ---")

	file, err := os.Open("./inputs/day10.txt")
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

	
	// for part 1
	coordinates9 := map[Vec2] map[Vec2]bool{} 
	// for part 2
	trailheadRatings := map[Vec2]int{}


	// explore, basically 4 way flood fill
	var find0 func (pos, initial9 Vec2)
	find0 = func (pos, initial9 Vec2){
		if (grid[pos.y][pos.x] == '0'){
			if _, ok := coordinates9[pos]; !ok{
				coordinates9[pos] = map[Vec2]bool{}
				trailheadRatings[pos] = 0
			}

			if _, ok := coordinates9[pos][initial9]; !ok{
				coordinates9[pos][initial9] = true
			}
			
			trailheadRatings[pos]++
			return 
		}

		left := Vec2{pos.x - 1, pos.y}
		right := Vec2{pos.x + 1, pos.y}
		up := Vec2{pos.x, pos.y - 1}
		down := Vec2{pos.x, pos.y + 1}

		explore := func(neighbour Vec2){
			if (isInsideRange(neighbour.x, neighbour.y) && int(grid[pos.y][pos.x] - grid[neighbour.y][neighbour.x]) == 1){
				find0(neighbour, initial9)
			}
		}

		explore(left)
		explore(right)
		explore(up)
		explore(down)
	}





	for row := range grid{
		for col := range grid[row]{
			if (grid[row][col] == '9'){
				find0(Vec2{col, row}, Vec2{col, row})
			}
		}
	}


	sum_p1 := 0
	sum_p2 := 0

	for _, connected9s := range coordinates9{
		score := len(connected9s)
		sum_p1 += score
	} 

	for _, rating := range trailheadRatings{
		sum_p2 += rating
	}

	fmt.Println("Part 1: ", sum_p1)
	fmt.Println("Part 2: ", sum_p2)



}