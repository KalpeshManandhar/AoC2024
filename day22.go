package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)


func Day22(){
	fmt.Println("--- Day 22: Monkey Market ---")

	file, err := os.Open("./inputs/day22.txt")

	if err != nil{
		panic(err)
	}
	defer file.Close()
	
	info, _ := file.Stat()
	
	bytes := make([]byte, info.Size())

	file.Read(bytes)


	buyers := strings.Split(string(bytes), "\r\n")
	
	N_NUMBER := 2000
	MOD := 16777216

	sum_p1 := 0
	
	const WINDOW_SIZE = 4
	
	type Key struct{
		window [WINDOW_SIZE]int
	}
	type Value struct{
		bananas int
		upto int
	}

	bananas := map[Key]Value{}

	for i, secret := range buyers{
		deltas := make([]int, N_NUMBER)
		numbers := make([]int, N_NUMBER)

		num, _ := strconv.Atoi(secret)
		
		for j := range(N_NUMBER){
			n := num
			n = ((n<<6)^n)%MOD
			n = ((n>>5)^n)%MOD
			n = ((n<<11)^n)%MOD

			deltas[j] = n%10 - num%10
			numbers[j] = n

			num = n
		}
		
		for j := WINDOW_SIZE - 1; j < N_NUMBER; j++{
			key := Key{}
			
			for k := range(WINDOW_SIZE){
				key.window[k] = deltas[j - WINDOW_SIZE + 1 + k]
			}

			if _, ok := bananas[key]; !ok{
				bananas[key] = Value{bananas: 0, upto: -1}
			}

			value := bananas[key]

			if value.upto < i {
				value.bananas += numbers[j]%10
				value.upto = i
				bananas[key] = value
			}
		}

		sum_p1 += num
	}

	
	max_bananas_p2 := 0
	for _, value := range bananas{
		if max_bananas_p2 < value.bananas{
			max_bananas_p2 = value.bananas
		}
	}


	fmt.Println("Part 1: ", sum_p1)
	fmt.Println("Part 2: ", max_bananas_p2)
}