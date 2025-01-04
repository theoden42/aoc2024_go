package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("disk_fragmenter.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	final_ans := int64(0)
	for scanner.Scan() {
		line := scanner.Text()
		new_line := []byte(line) // can't mutate strings in go
		prefix_sums := make([]int64, len(new_line)+1)
		prefix_sums[0] = 0
		for i := 0; i < len(new_line); i++ {
			prefix_sums[i+1] = prefix_sums[i] + int64(new_line[i]-'0')
		}
		for j := len(new_line) - 1; j >= 0; j-- {
			if j%2 == 0 {
				for i := 0; i < j; i++ {
					if i%2 == 1 && new_line[i] >= new_line[j] {
						sum_pos := prefix_sums[i]*int64(new_line[j]-'0') + int64(new_line[j]-'0')*(int64(new_line[j]-'0')-1)/2
						fmt.Println(sum_pos, new_line[j]-'0', new_line[i]-'0', i, j/2)
						final_ans += int64(j/2) * sum_pos
						new_line[i] -= new_line[j]
						new_line[i] += '0'
						prefix_sums[i] += int64(new_line[j] - '0')
						new_line[j] = '0'
						fmt.Println(final_ans)
						break
					}
				}
			}
		}

		for i := 0; i < len(new_line); i++ {
			if i%2 == 0 && new_line[i] != '0' {
				sum_pos := prefix_sums[i]*int64(new_line[i]-'0') + int64(new_line[i]-'0')*(int64(new_line[i]-'0')-1)/2
				final_ans += int64(i/2) * sum_pos
			}
		}
		fmt.Println(final_ans)
	}

}
