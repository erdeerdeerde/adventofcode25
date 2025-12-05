package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input_file, err := os.Open("cmd/2/input_real.txt")
	if err != nil {
		panic(err)
	}
	defer input_file.Close()
	scanner := bufio.NewScanner(input_file)
	dial := 50
	times := 0
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		if line[0] == 'R' {
			times = times + (dial+num)/100
			dial = (dial + num) % 100
			if dial > 100 {
				dial = dial - 100
			}
		} else {
			times = times + ((100-dial)+num)/100
			if dial == 0 {
				times = times - 1
			}
			dial = (dial - num) % 100
			if dial < 0 {
				dial = dial + 100
			}
		}
	}
	fmt.Printf("password: %d\n", times)
}
