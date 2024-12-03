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

		// fmt.Println(scanner.Text()[0:5], scanner.Text()[5+3:])
		// fmt.Println(scanner.Text())
	}
	slices.Sort(col1)
	slices.Sort(col2)
	var distance int
	fmt.Println(len(col1))

	for i := 0; i < len(col1); i++ {
		distance += abs(col1[i] - col2[i])

	}
	fmt.Println(distance)

}
