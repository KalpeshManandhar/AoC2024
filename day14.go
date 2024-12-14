package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
)


func Day14(){
	fmt.Println("--- Day 14: Restroom Redoubt ---")

	file, err := os.Open("./inputs/day14.txt")
	defer file.Close()
	
	if err != nil{
		panic(err)
	}
	info, _ := file.Stat()
	
	bytes := make([]byte, info.Size())

	file.Read(bytes)
	
	robots := strings.Split(string(bytes), "\r\n")
	
	w := 101
	h := 103
	nSeconds := 100

	regex := regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)

	initialPos := make([]Vec2, len(robots))
	velocity := make([]Vec2, len(robots))
	
	for i, robot := range robots{
		match := regex.FindAllStringSubmatch(robot, -1)
		
		p := Vec2{}
		v := Vec2{}
		p.x, _ = strconv.Atoi(match[0][1])
		p.y, _ = strconv.Atoi(match[0][2])
		v.x, _ = strconv.Atoi(match[0][3])
		v.y, _ = strconv.Atoi(match[0][4])
		
		initialPos[i] = p
		velocity[i] = v
	}
	
	simulate := func (steps int) []int{
		grid := make([]int, w*h)

		for j := range robots{
			final := Vec2{(initialPos[j].x + velocity[j].x * steps)%w, (initialPos[j].y + velocity[j].y * steps)%h}
			final = Vec2{(final.x + w)%w, (final.y + h)%h} // for negative coordinates
			
			grid[final.y * w + final.x] += 1
		}
		return grid
	}


	printGrid := func(grid []int){
		for i, count := range grid{
			x := i%w
			
			if (x == 0){
				fmt.Print("\n")
			}
			fmt.Print(count)
		}
		fmt.Print("\n")
	}

	// part 1
	robotCount := simulate(nSeconds)

	quadrantCount := [2][2]int64{{0,0},{0,0}}
	for i, count := range robotCount{
		x := i%w
		y := i/w

		if (x == w/2 || y == h/2){
			continue
		}

		quadrantX := x*2/w
		quadrantY := y*2/h
		
		quadrantCount[quadrantY][quadrantX] += int64(count)
	}
	
	product_p1 := quadrantCount[0][0] * quadrantCount[0][1] * quadrantCount[1][0] * quadrantCount[1][1]


	// part 2
	calcScore := func(steps int) int{
		grid := simulate(steps)
		score := 0

		
		for j := 0; j < h; j++ {
			for k := 0; k < w; k++ {
				count := 0
				
				check := func(x, y int) int{
					if (x >= 0 && x < w && y >= 0 && y < h){
						return grid[y * w + x]
					} 
					return 0
				}

				count += check(k, j) + check(k+1, j) + check(k-1, j)
				count += check(k, j+1) + check(k+1, j+1) + check(k-1, j+1)
				count += check(k, j-1) + check(k+1, j-1) + check(k-1, j-1)

				score += count * count
			}
			
		}

		return score
	}
	

	// find 
	nThreads := 8
	maxScores := [8]int{}
	maxScoresAt := [8]int{}
	
	each := 5000
	
	var wg sync.WaitGroup
	
	for thread := range nThreads{
		start := each * thread
		end := start + each
		wg.Add(1)
		go func(start, end int, threadID int){	
			maxScore := math.MinInt32
			maxScoreAt := start
			for i := start; i<end; i++{
				score := calcScore(i)
				if (maxScore < score){
					maxScore = score
					maxScoreAt = i
				}
			}

			maxScores[threadID] = maxScore
			maxScoresAt[threadID] = maxScoreAt
			wg.Done()
		}(start, end, thread)
	}

	wg.Wait()
	
	// find global max score
	maxScoreAtGlobal := 0
	for thread := range(nThreads){
		if (maxScores[maxScoreAtGlobal] < maxScores[thread]){
			maxScoreAtGlobal = thread
		}
		
	}
	
	fmt.Println("Part 1: ", product_p1)
	fmt.Println("Part 2: ", maxScoresAt[maxScoreAtGlobal])
	easterEgg := simulate(maxScoresAt[maxScoreAtGlobal])
	printGrid(easterEgg)

}