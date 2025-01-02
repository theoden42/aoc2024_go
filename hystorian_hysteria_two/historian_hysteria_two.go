package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	file, err := os.Open("historian_hysteria.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var first_list, second_list []int
	scanner := bufio.NewScanner(file)
	var final_ans int = 0
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, "   ")
		number_1, _ := strconv.Atoi(numbers[0])
		number_2, _ := strconv.Atoi(numbers[1])
		first_list = append(first_list, number_1)
		second_list = append(second_list, number_2)
	}
	frequency := make(map[int]int)
	for _, num := range second_list {
		frequency[num]++
	}
	for _, num := range first_list {
		final_ans += num * frequency[num]
	}

	fmt.Println(final_ans)

}