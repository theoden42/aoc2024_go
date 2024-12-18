package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
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
	sort.Ints(first_list)
	sort.Ints(second_list)
	for i := 0; i < len(first_list); i++ {
		final_ans += int(math.Abs(float64(first_list[i] - second_list[i])))
	}

	fmt.Println(final_ans)

}
