package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("red-nosed-reports.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var count_lines int = 0
	for scanner.Scan() {
		line := scanner.Text()
		numbers_string := strings.Split(line, " ")
		var numbers []int
		for _, number_string := range numbers_string {
			number, _ := strconv.Atoi(number_string)
			numbers = append(numbers, number)
		}
		var increasing bool = true
		var decreasing bool = true
		for i := 1; i < len(numbers); i++ {
			if numbers[i] < numbers[i-1] {
				increasing = false
			}
		}
		for i := 1; i < len(numbers); i++ {
			if numbers[i] > numbers[i-1] {
				decreasing = false
			}
		}
		if !increasing && !decreasing {
			continue
		}
		flag := true
		for i := 1; i < len(numbers); i++ {
			if numbers[i] == numbers[i-1] || int(math.Abs(float64(numbers[i]-numbers[i-1]))) > 3 {
				flag = false
			}
		}
		if flag {
			count_lines++
		}
	}

	fmt.Print(count_lines)
}
