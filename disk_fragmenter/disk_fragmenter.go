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
	// final_ans := 0
	for scanner.Scan() {
		line := scanner.Text()
		j := len(line) - 1
		if len(line)%2 == 0 {
			j -= 1
		}
		new_line := []byte(line) // can't mutate strings in go
		// the string is so short that we can just brute force it
		// otherwise use two pointers
		id_string := []int64{}
		for i := 0; i < len(new_line) && i <= j; i++ {
			if i%2 == 0 {
				for k := 0; k < int(new_line[i]-'0'); k++ {
					id_string = append(id_string, int64(i)/2)
				}
			} else {
				if j > i {
					for k := 0; k < min(int(new_line[i]-'0'), int(new_line[j]-'0')); k++ {
						id_string = append(id_string, int64(j)/2)
					}
					if new_line[i] > new_line[j] {
						new_line[i] -= new_line[j]
						new_line[i] += '0'
						i -= 1
						j -= 2
					} else {
						new_line[j] -= new_line[i]
						new_line[j] += '0'
					}
				}
			}
		}

		final_ans := int64(0)
		for i := 0; i < len(id_string); i++ {
			final_ans += id_string[i] * int64(i)
		}
		fmt.Println(final_ans)
	}

}
