package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dlclark/regexp2"
)

func main() {
	input_file, err := os.Open("cmd/4/input_real.txt")
	if err != nil {
		panic(err)
	}
	defer input_file.Close()

	csvReader := csv.NewReader(input_file)
	records, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	sum := 0
	for _, record := range records[0] {
		first, err := strconv.Atoi(strings.Split(record, `-`)[0])
		if err != nil {
			panic(err)
		}
		second, err := strconv.Atoi(strings.Split(record, `-`)[1])
		if err != nil {
			panic(err)
		}
		for i := first; i <= second; i++ {
			re := regexp2.MustCompile(`^(\d+)\1+$`, 0)
			match, _ := re.MatchString(strconv.Itoa(i))
			if match {
				sum += i
			}
		}
	}
	fmt.Printf("Sum: %d\n", sum)
}
