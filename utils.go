package main

import (
	"fmt"
)

type Vec2 struct{
	x,y int
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



type Score[T any] struct{
	score int
	info T
}

type ScoreHeap[T any] []Score[T]

func (h ScoreHeap[T]) Len() int{
	return len(h)
} 

func (h ScoreHeap[T]) Less(i, j int) bool{
	return h[i].score < h[j].score
} 

func (h ScoreHeap[T]) Swap(i, j int){
	h[i], h[j] = h[j], h[i]
} 

func (h *ScoreHeap[T]) Push(x any){
	*h = append(*h, x.(Score[T]))
}

func (h *ScoreHeap[T]) Pop() any{
	x := (*h)[len(*h) - 1]
	*h = (*h)[:len(*h) - 1]
	return x
}


func printGrid(grid any, w, h int){
	for i := range(h){
		for j := range(w){
			if _, ok := grid.([]int); ok{
				fmt.Print(grid.([]int)[i * w + j])
			
			} else if _, ok := grid.([]byte); ok{
				fmt.Print(string(grid.([]byte)[i * w + j]))
			}

		}
		fmt.Print("\n")
	}
}


func Absi(a int) int {
	if a < 0{
		a = -a
	}
	return a
}

func Contains[T comparable](arr []T, elem T) bool{
	for _, e := range arr {
		if elem == e{
			return true
		}
	}
	return false
}