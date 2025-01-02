package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("bridge_repair.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var final_value int64 = 0
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, ":")
		if err != nil {
			log.Fatal(err)
		}
		expected_ans, _ := strconv.Atoi(values[0])
		values = strings.Split(values[1], " ")
		values = values[1:]
		possible_table := make(map[int]bool)
		for _, value := range values {
			current_value, _ := strconv.Atoi(value)
			new_possible_table := make(map[int]bool)
			if len(possible_table) == 0 {
				possible_table[current_value] = true
				continue
			}
			for key := range possible_table {
				new_possible_table[key*current_value] = true
				new_possible_table[key+current_value] = true
				concatenated_value := strconv.Itoa(key) + strconv.Itoa(current_value)
				concatenated_value_int, _ := strconv.Atoi(concatenated_value)
				new_possible_table[concatenated_value_int] = true
			}
			possible_table = new_possible_table
		}
		value, exists := possible_table[expected_ans]
		if exists && value {
			// fmt.Println(expected_ans, values)
			final_value += int64(expected_ans)
		}
	}
	fmt.Println(final_value)
}
