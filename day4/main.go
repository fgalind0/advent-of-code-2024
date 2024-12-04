package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Modifier struct {
	IdxChange int
	PosChange int
}

var (
	acceptableLetter = []string{"X", "M", "A", "S"}
	// XMAS
	nextLetter = map[string]string{"X": "M", "M": "A", "A": "S", "S": "END"}
	// UU, UL, UR, LL, RR, DL, DD, DR
	nextLetterLoc = map[string]Modifier{"UU": {IdxChange: -1, PosChange: 0}, "UL": {IdxChange: -1, PosChange: -1}, "UR": {IdxChange: -1, PosChange: 1}, "LL": {IdxChange: 0, PosChange: -1}, "RR": {IdxChange: 0, PosChange: 1}, "DL": {IdxChange: 1, PosChange: -1}, "DD": {IdxChange: 1, PosChange: 0}, "DR": {IdxChange: 1, PosChange: 1}}
	numXmasFound  = 0
)

func main() {
	fmt.Println("Hello, World!")
	filepath := "input.txt"
	data, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	// use the fmt package to print the data to the console
	singleOutput := string(data)
	lines := strings.Split(singleOutput, "\n")
	part1(lines)
}

func part1(lines []string) {
	for idx, line := range lines {
		for pos, letter := range line {
			if letter == 'X' {
				xFound(&lines, idx, pos)
			}
		}
	}
	fmt.Printf("Part 1 solution: %v\n", numXmasFound)
}

// will handle kicking off searchfunc for each direction and handle incrementing when search returns true
func xFound(linesPtr *[]string, idx, pos int) {
	for direction := range nextLetterLoc {
		if searching(linesPtr, idx, pos, direction) {
			numXmasFound++
		}
	}
}

func searching(linesPtr *[]string, idx, pos int, direction string) bool {
	lines := *linesPtr
	currentLetter := string(lines[idx][pos])
	if !slices.Contains(acceptableLetter, currentLetter) {
		panic(fmt.Sprintf("Invalid letter %v\nindex: %v,position:%v", currentLetter, idx, pos))
	}
	desiredLetter := nextLetter[currentLetter]
	if desiredLetter == "END" {
		return true
	}
	newIdx := idx + nextLetterLoc[direction].IdxChange
	newPos := pos + nextLetterLoc[direction].PosChange
	// determining if the next index or position is out of bounds
	if newIdx < 0 || newIdx >= len(lines) || newPos < 0 || newPos >= len(lines[idx]) {
		return false
	}
	if string(lines[newIdx][newPos]) == desiredLetter {
		return searching(linesPtr, newIdx, newPos, direction)
	} else {
		return false
	}
}
