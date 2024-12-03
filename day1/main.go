package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// given a txt file of string data, inputDay1.txt, read the file and print the data
	// to the console
	// use the ioutil package to read the file
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	// use the fmt package to print the data to the console
	singleOutput := string(data)
	// conver a string to a slice based on the newline character
	// https://pkg.go.dev/strings#Split
	arr := strings.Split(singleOutput, "\n")
	finalArrLeft := []int{}
	finalArrRight := []int{}
	for _, element := range arr {
		// replace spaces in element with a comma
		elements := strings.Split(element, "   ")
		// convert the string to an int
		// https://pkg.go.dev/strconv#Atoi
		// int, err := strconv.Atoi(element)
		// append the int to the finalArr
		if element != "" {
			val0, _ := strconv.Atoi(elements[0])
			val1, _ := strconv.Atoi(elements[1])
			finalArrLeft = append(finalArrLeft, val0)
			finalArrRight = append(finalArrRight, val1)
		}
		// fmt.Printf("index: %v, element: %v\n", i, element)

	}
	if len(finalArrLeft) != len(finalArrRight) {
		panic("finalArrLeft and finalArrRight are not the same length")
	}
	puzzle1(finalArrLeft, finalArrRight)
	puzzle2(finalArrLeft, finalArrRight)

}

func puzzle2(finalArrLeft []int, finalArrRight []int) {
	totalSimilarity := 0
	// create a map with an int key and int value
	values := make(map[int]int)
	for i := 0; i < len(finalArrRight); i++ {
		// check if the value is in the map
		value, ok := values[finalArrRight[i]]
		// if the value is in the map, increment the value by 1
		if ok {
			values[finalArrRight[i]] = value + 1
		} else {
			// if the value is not in the map, add the value to the map with a value of 1
			values[finalArrRight[i]] = 1
		}
	}
	for i := 0; i < len(finalArrLeft); i++ {
		// check if the value is in the map
		_, ok := values[finalArrLeft[i]]
		// if the value is in the map, increment the value by 1
		if ok {
			totalSimilarity += values[finalArrLeft[i]] * finalArrLeft[i]
		}
	}
	// print out full value of totalDistance without using scientific notation
	fmt.Printf("totalSimilarity: %v\n", totalSimilarity)
}
func puzzle1(finalArrLeft []int, finalArrRight []int) {
	// for each element in the slice, print it to the console

	// print out finalArrLeft
	// fmt.Printf("finalArrLeft: %v\n", finalArrLeft)
	finalArrLeft = mergeSort(finalArrLeft)
	finalArrRight = mergeSort(finalArrRight)
	totalDistance := 0.0
	for i := 0; i < len(finalArrLeft); i++ {
		// add to total distance the absolute value of the difference between the two elements
		totalDistance += math.Abs(float64(finalArrLeft[i] - finalArrRight[i]))
	}
	// print out full value of totalDistance without using scientific notation
	fmt.Printf("totalDistance: %f\n", totalDistance)
}

// implement merge sort
func mergeSort(arr []int) []int {
	// base case
	if len(arr) <= 1 {
		return arr
	}
	// find the middle index
	mid := len(arr) / 2
	// recursively call mergeSort on the left and right halves
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	// merge the left and right halves
	return merge(left, right)
}
func merge(left, right []int) []int {
	// create a new slice to hold the merged values
	merged := make([]int, 0, len(left)+len(right))
	// create two index variables to keep track of the current index of the left and right slices
	l, r := 0, 0
	// iterate over the left and right slices
	for l < len(left) && r < len(right) {
		// compare the current values of the left and right slices
		// append the smaller value to the merged slice
		if left[l] < right[r] {
			merged = append(merged, left[l])
			l++
		} else {
			merged = append(merged, right[r])
			r++
		}
	}
	// append the remaining values of the left and right slices to the merged slice
	merged = append(merged, left[l:]...)
	merged = append(merged, right[r:]...)
	// return the merged slice
	return merged
}
