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

func check_level_is_good(numbers []int) bool {
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
		return false
	}
	flag := true
	for i := 1; i < len(numbers); i++ {
		if numbers[i] == numbers[i-1] || int(math.Abs(float64(numbers[i]-numbers[i-1]))) > 3 {
			flag = false
		}
	}
	return flag
}

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
		// there is O(n) way to do this where you can just count the incorrect positions in a set
		// and at max one is checked, however it is two much of a hassle

		for i := 0; i < len(numbers); i++ {
			var temp_numbers []int
			temp_numbers = append(temp_numbers, numbers[:i]...)
			temp_numbers = append(temp_numbers, numbers[i+1:]...)
			fmt.Println(temp_numbers)
			if check_level_is_good(temp_numbers) {
				count_lines++
				break
			}
		}
		// break
	}

	fmt.Print(count_lines)
}
