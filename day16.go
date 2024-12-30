package main

import (
	"container/heap"
	"fmt"
	"math"
	"os"
	"strings"
)


func findStartEnd(grid []string) (Vec2, Vec2){
	start := Vec2{}
	end := Vec2{}
	for i, row := range grid{
		for j, cell := range row{
			if (cell == 'S'){
				start.x = j
				start.y = i
			}
			if (cell == 'E'){
				end.x = j
				end.y = i
			}
		}
	}
	return start, end
}


type NodeInfo struct {
	score [4]int
	from []Vec2
}


var directionIndex = map[Vec2]int{
	DIR_RIGHT : 0,
	DIR_LEFT : 1,
	DIR_UP : 2,
	DIR_DOWN : 3,
}


type ScoreInfo_d16 struct{
	pos Vec2
	dir Vec2
}



func Day16(){
	fmt.Println("--- Day 16: Reindeer Maze ---")

	file, err := os.Open("./inputs/day16.txt")
	if err != nil{
		panic(err)
	}
	defer file.Close()
	
	info, _ := file.Stat()
	
	bytes := make([]byte, info.Size())

	file.Read(bytes)
	
	grid := strings.Split(string(bytes), "\r\n")
	m, n := len(grid), len(grid[0])
	visited := make([]bool, m*n)

	start, end := findStartEnd(grid)

	nodeHeap := make(ScoreHeap[ScoreInfo_d16], 0)
	heap.Init(&nodeHeap)
	heap.Push(&nodeHeap, Score[ScoreInfo_d16]{score: 0, info: ScoreInfo_d16{start, DIR_RIGHT}})

	paths := make([]NodeInfo, m*n)
	for i := range paths{
		paths[i] = NodeInfo{
			[4]int{math.MaxInt32, math.MaxInt32, math.MaxInt32, math.MaxInt32}, 
			make([]Vec2, 0),
		}
	}
	paths[start.y * n + start.x].score[directionIndex[DIR_RIGHT]] = 0


	score_p1 := 0
	

	// Dijkstra's
	for nodeHeap.Len() > 0{
		node := heap.Pop(&nodeHeap).(Score[ScoreInfo_d16])
		pos := node.info.pos
		if (visited[pos.y * n + pos.x]){
			continue
		}

		dir := node.info.dir
		visited[pos.y * n + pos.x] = true
		
		// set the cost for current node in each direction
		paths[pos.y * n + pos.x].score[directionIndex[dir]] = node.score
		paths[pos.y * n + pos.x].score[directionIndex[rotateLeft[dir]]] = node.score + 1000
		paths[pos.y * n + pos.x].score[directionIndex[rotateRight[dir]]] = node.score + 1000

		if (pos == end){
			score_p1 = node.score
			break
		}

		addNeighbour := func(inDirection Vec2){
			neighbour := Vec2{pos.x + inDirection.x, pos.y + inDirection.y}
			
			if (grid[neighbour.y][neighbour.x] == '#'){
				return
			}
			
			score := paths[pos.y * n + pos.x].score[directionIndex[inDirection]] + 1
			
			// find minimum cost in current direction for neighbour
			if (score <= paths[neighbour.y * n + neighbour.x].score[directionIndex[inDirection]]){
				paths[neighbour.y * n + neighbour.x].score[directionIndex[inDirection]] = score
			}
			paths[neighbour.y * n + neighbour.x].from = append(paths[neighbour.y*n+neighbour.x].from, pos)
			
			if (visited[neighbour.y *n + neighbour.x]){
				return
			}
			
			heap.Push(&nodeHeap, Score[ScoreInfo_d16]{score: score, info: ScoreInfo_d16{neighbour, inDirection}})
		}
		
		

		addNeighbour(dir)
		addNeighbour(rotateLeft[dir])
		addNeighbour(rotateRight[dir])

	}
	

	// part 2
	counted := make([]bool, m*n)
	
	var count func(current Vec2) int
	count = func(current Vec2) int{
		// if already counted skip
		if (counted[current.y * n + current.x]){
			return 0
		}
		
		c := 1
		counted[current.y * n + current.x] = true
		
		for _, from := range paths[current.y * n + current.x].from {
			direction := Vec2{current.x - from.x, current.y - from.y}
			
			// check if the node has score equal to the score when going back from end with best path
			score := paths[current.y * n + current.x].score[directionIndex[direction]]
			fromScore := score - 1
			
			// check if lies in best path
			if (fromScore < paths[from.y * n + from.x].score[directionIndex[direction]]){
				continue
			}
			
			// if lies in best path, update the cost for rotating
			paths[from.y * n + from.x].score[directionIndex[direction]] = fromScore
			paths[from.y * n + from.x].score[directionIndex[rotateLeft[direction]]] = fromScore - 1000
			paths[from.y * n + from.x].score[directionIndex[rotateRight[direction]]] = fromScore - 1000

			c += count(from)
		}
		return c
	} 

	nSeats_p2 := count(end)


	for i, row := range grid{
		for j := range row{
			if counted[i * n + j] {
				fmt.Print("O")
			}else{
				fmt.Print(string(grid[i][j]))
			}
		}
		fmt.Print("\n")
	}
	
	fmt.Println("Part 1: ", score_p1)
	fmt.Println("Part 2: ", nSeats_p2)

	
}