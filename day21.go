package main

import (
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)


var numKeypadMap = map[byte]Vec2{
	'7' : {0,0},
	'8' : {1,0},
	'9' : {2,0},
	'4' : {0,1},
	'5' : {1,1},
	'6' : {2,1},
	'1' : {0,2},
	'2' : {1,2},
	'3' : {2,2},
	'0' : {1,3},
	'A' : {2,3},
}


var directionalKeypadMap = map[byte]Vec2{
	'^' : {1,0},
	'A' : {2,0},
	'<' : {0,1},
	'v' : {1,1},
	'>' : {2,1},
}


var directionToByteMap = map[Vec2]byte{
	DIR_LEFT : '<',
	DIR_RIGHT : '>',
	DIR_UP : '^',
	DIR_DOWN : 'v',
}

var numKeypadGrid = [][]byte{
	{'7', '8', '9'},
	{'4', '5', '6'},
	{'1', '2', '3'},
	{'#', '0', 'A'},
}

var directionalKeypadGrid = [][]byte{
	{'#', '^', 'A'},
	{'<', 'v', '>'},
}



type Info_d21 struct{
	pos Vec2
	inputMoves string 
}



func costFunc(moves string) int {
	cost := 0

	moves = moves + "A"
	pos := directionalKeypadMap['A']
	for i:=0; i< len(moves); i++{
		nextPos := directionalKeypadMap[byte(moves[i])]
		cost += Absi(pos.x - nextPos.x) + Absi(pos.y - nextPos.y) + 1
		pos = nextPos
	}

	return cost
}


type Move struct{
	from, to Vec2 
}


func findOptimalMoves(keypadMap map[byte]Vec2, grid [][]byte) map[Move][]string {
	optimalMoves := map[Move][]string{}
	
	
	for _, from := range keypadMap{
		for _, to := range keypadMap{
			optimalCost := math.MaxInt32
	
			m, n := len(grid), len(grid[0])

	
			costs := make([]int, m*n)

			for i := range costs{
				costs[i] = math.MaxInt32
			}

			nodeHeap := make(ScoreHeap[Info_d21], 0)

			heap.Init(&nodeHeap)
			heap.Push(&nodeHeap, Score[Info_d21]{score: 0, info: Info_d21{pos : from, inputMoves: ""}})
		
			for nodeHeap.Len() > 0{
				node := heap.Pop(&nodeHeap).(Score[Info_d21])
				pos := node.info.pos
				movesTillNow := node.info.inputMoves
				

				if node.score < costs[pos.y * n + pos.x] {
					costs[pos.y * n + pos.x] = node.score
				}
				


				if (pos == to){
					cost := costFunc(movesTillNow + "A")
					
					if cost < optimalCost {
						optimalMoves[Move{from, to}] = make([]string, 0)
						optimalMoves[Move{from, to}] = append(optimalMoves[Move{from, to}], movesTillNow + "A")
						optimalCost = cost

					} else if cost == optimalCost{
						optimalMoves[Move{from, to}] = append(optimalMoves[Move{from, to}], movesTillNow + "A")
					}
				}


				addNeighbour := func(inDirection Vec2){
					neighbour := Vec2{pos.x + inDirection.x, pos.y + inDirection.y}
					
					if (neighbour.x < 0 || neighbour.y < 0 || neighbour.x >= len(grid[0]) || neighbour.y >= len(grid)){
						return
					}

					if (grid[neighbour.y][neighbour.x] == '#'){
						return
					}
					
					
					moves := movesTillNow + string(directionToByteMap[inDirection])
					cost := costFunc(moves)
					if (cost <= costs[neighbour.y * n + neighbour.x]){
						heap.Push(&nodeHeap, Score[Info_d21]{score: cost, info: Info_d21{pos: neighbour, inputMoves: moves}})
					}
				}
				
				
				
				addNeighbour(DIR_LEFT)
				addNeighbour(DIR_RIGHT)
				addNeighbour(DIR_UP)
				addNeighbour(DIR_DOWN)
				
			
			}

		}
	}

	
	return optimalMoves

}



type CacheKey_d21 struct {
	robotLevel int
	humanLevel int
	reqOutput string
}

var cache_d21 = map[CacheKey_d21]int64{}

func findShortest(robotLevel int, humanLevel int, reqOutput string, optimalMoves map[Move][]string) int64 {
	if val, ok := cache_d21[CacheKey_d21{robotLevel: robotLevel, humanLevel : humanLevel,  reqOutput: reqOutput}]; ok{
		return val
	}

	if robotLevel == humanLevel{
		cache_d21[CacheKey_d21{robotLevel: robotLevel, humanLevel : humanLevel,  reqOutput: reqOutput}] = int64(len(reqOutput))
		return int64(len(reqOutput))
	}
	

	pos := directionalKeypadMap['A']
	pathLength := int64(0)
	for _, button := range reqOutput{
		nextPos := directionalKeypadMap[byte(button)]
		
		possiblePaths := optimalMoves[Move{pos, nextPos}]
		shortestPathLen := int64(math.MaxInt64)
		
		for _, path := range possiblePaths{
			length := findShortest(robotLevel + 1, humanLevel, path, optimalMoves)
			if length < shortestPathLen{
				shortestPathLen = length
			}
		}

		pathLength += shortestPathLen
		pos = nextPos
	}
	
	cache_d21[CacheKey_d21{robotLevel: robotLevel, humanLevel : humanLevel, reqOutput: reqOutput}] = pathLength
	return pathLength
}







func Day21(){
	fmt.Println("--- Day 21: Keypad Conundrum ---")

	file, err := os.Open("./inputs/day21.txt")

	if err != nil{
		panic(err)
	}
	defer file.Close()
	
	info, _ := file.Stat()
	
	bytes := make([]byte, info.Size())

	file.Read(bytes)
	
	codes := strings.Split(string(bytes), "\r\n")
	
	
	// find optimal moves for each (from, to) pair in both numeric and directional keypad
	optimalMovesNumKeypad := findOptimalMoves(numKeypadMap, numKeypadGrid)
	optimalMovesDirKeypad := findOptimalMoves(directionalKeypadMap, directionalKeypadGrid)
	
	
	const N_P1 = 25
	const N_P2 = 25

	sum_p1 := int64(0)
	sum_p2 := int64(0)
	
	// find min length for each robot for a given sequence of inputs
	for _, code := range codes{
		fmt.Println(code)
		
		pos := numKeypadMap['A']
		pathLength_p1 := int64(0)
		pathLength_p2 := int64(0)

		for _, button := range code{
			nextPos := numKeypadMap[byte(button)]
		
			possiblePaths := optimalMovesNumKeypad[Move{pos, nextPos}]
		
			shortestLength_p1 := int64(math.MaxInt64)
			shortestLength_p2 := int64(math.MaxInt64)
			
			for _, path := range possiblePaths{
				length_p1 := findShortest(0, N_P1, path, optimalMovesDirKeypad)
				if length_p1 < shortestLength_p1{
					shortestLength_p1 = length_p1
				}

				length_p2 := findShortest(0, N_P2, path, optimalMovesDirKeypad)
				if length_p2 < shortestLength_p2{
					shortestLength_p2 = length_p2
				}
			}

			pathLength_p1 += shortestLength_p1
			pathLength_p2 += shortestLength_p2
			pos = nextPos
		}

		val, _ := strconv.Atoi(code[:len(code)-1])
		sum_p1 += int64(val) * pathLength_p1
		sum_p2 += int64(val) * pathLength_p2

	}

	fmt.Println("Part 1: ", sum_p1)
	fmt.Println("Part 2: ", sum_p2)

}