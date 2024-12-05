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
	case 0:
		fmt.Println("No puzzles selected")
	case 1:
		part1(lines)
	case 2:
		part2(lines)
	default:
		part1(lines)
		part2(lines)
	}
}

func part1(lines []string) {
	validTotal := 0
	for _, updateLine := range lines {
		updatePages := strings.Split(updateLine, ",")

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
					break checkLine
				}
			}
			pageAlreadySeen[currPage] = true
		}
		if valid {
			addNum, _ := strconv.Atoi(updatePages[len(updatePages)/2])
			validTotal += addNum
		}
	}
	fmt.Printf("Part 1 solution: %v\n", validTotal)
}

func part2(lines []string) {

}
