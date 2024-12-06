package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	// key must come first
	// value must come second
	orderRules = map[string][]string{}
	part2Sum   = 0
)

func main() {
	fmt.Println("Hello, World!")
	filepath := "input.txt"
	puzzles := 1
	data, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	// use the fmt package to print the data to the console
	singleOutput := string(data)
	split := strings.Split(singleOutput, "\n\n")
	orders := strings.Split(split[0], "\n")
	for _, order := range orders {
		orderSplit := strings.Split(order, "|")
		orderRules[orderSplit[0]] = append(orderRules[orderSplit[0]], orderSplit[1])
	}

	lines := strings.Split(split[1], "\n")

	switch puzzles {
	case 1:
		part1(lines)
	default:
		fmt.Println("No puzzles selected")
	}
}

func part1(lines []string) {
	validTotal := 0
	for _, updateLine := range lines {
		updatePages := strings.Split(updateLine, ",")
		pagesIndex := map[string]int{}
		pagesIndex[updatePages[0]] = 0
		pageAlreadySeen := map[string]bool{}
		pageAlreadySeen[updatePages[0]] = true
		valid := true
	checkLine:
		for i := 1; i < len(updatePages); i++ {
			// go through page dependencies
			currPage := updatePages[i]
			for _, secondaryPage := range orderRules[currPage] {
				// if the secondary page has already been seen then this is an invalid update order
				if pageAlreadySeen[secondaryPage] {
					// this is out of order
					valid = false
					iToMove := pagesIndex[secondaryPage]
					part2(append(updatePages[:iToMove], append(updatePages[iToMove+1:], updatePages[iToMove])...))
					break checkLine
				}
			}
			pageAlreadySeen[currPage] = true
			pagesIndex[currPage] = i
		}
		if valid {
			addNum, _ := strconv.Atoi(updatePages[len(updatePages)/2])
			validTotal += addNum
		}
	}
	fmt.Printf("Part 1 solution: %v\n", validTotal)
	fmt.Printf("Part 2 solution: %v\n", part2Sum)
}

func part2(updatePages []string) {
	pagesIndex := map[string]int{}
	pageAlreadySeen := map[string]bool{}
	pageAlreadySeen[updatePages[0]] = true
	pagesIndex[updatePages[0]] = 0
	valid := true
checkLine:
	for i := 1; i < len(updatePages); i++ {
		// go through page dependencies
		currPage := updatePages[i]
		for _, secondaryPage := range orderRules[currPage] {
			// if the secondary page has already been seen then this is an invalid update order
			if pageAlreadySeen[secondaryPage] {
				// this is out of order
				valid = false
				iToMove := pagesIndex[secondaryPage]
				// sItoEnd := append(updatePages[iToMove+1:], updatePages[iToMove])
				// upToI := updatePages[:iToMove]
				// updatePages = append(upToI, sItoEnd...)
				updatePages = append(updatePages[:iToMove], append(updatePages[iToMove+1:], updatePages[iToMove])...)
				break checkLine
			}
		}
		pageAlreadySeen[currPage] = true
		pagesIndex[currPage] = i
	}
	if valid {
		addNum, _ := strconv.Atoi(updatePages[len(updatePages)/2])
		part2Sum += addNum
	} else {
		part2(updatePages)
	}
}
