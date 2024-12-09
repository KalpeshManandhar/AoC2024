package main

import (
	"fmt"
	"os"
	"strings"
)

type Vec2 struct{
	x,y int
}


func findGuardPos(lines []string) (pos Vec2, dir Vec2){
	pos = Vec2{}
	dir = Vec2{}
	
	directions := map[byte]Vec2{
		'<' : {x: -1, y: 0},
		'^' : {x: 0, y: -1},
		'>' : {x: 1, y: 0},
		'v' : {x: 0, y: 1},
	}
	
	for i := 0; i < len(lines); i++ {
		findGuard := func(c byte) bool{
			index := strings.IndexByte(lines[i], c)
			if (index == -1){
				return false
			}
			pos.y = i
			pos.x = index 
			return true
		}
		
		for key, val := range directions{
			if (findGuard(key)){
				dir = val
				return pos, dir
			}
		}
	}
	return pos, dir
}

var DIR_RIGHT = Vec2{x: 1, y: 0}
var DIR_LEFT = Vec2{x: -1, y: 0}
var DIR_DOWN = Vec2{x: 0, y: 1}
var DIR_UP = Vec2{x: 0, y: -1}


var rotateRight = map[Vec2]Vec2{
	DIR_DOWN : DIR_LEFT,
	DIR_LEFT  : DIR_UP,
	DIR_UP  : DIR_RIGHT,
	DIR_RIGHT : DIR_DOWN,
}
var rotateLeft = map[Vec2]Vec2{
	DIR_UP : DIR_LEFT,
	DIR_RIGHT  : DIR_UP,
	DIR_DOWN  : DIR_RIGHT,
	DIR_LEFT : DIR_DOWN,
}


var directions = map[Vec2]int{
	DIR_LEFT : 0x1,
	DIR_UP   : 0x2,
	DIR_RIGHT: 0x4,
	DIR_DOWN : 0x8,
}



func Day6(){
	fmt.Println("--- Day 6: Guard Gallivant ---")
	
	file, err := os.Open("./inputs/day6.txt")
	defer file.Close()
	
	if err != nil{
		panic(err)
	}
	info, _ := file.Stat()
	
	bytes := make([]byte, info.Size())
	
	file.Read(bytes)
	
	lines := strings.Split(string(bytes), "\r\n")
	
	
	m, n := len(lines), len(lines[0])
	toPut := make([]bool, m*n)
	

	isInsideRange := func(x, y int) bool {
		return x >= 0 && x < n && y >= 0 && y < m
	}
	

	ipos, idir := findGuardPos(lines)
	pos, dir := ipos, idir
	count := 0

	
	nWays := 0
	visitedDir := make([]int, m * n)
	
	doesLoop := func (currentPos, direction, obs Vec2) bool{
		visitedDirection := make([]int, m * n)
		copy(visitedDirection, visitedDir)
		for {
			// check if previously visited
			if (visitedDirection[currentPos.y * n + currentPos.x] & directions[direction]) != 0{
				return true
			}
			
			// if not visited, track current visit and direction
			visitedDirection[currentPos.y * n + currentPos.x] |= directions[direction]
			
			
			newPos := Vec2{}
			newPos.x = currentPos.x + direction.x
			newPos.y = currentPos.y + direction.y
	
	
			// if outside range, then break
			if !isInsideRange(newPos.x, newPos.y){
				break
			}
			
			// change direction
			if (lines[newPos.y][newPos.x] == '#' || (newPos.x == obs.x && newPos.y == obs.y)){
				direction = rotateRight[direction]	
				
			} else {
				// else move forward
				currentPos = newPos
			}
	
		}
		return false
	}

	
	checked := make([]bool, m * n)
	for {
		
		visitedDir[pos.y * n + pos.x] |= directions[dir]

		newPos := Vec2{}
		newPos.x = pos.x + dir.x
		newPos.y = pos.y + dir.y
		


		// if outside range, then break
		if !isInsideRange(newPos.x, newPos.y){
			break
		}
			


		// change direction
		if (lines[newPos.y][newPos.x] == '#'){
			dir = rotateRight[dir]	

			
		} else {
			if (!checked[newPos.y * n + newPos.x]){	
				// check if a loop would be formed if the next pos had an obstacle placed 
				if doesLoop(pos, rotateRight[dir], newPos){
					toPut[newPos.y * n + newPos.x] = true
				}
				checked[newPos.y * n + newPos.x] = true
			}
			// move forward
			pos = newPos
		}

	}

	// count number of cells where obstacles can be placed
	for i, put := range toPut{
		if i == (ipos.y * n + ipos.x){
			continue
		}
		if (put){
			nWays++
		}
	}
	
	// count number of cells where obstacles can be placed
	for _, visit := range visitedDir{
		if (visit > 0){
			count++
		}
	}


	fmt.Println("Part 1: ", count)
	fmt.Println("Part 2: ", nWays)
	

}
