package main

import (
	"fmt"
	"os"
	"strings"
)


func findRobotPos(grid [][]byte) Vec2{
	for i, row := range grid{
		for j, tile := range row{
			if (tile == '@'){
				return Vec2{j, i}
			}
		}
	}

	return Vec2{-1,-1}
}


func isConnectedToWall(grid [][]byte, pos Vec2, dir Vec2) bool {
	for grid[pos.y][pos.x] != '#'{
		if grid[pos.y][pos.x] == '.'{
			return false
		}
		pos = Vec2{pos.x + dir.x, pos.y + dir.y}
	}

	return true
}

func moveRobot(grid [][]byte, pos, dir Vec2){
	for grid[pos.y][pos.x] != '.'{
		pos = Vec2{pos.x + dir.x, pos.y + dir.y}
	}
	
	for grid[pos.y][pos.x] != '@'{
		prev := Vec2{pos.x - dir.x, pos.y - dir.y}
		grid[pos.y][pos.x] = grid[prev.y][prev.x]
		
		pos = prev
	}
	
	grid[pos.y][pos.x] = '.'
}



func moveRobot_p2(grid [][]byte, pos, dir Vec2){
	// horizontal movement remains the same
	if dir.y == 0{
		for grid[pos.y][pos.x] != '.'{
			pos = Vec2{pos.x + dir.x, pos.y + dir.y}
		}
		
		for grid[pos.y][pos.x] != '@'{
			prev := Vec2{pos.x - dir.x, pos.y - dir.y}
			grid[pos.y][pos.x] = grid[prev.y][prev.x]
			
			pos = prev
		}
		
		grid[pos.y][pos.x] = '.'
	
	} else {
		if (grid[pos.y][pos.x] == '.'){
			return
		}
		
		if (grid[pos.y][pos.x] == '['){
			// for a chain of [] on top of each other so that there wont be multiple move for one cell
			for (grid[pos.y][pos.x] == '['){
				pos = Vec2{pos.x + dir.x, pos.y + dir.y}
			}
			
			moveRobot_p2(grid, Vec2{pos.x, pos.y}, dir)
			moveRobot_p2(grid, Vec2{pos.x + 1, pos.y}, dir)

			pos = Vec2{pos.x - dir.x, pos.y - dir.y}

			for (grid[pos.y][pos.x] == '['){
				grid[pos.y + dir.y][pos.x + dir.x] = grid[pos.y][pos.x]
				grid[pos.y + dir.y][pos.x + dir.x + 1] = grid[pos.y][pos.x + 1]
				grid[pos.y][pos.x] = '.'
				grid[pos.y][pos.x + 1] = '.'
				pos = Vec2{pos.x - dir.x, pos.y - dir.y}
			}
			
			return
		}
		
		if (grid[pos.y][pos.x] == ']'){
			for (grid[pos.y][pos.x] == ']'){
				pos = Vec2{pos.x + dir.x, pos.y + dir.y}
			}
			
			moveRobot_p2(grid, Vec2{pos.x, pos.y}, dir)
			moveRobot_p2(grid, Vec2{pos.x - 1, pos.y}, dir)
			
			pos = Vec2{pos.x - dir.x, pos.y - dir.y}
			
			for (grid[pos.y][pos.x] == ']'){
				grid[pos.y + dir.y][pos.x + dir.x] = grid[pos.y][pos.x]
				grid[pos.y + dir.y][pos.x + dir.x - 1] = grid[pos.y][pos.x - 1]
				grid[pos.y][pos.x] = '.'
				grid[pos.y][pos.x - 1] = '.'
				pos = Vec2{pos.x - dir.x, pos.y - dir.y}
			}

			return
		}

		if (grid[pos.y][pos.x] == '@'){
			next := Vec2{pos.x + dir.x, pos.y + dir.y}
			moveRobot_p2(grid, next, dir)
			grid[next.y][next.x] = '@'
			grid[pos.y][pos.x] = '.'
		}
	
	}
	
}

