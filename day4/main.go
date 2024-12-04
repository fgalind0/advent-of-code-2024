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
	acceptableLetter    = []string{"X", "M", "A", "S"}
	acceptableLetterPt2 = []string{"M", "S"}
	// XMAS
	nextLetter    = map[string]string{"X": "M", "M": "A", "A": "S", "S": "END"}
	nextLetterPt2 = map[string]string{"M": "S", "S": "M"}
	nextLetterDir = map[string]string{"UL": "DR", "DL": "UR"}
	// UU, UL, UR, LL, RR, DL, DD, DR
	nextLetterLoc = map[string]Modifier{"UU": {IdxChange: -1, PosChange: 0}, "UL": {IdxChange: -1, PosChange: -1}, "UR": {IdxChange: -1, PosChange: 1}, "LL": {IdxChange: 0, PosChange: -1}, "RR": {IdxChange: 0, PosChange: 1}, "DL": {IdxChange: 1, PosChange: -1}, "DD": {IdxChange: 1, PosChange: 0}, "DR": {IdxChange: 1, PosChange: 1}}
	numXmasFound  = 0
	numX_masFound = 0
)

func main() {
	fmt.Println("Hello, World!")
	filepath := "input.txt"
	puzzles := 2
	data, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	// use the fmt package to print the data to the console
	singleOutput := string(data)
	lines := strings.Split(singleOutput, "\n")
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
	if newIdx < 0 || newIdx >= len(lines)-1 || newPos < 0 || newPos >= len(lines[idx]) {
		return false
	}
	if string(lines[newIdx][newPos]) == desiredLetter {
		return searching(linesPtr, newIdx, newPos, direction)
	} else {
		return false
	}
}

func part2(lines []string) {
	for idx, line := range lines {
		for pos, letter := range line {
			if letter == 'A' {
				aFound(&lines, idx, pos)
			}
		}
	}
	fmt.Printf("Part 2 solution: %v\n", numX_masFound)
}

func aFound(linesPtr *[]string, idx, pos int) {
	firstDiagonal := searchingDiagonal(linesPtr, idx, pos, "UL")
	secondDiagonal := searchingDiagonal(linesPtr, idx, pos, "DL")
	if firstDiagonal && secondDiagonal {
		numX_masFound++
	}
}

func searchingDiagonal(linesPtr *[]string, idx, pos int, direction string) bool {
	lines := *linesPtr

	firstNewIdx := idx + nextLetterLoc[direction].IdxChange
	firstNewPos := pos + nextLetterLoc[direction].PosChange

	// determine if the first next index or position is out of bounds
	if firstNewIdx < 0 || firstNewIdx >= len(lines)-1 || firstNewPos < 0 || firstNewPos >= len(lines[idx]) {
		return false
	}

	nextDir := nextLetterDir[direction]
	// determine if the second next index or position is out of bounds
	secondNewIdx := idx + nextLetterLoc[nextDir].IdxChange
	secondNewPos := pos + nextLetterLoc[nextDir].PosChange
	if secondNewIdx < 0 || secondNewIdx >= len(lines)-1 || secondNewPos < 0 || secondNewPos >= len(lines[idx]) {
		return false
	}
	// now we know values are in bounds
	firstLetter := string(lines[firstNewIdx][firstNewPos])
	if !slices.Contains(acceptableLetterPt2, firstLetter) {
		return false
	}
	desiredLetter := nextLetterPt2[firstLetter]
	secondLetter := string(lines[secondNewIdx][secondNewPos])
	return secondLetter == desiredLetter
}
