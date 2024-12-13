package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)


func det2(m00, m01, m10, m11 int) int{
	return m00 * m11 - m01 * m10
}


func solve(a0, b0, c0, a1, b1, c1 int) (x, y float64){
	d := float64(det2(a0, b0, a1, b1))
	
	if (d == 0){
		panic("Determinant 0")
	}
	
	d1 := float64(det2(c0, b0, c1, b1))
	d2 := float64(det2(a0, c0, a1, c1))

	return d1/d, d2/d
}



func Day13(){
	fmt.Println("--- Day 13: Claw Contraption ---")

	file, err := os.Open("./inputs/day13.txt")
	defer file.Close()
	
	if err != nil{
		panic(err)
	}
	info, _ := file.Stat()
	
	bytes := make([]byte, info.Size())

	file.Read(bytes)
	
	machines := strings.Split(string(bytes), "\r\n\r\n")

	regex := regexp.MustCompile(`Button A: X\+(\d+), Y\+(\d+)\r\nButton B: X\+(\d+), Y\+(\d+)\r\nPrize: X=(\d+), Y=(\d+)`)
	
	nTokens_p1 := 0
	nTokens_p2 := 0
	tokensUsedA := 3
	tokensUsedB := 1

	conversionError := 10000000000000

	for _, machineInfo := range machines{
		match := regex.FindAllStringSubmatch(machineInfo, -1)
		
		Ax, _ := strconv.Atoi(match[0][1])
		Ay, _ := strconv.Atoi(match[0][2])
		Bx, _ := strconv.Atoi(match[0][3])
		By, _ := strconv.Atoi(match[0][4])
		T1x, _ := strconv.Atoi(match[0][5])
		T1y, _ := strconv.Atoi(match[0][6])

		T2x := T1x + conversionError
		T2y := T1y + conversionError
		
		x1, y1 := solve(Ax, Bx, T1x, Ay, By, T1y)
		x2, y2 := solve(Ax, Bx, T2x, Ay, By, T2y)
		
		// fmt.Printf("P1: (%f, %f)\n", x1, y1)
		// fmt.Printf("P2: (%f, %f)\n", x2, y2)

		
		hasFracValue := func(x float64) bool{
			_, frac := math.Modf(x)
			return frac > 0
		}

		if !hasFracValue(x1) && !hasFracValue(y1){
			nTokens_p1 += (tokensUsedA * int(x1) + tokensUsedB * int(y1))
		}

		if !hasFracValue(x2) && !hasFracValue(y2){
			nTokens_p2 += (tokensUsedA * int(x2) + tokensUsedB * int(y2))
		}

		
	}

	fmt.Println("Part 1: ", nTokens_p1)
	fmt.Println("Part 2: ", nTokens_p2)
 

}