func isConnectedToWall_p2(grid [][]byte, pos Vec2, dir Vec2) bool {
	if dir.y == 0{
		for grid[pos.y][pos.x] != '#'{
			if grid[pos.y][pos.x] == '.'{
				return false
			}
			pos = Vec2{pos.x + dir.x, pos.y + dir.y}
		}
		return true

	} else {
		for grid[pos.y][pos.x] != '#'{
			if grid[pos.y][pos.x] == '.'{
				return false
			}
			
			// for [  or  ], is connected to wall if any is connected to wall
			if (grid[pos.y][pos.x] == '['){
				isLeftConnected := isConnectedToWall_p2(grid, Vec2{pos.x + dir.x, pos.y + dir.y}, dir) 
				isRightConnected := isConnectedToWall_p2(grid, Vec2{pos.x + 1 + dir.x, pos.y + dir.y}, dir)
				
				return isLeftConnected || isRightConnected
			}

			if (grid[pos.y][pos.x] == ']'){
				isLeftConnected := isConnectedToWall_p2(grid, Vec2{pos.x - 1 + dir.x, pos.y + dir.y}, dir)
				isRightConnected := isConnectedToWall_p2(grid, Vec2{pos.x + dir.x, pos.y + dir.y}, dir) 
				
				return isLeftConnected || isRightConnected
			}
			pos = Vec2{pos.x + dir.x, pos.y + dir.y}
		}

		return true
	}

}






func Day15(){
	fmt.Println("--- Day 15: Warehouse Woes ---")

	file, err := os.Open("./inputs/day15.txt")
	
	if err != nil{
		panic(err)
	}
	defer file.Close()
	info, _ := file.Stat()
	
	bytes := make([]byte, info.Size())

	file.Read(bytes)
	
	grid_moves := strings.Split(string(bytes), "\r\n\r\n")

	gridStr := strings.Split(grid_moves[0], "\r\n")
	
	grid_p1 := make([][]byte, len(gridStr))
	grid_p2 := make([][]byte, len(gridStr))
	
	// convert to byte array as strings are immutable
	for i := range len(gridStr){
		grid_p1[i] = []byte(gridStr[i])
		
		// for part 2, expand each cell twice
		grid_p2[i] = make([]byte, 2 * len(gridStr[i]))
		for j := range len(gridStr[i]){
			if gridStr[i][j] == 'O'{
				grid_p2[i][2*j] = '['
				grid_p2[i][2*j+1] = ']'
			} else if gridStr[i][j] == '@'{
				grid_p2[i][2*j] = '@'
				grid_p2[i][2*j+1] = '.'
			} else {
				grid_p2[i][2*j] = gridStr[i][j]
				grid_p2[i][2*j+1] = gridStr[i][j]
			}
		}
	
	}
	

	
	moves := strings.Join(strings.Split(grid_moves[1], "\r\n"), "")

	robotPos_p1 := findRobotPos(grid_p1)	
	robotPos_p2 := findRobotPos(grid_p2)	

	for _, move := range moves{
		moveDirection := map[byte]Vec2{
			'>': DIR_RIGHT, 
			'<': DIR_LEFT, 
			'^': DIR_UP, 
			'v': DIR_DOWN, 
		}

		direction := moveDirection[byte(move)]
		
		if !isConnectedToWall(grid_p1, robotPos_p1, direction){
			moveRobot(grid_p1, robotPos_p1, direction)
			robotPos_p1 = Vec2{robotPos_p1.x + direction.x, robotPos_p1.y + direction.y}
		}
		
		if !isConnectedToWall_p2(grid_p2, robotPos_p2, direction){
			moveRobot_p2(grid_p2, robotPos_p2, direction)
			robotPos_p2 = Vec2{robotPos_p2.x + direction.x, robotPos_p2.y + direction.y}
		}
	}
	
	sum_p1 := 0 
	for i := range grid_p1{
		for j := range grid_p1[i]{
			if grid_p1[i][j] == 'O'{
				sum_p1 += i * 100 + j 
			}
			fmt.Print(string(grid_p1[i][j]))
		}
		fmt.Print("\n")
	}
	
	sum_p2 := 0 
	for i := range grid_p2{
		for j := range grid_p2[i]{
			if grid_p2[i][j] == '['{
				sum_p2 += i * 100 + j 
			}
			fmt.Print(string(grid_p2[i][j]))
		}
		fmt.Print("\n")
	}

	fmt.Println("Part 1: ", sum_p1)
	fmt.Println("Part 2: ", sum_p2)
}