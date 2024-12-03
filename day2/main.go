package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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
	// conver a string to a slice based on the newline character
	// https://pkg.go.dev/strings#Split
	reports := strings.Split(singleOutput, "\n")
	fmt.Printf("Part 1 num safe: %v\nPart 2 numSafe: %v\n", puzzle1(reports), puzzle2(reports))

}

func puzzle1(reports []string) int {
	numSafe := 0
	for _, report := range reports {
		// replace spaces in element with a comma
		levels := strings.Split(report, " ")
		// convert the string to an int
		// https://pkg.go.dev/strconv#Atoi
		// int, err := strconv.Atoi(element)
		// append the int to the finalArr
		safe := true
		prevSign := 0
		sign := 0
		if report != "" {
			for i := 1; i < len(levels); i++ {
				prev, _ := strconv.Atoi(levels[i-1])
				curr, _ := strconv.Atoi(levels[i])
				diff := curr - prev
				// update sign if diff is positive
				if diff > 0 {
					sign = 1
				} else if diff < 0 {
					sign = -1
				}
				if diff > 3 || diff < -3 || diff == 0 || sign+prevSign == 0 {
					fmt.Printf("prev: %v, curr: %v, diff: %v\n", prev, curr, diff)
					safe = false
					break
				}
				prevSign = sign

			}
			if safe {
				numSafe++
			}
			// fmt.Printf("index: %v, element: %v\n", i, element)

		}
	}
	return numSafe
}

func puzzle2(reports []string) int {

	numSafe := 0
	for _, report := range reports {
		// replace spaces in element with a comma
		levels := strings.Split(report, " ")
		// convert the string to an int
		// https://pkg.go.dev/strconv#Atoi
		// int, err := strconv.Atoi(element)
		// append the int to the finalArr
		safe := true
		prevSign := 0
		sign := 0
		if report != "" {
			for i := 1; i < len(levels); i++ {
				prev, _ := strconv.Atoi(levels[i-1])
				curr, _ := strconv.Atoi(levels[i])
				diff := curr - prev
				// update sign if diff is positive
				if diff > 0 {
					sign = 1
				} else if diff < 0 {
					sign = -1
				}
				if diff > 3 || diff < -3 || diff == 0 || sign+prevSign == 0 {
					fmt.Printf("Report was not safe: %v,\nchecking if removing one will resolve issue\n", report)
					fmt.Printf("prev: %v, curr: %v, diff: %v\nprevious sign: %v, current sign: %v\n", prev, curr, diff, prevSign, sign)
					for j := 0; j < len(levels); j++ {
						safeWithTolerance := falutTolerance(j, report)
						if safeWithTolerance {
							fmt.Printf("Safe after removing index: %v\n", j)
							safe = true
							break
						}
						safe = false
					}

					if !safe {
						fmt.Printf("Still not safe\n\n")
					}
					break
				}
				prevSign = sign

			}
			if safe {
				numSafe++
			}
			// fmt.Printf("index: %v, element: %v\n", i, element)

		}
	}
	fmt.Printf("numSafe: %v\n", numSafe)
	fmt.Printf("numReports: %v\n", len(reports))
	return numSafe
}

func falutTolerance(levelOmitt int, report string) bool {
	levels := strings.Split(report, " ")
	levels = remove(levels, levelOmitt)
	safe := true
	prevSign := 0
	sign := 0

	for i := 1; i < len(levels); i++ {
		prev, _ := strconv.Atoi(levels[i-1])
		curr, _ := strconv.Atoi(levels[i])
		diff := curr - prev
		// update sign if diff is positive
		if diff > 0 {
			sign = 1
		} else if diff < 0 {
			sign = -1
		}
		if diff > 3 || diff < -3 || diff == 0 || sign+prevSign == 0 {
			fmt.Printf("\tprev: %v, curr: %v, diff: %v\n\tprevious sign: %v, current sign: %v\n", prev, curr, diff, prevSign, sign)
			return false
		}
		prevSign = sign

	}
	return safe
}

func remove(levels []string, i int) []string {
	if i == len(levels)-1 {
		return levels[:i]

	}
	if i == 0 {
		return levels[1:]
	}
	return append(levels[:i], levels[i+1:]...)
}
