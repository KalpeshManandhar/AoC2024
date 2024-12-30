package main

import (
	"fmt"
	"os"
	"strings"
)




func Day20(){
	fmt.Println("--- Day 20: Race Condition ---")

	file, err := os.Open("./inputs/day20.txt")
	
	if err != nil{
		panic(err)
	}
	defer file.Close()
	info, _ := file.Stat()
	
	bytes := make([]byte, info.Size())

	file.Read(bytes)
	
	grid := strings.Split(string(bytes), "\r\n")
	
	start, end := findStartEnd(grid)

	MIN_TIME_SAVED_P1 := 100
	MIN_TIME_SAVED_P2 := 100
	
	computePathScores := func(grid []string, start, end Vec2) []int{
		m, n := len(grid), len(grid[0])
		scoreGrid := make([]int, m*n)
		for i := range scoreGrid{
			scoreGrid[i] = -1
		}
		
		pos := start
		score := 0
		scoreGrid[pos.y * n + pos.x] = 0

		for pos != end{
			for dir := range directions{
				next := Vec2{pos.x + dir.x, pos.y + dir.y}
				if (grid[next.y][next.x] != '#' && scoreGrid[next.y *n + next.x] == -1){
					score++
					pos = next
					scoreGrid[next.y * n + next.x] = score
					break
				}
			}
		}

		return scoreGrid
	}
	
	// compute scores for each cell in the only path
	scoreGrid := computePathScores(grid, start, end)
	
	// count the number of cheats that would save atleast required time
	findAllCheats := func(grid []string, start, end Vec2, scoreGrid []int, minTimeSaved, cheatTime int) int{
		m, n := len(grid), len(grid[0])
		
		pos := start
		count := 0

		for pos != end{
			nextPos := Vec2{}
			// find the next cell
			for dir := range directions{
				next := Vec2{pos.x + dir.x, pos.y + dir.y}
				
				if (next.x < 0 || next.y < 0 || next.x >= n || next.y >= m){
					continue
				}
				cheatStartScore := scoreGrid[pos.y * n + pos.x]

				if (grid[next.y][next.x] != '#'){
					if (scoreGrid[next.y * n + next.x] == cheatStartScore + 1){
						nextPos = next
					}
					continue
				}
			}
			
			// for each cheat length, find all the cheat combinations
			// cheat end is be given by all the points which are at a manhattan distance of cheat length from current pos
			for cheatLength := 2; cheatLength <= cheatTime; cheatLength++{
				x, y := cheatLength, 0
				
				quadrants := [4]Vec2{
					{1,1},
					{-1,1},
					{-1,-1},
					{1,-1},
				}

				for x >= 0{
					findTimeSaved := func(start, end Vec2) int {
						if (end.x < 0 || end.y < 0 || end.x >= n || end.y >= m){
							return 0
						}

						if grid[end.y][end.x] == '#'{
							return 0
						}
						cheatStartScore := scoreGrid[start.y * n + start.x]
						cheatEndScore := scoreGrid[end.y * n + end.x]
						
						return cheatEndScore - cheatStartScore - cheatLength
					}
					

					endSet := map[Vec2]bool{}

					for _, quadrant := range quadrants{
						cheatEnd := Vec2{pos.x + quadrant.x * x, pos.y + quadrant.y * y}
						
						if _, ok := endSet[cheatEnd]; ok{
							continue
						}
						endSet[cheatEnd] = true
						
						timeSaved := findTimeSaved(pos, cheatEnd)
					
						if timeSaved >= minTimeSaved{
							count++
						}
					}
					
					x--
					y++
				}

			}

			pos = nextPos
		}

		return count
	}
	
	count_p1 := findAllCheats(grid, start, end, scoreGrid, MIN_TIME_SAVED_P1, 2)
	count_p2 := findAllCheats(grid, start, end, scoreGrid, MIN_TIME_SAVED_P2, 20)
	
	fmt.Println("Part 1: ", count_p1)
	fmt.Println("Part 2: ", count_p2)

}