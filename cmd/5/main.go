package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	input_file, err := os.Open("cmd/5/input_real.txt")
	if err != nil {
		panic(err)
	}
	defer input_file.Close()
	scanner := bufio.NewScanner(input_file)

	jolts := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		sum := ""
		for i := 9; i > 0; i-- {
			loc := findLocation(strconv.Itoa(i), line)
			if loc != -1 && loc != len(line)-1 {
				sum += strconv.Itoa(i)
				for j := 9; j > 0; j-- {
					loc2 := findLocation(strconv.Itoa(j), line[loc+1:])
					if loc2 != -1 {
						sum += strconv.Itoa(j)
						break
					}
				}
				break
			}
		}
		jolts = append(jolts, sum)
		//fmt.Printf("Line: %s Sum: %d\n", line, sum)
	}
	totalJolts := 0
	for _, jolt := range jolts {
		j, _ := strconv.Atoi(jolt)
		totalJolts += j
	}
	fmt.Printf("Total jolts: %d\n", totalJolts)
}

func findLocation(pattern string, text string) int {
	//fmt.Printf("Pattern: %s\n", pattern)
	//fmt.Printf("Text: %s\n", text)
	r, _ := regexp.Compile(pattern)
	loc := r.FindStringIndex(text)
	if loc != nil {
		return loc[0]
	}
	return -1
}
