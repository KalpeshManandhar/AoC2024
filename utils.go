package main

import (
	"fmt"
)

type Score struct{
	score int
	pos Vec2
	dir Vec2
}

type ScoreHeap []Score

func (h ScoreHeap) Len() int{
	return len(h)
} 

func (h ScoreHeap) Less(i, j int) bool{
	return h[i].score < h[j].score
} 

func (h ScoreHeap) Swap(i, j int){
	h[i], h[j] = h[j], h[i]
} 

func (h *ScoreHeap) Push(x any){
	*h = append(*h, x.(Score))
}

func (h *ScoreHeap) Pop() any{
	x := (*h)[len(*h) - 1]
	*h = (*h)[:len(*h) - 1]
	return x
}


func printGrid(grid []byte, w, h int){
	for i := range(h){
		for j := range(w){
			fmt.Print(string(grid[i * w + j]))
		}
		fmt.Print("\n")
	}
}