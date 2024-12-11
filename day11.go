package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)


func countDigits(n int) int{
	count := 0
	for n > 0{
		n /= 10
		count++
	}
	return count
}




/*
1
2024
20 24
2 0 2 4
4048 1 4048 8096
40 48 2024 4048 8096  



*/

type Key struct{
	val, n int
}


var cache map[Key]int


func count(val, n int) int{
	if (cache == nil){
		cache = map[Key]int{}
	}

	if c, ok := cache[Key{val, n}]; ok{
		return c
	}
	
	if (n == 0){
		return 1
	}
	
	c := 0
	ndigits := countDigits(val)

	if (val == 0){
		c = count(1, n - 1)
	}else if (ndigits % 2 == 0){
		halfDigitsPow10 := int(math.Pow(10, float64(ndigits)/2))
		front := val/halfDigitsPow10
		back := val%halfDigitsPow10
		c = count(front, n-1) + count(back, n-1)
	} else{
		c = count(val * 2024, n-1) 
	}
	
	cache[Key{val, n}] = c
	return c;
}


func Day11(){
	fmt.Println("--- Day 11: Plutonian Pebbles ---")

	file, err := os.Open("./inputs/day11.txt")
	defer file.Close()
	
	if err != nil{
		panic(err)
	}
	info, _ := file.Stat()
	
	bytes := make([]byte, info.Size())

	file.Read(bytes)
	rocks := strings.Split(string(bytes), " ")
	
	nTurns_p1 := 25
	nTurns_p2 := 75

	count_p1 := 0
	count_p2 := 0
	for _, rock := range rocks{
		rockVal, _ := strconv.Atoi(rock)

		count_p1 += count(rockVal, nTurns_p1)
		count_p2 += count(rockVal, nTurns_p2)
	}

	fmt.Println("Part 1: ", count_p1)
	fmt.Println("Part 2: ", count_p2)


}