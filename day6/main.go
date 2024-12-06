package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
)

type GuardLoc struct {
	Idx int
	Pos int
	Dir rune
}

type LocMod struct {
	IdxChange int
	PosChange int
}

var (
	guardLoc = GuardLoc{Idx: 61, Pos: 78, Dir: 'U'}
	nextDir  = map[rune]rune{'U': 'R', 'R': 'D', 'D': 'L', 'L': 'U'}
	locMods  = map[rune]LocMod{'U': {IdxChange: -1, PosChange: 0}, 'R': {IdxChange: 0, PosChange: 1}, 'D': {IdxChange: 1, PosChange: 0}, 'L': {IdxChange: 0, PosChange: -1}}
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
	// mapOfLab := string(data)
	mapRows := bytes.Split(data, []byte("\n"))
	if filepath == "sample.txt" {
		guardLoc.Idx = 6
		guardLoc.Pos = 4
	}
	guard := mapRows[guardLoc.Idx][guardLoc.Pos]
	if guard == '^' {
		fmt.Println("Guard found")
	} else {
		panic(fmt.Sprintf("Guard not found. Instead found %c ", guard))
	}
	switch puzzles {
	case 1:
		part1(mapRows)
	default:
		fmt.Println("No puzzles selected")
	}

}

func part1(mapRows [][]byte) {
	// first we must modify current postion to become X
	mapRows[guardLoc.Idx][guardLoc.Pos] = 'X'
	// we must check next postion to confirm we can move
	nextIdx := guardLoc.Idx + locMods[guardLoc.Dir].IdxChange
	nextPos := guardLoc.Pos + locMods[guardLoc.Dir].PosChange
	// first need to see if next position is valid
	if nextIdx < 0 || nextIdx >= len(mapRows) || nextPos < 0 || nextPos >= len(mapRows[nextIdx]) {
		// convert mapRows to a string
		stringMap := string(bytes.Join(mapRows, []byte("\n")))
		// regex to find all "X"
		regex := regexp.MustCompile(`X`)
		fmt.Printf("Part 1 solution: %v\n", len(regex.FindAllString(stringMap, -1)))
	} else if mapRows[nextIdx][nextPos] != '#' {
		// we can move
		guardLoc.Idx = nextIdx
		guardLoc.Pos = nextPos
		part1(mapRows)
	} else {
		// we must change direction
		guardLoc.Dir = nextDir[guardLoc.Dir]
		part1(mapRows)
	}

}
