package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// inputFile := "example.txt"
	inputFile := "input.txt"
	// part1(inputFile)
	part2(inputFile)
}

func parseInt(stringNumber string) int {
	num, err := strconv.Atoi(stringNumber)
	if err != nil {
		fmt.Println("Error converting string to int, ", err)
		panic(err)
	}
	return num
}

func part1(filepath string) {
	readFile, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error reading file, ", err)
	}

	filescanner := bufio.NewScanner(readFile)

	filescanner.Split(bufio.ScanLines)

	lines := []string{}
	for filescanner.Scan() {
		line := filescanner.Text()
		lines = append(lines, line)
	}

	readFile.Close()

	left := []int{}
	right := []int{}
	for _, line := range lines {
		split := strings.Split(line, "   ")
		left = append(left, parseInt(split[0]))
		right = append(right, parseInt(split[1]))
	}
	sort.Sort(sort.IntSlice(left))
	sort.Sort(sort.IntSlice(right))

	distances := []int{}
	for i := 0; i < len(left); i++ {
		distance := int(math.Abs(float64(left[i] - right[i])))
		distances = append(distances, distance)
	}

	sum := 0
	for _, distance := range distances {
		sum += distance
	}

	fmt.Println("Part 1: ", sum)
}

func part2(filepath string) {
	readFile, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error reading file, ", err)
	}

	filescanner := bufio.NewScanner(readFile)

	filescanner.Split(bufio.ScanLines)

	lines := []string{}
	for filescanner.Scan() {
		line := filescanner.Text()
		lines = append(lines, line)
	}

	readFile.Close()

	left := []int{}
	right := []int{}
	for _, line := range lines {
		split := strings.Split(line, "   ")
		left = append(left, parseInt(split[0]))
		right = append(right, parseInt(split[1]))
	}
	sort.Sort(sort.IntSlice(left))
	sort.Sort(sort.IntSlice(right))

	products := []int{}
	for i := 0; i < len(left); i++ {
		currSum := 0
		for j := 0; j < len(right); j++ {
			if left[i] != right[j] {
				continue
			}
			currSum += left[i]
		}
		products = append(products, currSum)
	}

	sum := 0

	for _, product := range products {
		sum += product
	}

	fmt.Println("Part 1: ", sum)
}
