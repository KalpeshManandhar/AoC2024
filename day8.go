package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)


func Day8(){
	fmt.Println("--- Day 8: Resonant Collinearity ---")

	file, err := os.Open("./inputs/day8.txt")
	defer file.Close()
	
	if err != nil{
		panic(err)
	}
	info, _ := file.Stat()
	
	bytes := make([]byte, info.Size())

	file.Read(bytes)
	grid := strings.Split(string(bytes), "\r\n")
	m, n := len(grid), len(grid[0])

	antinodes := make([]bool, m * n)
	
	isInsideRange := func(x, y int) bool{
		return x >=0 && x < n && y >= 0 && y < m
	}
	
	setAntinode := func(x, y int){
		if (isInsideRange(x, y)){
			antinodes[y * n + x] = true
		}
	}
	
	
	
	freqAntennas := make([][]Vec2, 26 + 26 + 10)
	for i := range freqAntennas{
		freqAntennas[i] = make([]Vec2, 0)
	}

	for i := range grid{
		for j := range grid[i]{
			if (unicode.IsDigit(rune(grid[i][j]))){
				freqAntennas[grid[i][j] - '0'] = append(freqAntennas[grid[i][j] - '0'], Vec2{j, i})
			} else if (unicode.IsLower(rune(grid[i][j]))){
				freqAntennas[grid[i][j] - 'a' + 10] = append(freqAntennas[grid[i][j] - 'a' + 10], Vec2{j, i})		
			} else if (unicode.IsUpper(rune(grid[i][j]))){
				freqAntennas[grid[i][j] - 'A' + 10 + 26] = append(freqAntennas[grid[i][j] - 'A' + 10 + 26], Vec2{j, i})		
			}
		}
	}
	
	for _, antennas := range freqAntennas{
		for i, _ := range antennas{
			for j := i+1; j < len(antennas); j++{
				diff := Vec2{x: antennas[i].x - antennas[j].x, y: antennas[i].y - antennas[j].y}

				one := Vec2{x: antennas[i].x + diff.x, y: antennas[i].y + diff.y}
				two := Vec2{x: antennas[j].x - diff.x, y: antennas[j].y - diff.y}
				
				setAntinode(one.x, one.y)
				setAntinode(two.x, two.y)
			}
		}
	}

	nAntinodes_p1 := 0
	
	for _, isAntinode := range antinodes{
		if (isAntinode){
			nAntinodes_p1++
		}
	}
	
	nAntinodes_p2 := 0

	for _, antennas := range freqAntennas{
		for i, _ := range antennas{
			for j := i+1; j < len(antennas); j++{
				diff := Vec2{x: antennas[i].x - antennas[j].x, y: antennas[i].y - antennas[j].y}
				
				k := 0
				for {
					one := Vec2{x: antennas[i].x + k * diff.x, y: antennas[i].y + k * diff.y}
					if (!isInsideRange(one.x, one.y)){
						break
					}
					setAntinode(one.x, one.y)
					k++
				}
				
				k = 0
				for {
					two := Vec2{x: antennas[j].x - k * diff.x, y: antennas[j].y - k * diff.y}
					if (!isInsideRange(two.x, two.y)){
						break
					}
					setAntinode(two.x, two.y)
					k++
				}
			}
		}
	}

	for _, isAntinode := range antinodes{
		if (isAntinode){
			nAntinodes_p2++
		}
	}


	fmt.Println("Part 1: ", nAntinodes_p1)
	fmt.Println("Part 2: ", nAntinodes_p2)

}