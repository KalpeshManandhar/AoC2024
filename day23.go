package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)



func Day23(){
	fmt.Println("--- Day 23: LAN Party ---")

	file, err := os.Open("./inputs/day23.txt")

	if err != nil{
		panic(err)
	}
	defer file.Close()
	
	info, _ := file.Stat()
	
	bytes := make([]byte, info.Size())

	file.Read(bytes)


	lines := strings.Split(string(bytes), "\r\n\r\n")

	type Group struct{
		computers [3]string
	}


	compConnectionsMap := map[string][]string{}
	// add connections for each computer
	for _, line := range lines{
		computers := strings.Split(line, "-")
		
		if _, ok := compConnectionsMap[computers[0]]; !ok{
			compConnectionsMap[computers[0]] = make([]string, 0)
		}
		if _, ok := compConnectionsMap[computers[1]]; !ok{
			compConnectionsMap[computers[1]] = make([]string, 0)
		}

		compConnectionsMap[computers[0]] = append(compConnectionsMap[computers[0]], computers[1])
		compConnectionsMap[computers[1]] = append(compConnectionsMap[computers[1]], computers[0])
	}

	groups := map[Group]bool{}
	
	// form each group of 3
	for computer, connections := range compConnectionsMap{
		group := Group{}
		for i := 0; i < len(connections); i++{
			
			for j := i+1; j < len(connections); j++{

				if Contains(compConnectionsMap[connections[i]], connections[j]){
					group.computers[2] = connections[j]
					group.computers[1] = connections[i]
					group.computers[0] = computer

					sort.Strings(group.computers[:])
				
					groups[group] = true
				}
			}
		}
	}

	// part 1, count number of groups that have a computer starting with t
	nComputers_p1 := 0
	for group := range groups{
		for _, computer := range group.computers{
			if (computer[0] == 't'){
				nComputers_p1++
				break
			}
		}
 	}


	
	// part 2
	computers := make([]string, len(compConnectionsMap))
	i := 0
	for computer := range compConnectionsMap{
		computers[i] = computer
		i++
	}
	sort.Strings(computers)
 
	
	// find all the fully connected groupings of strings given by appending and not appending a given computer
	connectedGroups := make([][]string, 0)
	connectedGroups = append(connectedGroups, make([]string, 0))

	for i := range computers {
		newGroupsTaken := make([][]string, 0)
		for j := range connectedGroups {
			
			// check if fully connected
			// since all those in the groups are already fully connected, only check if the current string to be appended is fully connected to current group
			isConnected := true
			for _, comp := range connectedGroups[j]{
				if !Contains(compConnectionsMap[comp], computers[i]){
					isConnected = false
					break
				}
			}

			if isConnected{
				newGroup := append([]string{}, connectedGroups[j]...)
				newGroup = append(newGroup, computers[i])
				newGroupsTaken = append(newGroupsTaken, newGroup)
			}
		}

		connectedGroups = append(connectedGroups, newGroupsTaken...)
	}


	maxConnectedGroup := ""
	maxConnectedGroupIndex := 0
	for i := range connectedGroups{
		if len(connectedGroups[maxConnectedGroupIndex]) < len(connectedGroups[i]){
			maxConnectedGroupIndex = i 
		}
	}

	maxConnectedGroup = strings.Join(connectedGroups[maxConnectedGroupIndex], ",")


	fmt.Println("Part 1: ", nComputers_p1)
	fmt.Println("Part 2: ", maxConnectedGroup)
}