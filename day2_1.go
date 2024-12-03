package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isIncrease(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] >= arr[i+1] {
			return false
		}
	}
	return true
}
func isDecrease(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] <= arr[i+1] {
			return false
		}

	}
	return true
}

func main() {
	file, err := os.Open("problem/day2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	// var content string
	// fmt.Println(scanner.Text())
	var safe_count int
	for scanner.Scan() {
		line := scanner.Text()
		text := strings.Split(line, " ")
		var numbers []int
		for _, val := range text {
			converted, _ := strconv.Atoi(string(val))
			numbers = append(numbers, converted)

		}
		if numbers[0] < numbers[1] {
			if !isIncrease(numbers) {
				continue
			}
			flag := false
			for i := 0; i < len(numbers)-1; i++ {
				if numbers[i+1] > numbers[i]+3 {
					flag = true
					break
				}
			}
			if flag {
				continue
			}
		} else if numbers[0] > numbers[1] {
			// fmt.Println(numbers)
			if !isDecrease(numbers) {
				continue
			}
			flag := false
			for i := 0; i < len(numbers)-1; i++ {
				if numbers[i+1] < numbers[i]-3 {
					flag = true
					break
				}
			}
			if flag {
				continue
			}
		} else if numbers[0] == numbers[1] {
			continue
		}
		// fmt.Println(numbers)
		safe_count++
		// line++
	}
	fmt.Println(safe_count)

}
