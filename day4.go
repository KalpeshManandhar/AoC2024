package main

import (
	"os"
	"fmt"
	"strings"
)


func Day4(){
	fmt.Println("--- Day 4: Ceres Search ---")

	file, err := os.Open("./inputs/day4.txt")
	
	if err != nil{
		panic(err)
	}
	defer file.Close()
	info, _ := file.Stat()
	
	bytes := make([]byte, info.Size())

	file.Read(bytes)
	
	rows := strings.Split(string(bytes), "\r\n")
	
	m,n := len(rows), len(rows[0])

	
	findXMAS := func (x,y int) int {
		count := 0
		str := "XMAS"
			
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				isValid := true
				
				for k := 0; k < len(str); k++ {
					yy := y + k * i
					xx := x + k * j
					
					if (yy < 0 || yy >= m){
						isValid = false
						break;
					}
					if (xx < 0 || xx >= n){
						isValid = false
						break;
					}
					
					if (rows[yy][xx] != str[k]){
						isValid = false
						break
					}
				}
				
				if (isValid){
					count++
				}
			}
			
		}
		
		return count
	}
	
	findX_MAS := func (x,y int) int {
		dirs := [][]int{{1,1}, {-1,1}}
		
		for _, dir := range dirs{
			y1 := y - dir[0]
			y2 := y + dir[0]
			x1 := x - dir[1]
			x2 := x + dir[1]
			
			if (y1 < 0 || y1 >= m || y2 < 0 || y2 >= m){
				return 0;
			}
			if (x1 < 0 || x1 >= n || x2 < 0 || x2 >= n){
				return 0;
			}
			
			if !((rows[y1][x1] == 'M' && rows[y2][x2] == 'S') || (rows[y1][x1] == 'S' && rows[y2][x2] == 'M')){
				return 0
			}
			
		}
		
		return 1
	}
	
	totalCount_XMAS := 0
	totalCount_X_MAS := 0


	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (rows[i][j] == 'X'){
				totalCount_XMAS += findXMAS(j, i)
			}
			if (rows[i][j] == 'A'){
				totalCount_X_MAS += findX_MAS(j, i)
			}
		}
	}

	fmt.Println("Part 1: ", totalCount_XMAS)
	fmt.Println("Part 2: ", totalCount_X_MAS)
	



}