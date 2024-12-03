package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func part1(singleOutput string) {
	instructionRegex := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	valuesRegex := regexp.MustCompile(`\d{1,3}`)
	matches := instructionRegex.FindAllString(singleOutput, -1)
	result := 0

	for _, match := range matches {
		values := valuesRegex.FindAllString(match, -1)
		if len(values) != 2 {
			panic(fmt.Sprintf("Invalid values %v", match))
		}
		first, _ := strconv.Atoi(values[0])
		second, _ := strconv.Atoi(values[1])
		result += first * second
	}
	fmt.Printf("Part 1 program output: %v\n", result)
}

type Zone struct {
	Start    int
	End      int
	EndIndex int
}

func part2(singleOutput string) {
	singleOutput = strings.ReplaceAll(singleOutput, "\n", "")
	validZones := []Zone{}
	dontRegex := regexp.MustCompile(`don\'t\(\)`)
	doRegex := regexp.MustCompile(`do\(\)`)
	donts := dontRegex.FindAllStringIndex(singleOutput, -1)
	dos := doRegex.FindAllStringIndex(singleOutput, -1)
	// valid zones start at 0 and go on until we hit the first don't
	// then a valid zone starts again until we hit the next don't
	// first valud zone is going to be zero and first index of dont
	validZones = append(validZones, Zone{Start: 0, End: donts[0][0], EndIndex: 0})
	// to determine rest of valid zones we will loop through all do s and make that the start
	shouldBreak := false
	for _, val := range dos {
		start := val[0]
		lastInsert := validZones[len(validZones)-1]
		if lastInsert.End > start {
			continue
		}
		end := 0
		endIndex := 0
		if lastInsert.EndIndex+1 < len(donts) {
			end = donts[lastInsert.EndIndex][0]
			endIndex = lastInsert.EndIndex + 1
		} else {
			end = len(singleOutput)
			endIndex = len(donts)
			shouldBreak = true
		}
		validZones = append(validZones, Zone{Start: start, End: end, EndIndex: endIndex})
		if shouldBreak {
			break
		}
	}
	// fmt.Printf("Valid Zones: %+v\n", validZones)
	instructionRegex := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	valuesRegex := regexp.MustCompile(`\d{1,3}`)
	matches := instructionRegex.FindAllString(singleOutput, -1)
	matchesLocations := instructionRegex.FindAllStringIndex(singleOutput, -1)
	result := 0

	for i, match := range matches {
		// must first check matchLocation is in a valid zone
		for _, zone := range validZones {
			if matchesLocations[i][0] > zone.Start && matchesLocations[i][0] < zone.End {
				values := valuesRegex.FindAllString(match, -1)
				if len(values) != 2 {
					panic(fmt.Sprintf("Invalid values %v", match))
				}
				first, _ := strconv.Atoi(values[0])
				second, _ := strconv.Atoi(values[1])
				result += first * second
			}
		}
	}
	fmt.Printf("Part 2 OG program output: %v\n", result)

}

func part2try2(singleOutput string) {
	singleOutput = strings.ReplaceAll(singleOutput, "\n", "")
	disregardRegex := regexp.MustCompile(`don\'t\(\).*?do\(\)`)
	singleOutput = disregardRegex.ReplaceAllString(singleOutput, "")

	// fmt.Printf("Valid Zones: %+v\n", validZones)
	instructionRegex := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	valuesRegex := regexp.MustCompile(`\d{1,3}`)
	matches := instructionRegex.FindAllString(singleOutput, -1)
	result := 0

	for _, match := range matches {
		values := valuesRegex.FindAllString(match, -1)
		if len(values) != 2 {
			panic(fmt.Sprintf("Invalid values %v", match))
		}
		first, _ := strconv.Atoi(values[0])
		second, _ := strconv.Atoi(values[1])
		result += first * second
	}
	fmt.Printf("Part 2 t2 program output: %v\n", result)

}

func main() {
	fmt.Println("Hello, World!")
	filepath := "input.txt"
	data, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	// use the fmt package to print the data to the console
	singleOutput := string(data)
	part1(singleOutput)
	part2(singleOutput)
	part2try2(singleOutput)
}
