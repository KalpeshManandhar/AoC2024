package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)


func checkIfSafe(levels []string) bool{
	diff1 := make([]int, len(levels))

	increasing := 0
	decreasing := 0
	same := 0

	for i := 0; i < len(levels) - 1; i++ {
		next, _ := strconv.Atoi(levels[i+1])
		current, _ := strconv.Atoi(levels[i])
		diff1[i] = next - current
		
		if (diff1[i] > 0){
			increasing++
		}else if (diff1[i] < 0){
			decreasing++
		}else {
			same++
		}
	}
	diff1[len(diff1)-1] = 1
	
	if (same > 1){
		return false
	}
	
	var isIncreasing bool

	if (increasing >= decreasing){
		if (decreasing > 1){
			return false
		}
		isIncreasing = true 
	}else{
		if (increasing > 1){
			return false
		}
		isIncreasing = false
	}

	for i := 0; i<len(diff1)-1; i++{
		if !isDiffValid(diff1[i], isIncreasing){
			return false
		}		
	}

	return true
}

func isDiffValid(diff int, isDirectionIncreasing bool) bool{
	sign := true

	if (isDirectionIncreasing){
		sign = diff > 0
	} else {
		sign = diff < 0
	}

	val := (math.Abs(float64(diff)) >= 1) && (math.Abs(float64(diff)) <= 3)
	
	valid := sign && val

	return valid
}


func checkIfSafeWDamping(levels []string) bool{
	diff1 := make([]int, len(levels))
	diff2 := make([]int, len(levels))
	
	increasing := 0
	decreasing := 0
	same := 0
	
	// calculate differences between consecutive numbers
	for i := 0; i < len(levels) - 1; i++ {
		next, _ := strconv.Atoi(levels[i+1])
		current, _ := strconv.Atoi(levels[i])
		diff1[i] = next - current
		
		if (diff1[i] > 0){
			increasing++
		}else if (diff1[i] < 0){
			decreasing++
		}else {
			same++
		}
	}
	diff1[len(diff1)-1] = 1
	
	// determine direction : only one diff can be resolved by damping
	if (same > 1){
		return false
	}
	
	var isIncreasing bool

	if (increasing >= decreasing){
		if (decreasing > 1){
			return false
		}
		isIncreasing = true 
	}else{
		if (increasing > 1){
			return false
		}
		isIncreasing = false
	}

	

	// edge diff2
	if isIncreasing{
		diff2[0] = 1
		diff2[len(diff2)-1] = 1
	} else{
		diff2[0] = -1
		diff2[len(diff2)-1] = -1
	}
	
	// calculate the diff if that element is removed
	for i := 1; i < len(diff2)-1; i++ {
		diff2[i] = diff1[i-1] + diff1[i]
	}
	
	isDamped := false 

	for i := 0; i<len(diff1)-1; i++{
		if isDiffValid(diff1[i], isIncreasing){
			continue
		}
		if isDamped{
			return false
		}
		if (!isDiffValid(diff2[i], isIncreasing) && !isDiffValid(diff2[i+1], isIncreasing)){
			return false
		}
		if (!isDiffValid(diff1[i+1], isIncreasing) && isDiffValid(diff2[i+1], isIncreasing)){
			i++
		}

		isDamped = true
	}

	return true
}


func Day2(){
	fmt.Println("--- Day 2: Red-Nosed Reports ---")

	file, err := os.Open("./inputs/day2.txt")
	defer file.Close()
	
	if err != nil{
		panic(err)
	}
	info, _ := file.Stat()
	
	bytes := make([]byte, info.Size())

	file.Read(bytes)
	
	lines := strings.Split(string(bytes), "\r\n")
	
	nSafeReports := 0
	nSafeReportsWDamping := 0

	for _, line := range lines{
		levels := strings.Split(line, " ")
		
		if checkIfSafe(levels){
			nSafeReports++
		}		
		
		if checkIfSafeWDamping(levels){
			nSafeReportsWDamping++
		}		
	}

	fmt.Println("Part 1: ", nSafeReports)
	fmt.Println("Part 2: ", nSafeReportsWDamping)

}