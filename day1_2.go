package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func abs(num int) int {
	if num < 0 {
		return num * -1
	}
	return num
}

func main() {
	// reading file section
	file, err := os.Open("problem/day1.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	var col1 []int
	var col2 []int
	for scanner.Scan() {
		col1_num, err := strconv.Atoi(scanner.Text()[0:5])
		if err != nil {
			panic(err)
		}
		col2_num, err := strconv.Atoi(scanner.Text()[5+3:])
		if err != nil {
			panic(err)
		}
		col1 = append(col1, col1_num)
		col2 = append(col2, col2_num)
	}
	// solution section
	// counting the variable
	value_count := map[int]int{}
	for _, val := range col2 {
		if slices.Contains(col1, val) {
			value_count[val]++
		}
	}
	// count similarity
	var similarity int
	for _, val := range col1 {
		occurences, ok := value_count[val] //check whether key exist in map
		if ok {
			similarity += val * occurences
		} else {
			continue
		}
	}
	fmt.Println(similarity)
}
