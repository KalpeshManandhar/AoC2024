package main

import (
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)



func Day18(){
	fmt.Println("--- Day 18: RAM Run ---")

	file, err := os.Open("./inputs/day18.txt")
	defer file.Close()
	
	if err != nil{
		panic(err)
	}
	info, _ := file.Stat()
	
	bytes := make([]byte, info.Size())

	file.Read(bytes)


	start := Vec2{0,0}
	end := Vec2{70,70}
	
	N_SIMULATE := 1024

	m, n := end.y + 1, end.x + 1
	
	
	

	// simulate the bytes falling on the space
	pairs := strings.Split(string(bytes), "\r\n")
	
	simulateBytesFalling := func(upto int) []byte{
		grid := make([]byte, m*n)
		for i := range grid{
			grid[i] = '.'
		}
		
		for i:=0; i<upto; i++{
			nums := strings.Split(pairs[i], ",")
			
			x, _ := strconv.Atoi(nums[0])
			y, _ := strconv.Atoi(nums[1])
	
			grid[y * n + x] = '#'
		}

		return grid
	}

	// part 1
	grid := simulateBytesFalling(N_SIMULATE)
	// printGrid(grid, n, m)


	// find shortest path
	findMinMoves := func (grid []byte) int {
		visited := make([]bool, m*n)
		
		nodeHeap := make(ScoreHeap, 0)
		heap.Init(&nodeHeap)
		heap.Push(&nodeHeap, Score{score: 0, pos: start})
		visited[start.y * n + start.x] = true

		minMoves := 0
		
		// Dijkstra's
		for nodeHeap.Len() > 0{
			node := heap.Pop(&nodeHeap).(Score)
			pos := node.pos
			score := node.score
			
			if (pos == end){
				minMoves = score
				break
			}

			addNeighbour := func(inDirection Vec2){
				neighbour := Vec2{pos.x + inDirection.x, pos.y + inDirection.y}
				if (neighbour.x < 0 || neighbour.y < 0 || neighbour.x >= n || neighbour.y >= m){
					return
				}

				if (grid[neighbour.y * n + neighbour.x] == '#'){
					return
				}
				
				if (visited[neighbour.y *n + neighbour.x]){
					return
				}

				visited[neighbour.y * n + neighbour.x] = true
				heap.Push(&nodeHeap, Score{score: score + 1, pos:neighbour})
			}

			addNeighbour(DIR_LEFT)
			addNeighbour(DIR_RIGHT)
			addNeighbour(DIR_UP)
			addNeighbour(DIR_DOWN)
		}

		return minMoves
	}

	// part 1
	minMoves_p1 := findMinMoves(grid)
	fmt.Println("Part 1: ", minMoves_p1)
	
	
	
	// part 2
	l := 0
	r := len(pairs) - 1
	for l <= r {
		mid := (l + r)/2
		
		grid := simulateBytesFalling(mid + 1)
		moves :=  findMinMoves(grid)
		if (moves > 0){
			l = mid + 1
		} else{
			r = mid - 1
		}
	}

	firstByteBlockingEnd_p2 := pairs[l]
		
		
	fmt.Println("Part 2: ", firstByteBlockingEnd_p2)
	
}