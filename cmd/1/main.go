package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fi, err := os.Open("cmd/1/input_real.txt")
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	scanner := bufio.NewScanner(fi)
	d := 50
	t := 0
	for scanner.Scan() {
		l := scanner.Text()
		i, err := strconv.Atoi(l[1:])
		if err != nil {
			panic(err)
		}
		if l[0] == 'R' {
			d = (d + i) % 100
			if d > 100 {
				d = d - 100
			}
		} else {
			d = (d - i) % 100
			if d < 0 {
				d = d + 100
			}
		}
		fmt.Printf("%d\n", d)
		if d == 0 {
			t = t + 1
		}
	}
	fmt.Printf("password: %d\n", t)
}
