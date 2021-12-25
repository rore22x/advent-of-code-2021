package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	Day1Part1()
	Day1Part2()
}

func Day1Part2() {
	file, err := ioutil.ReadFile("day-1.txt")
	if err != nil {
		log.Fatal(err)
	}
	splitted := strings.Split(string(file), "\n")

	vals := make([]int, 0)
	var increaseCounter = 0
	for index, line := range splitted {
		value, _ := strconv.Atoi(line)
		vals = append(vals, value)
		if index > 2 {
			lastValue := vals[index-1] + vals[index-2] + vals[index-3]
			currentValue := vals[index] + vals[index-1] + vals[index-2]
			if currentValue > lastValue {
				increaseCounter++
			}
		}
	}
	fmt.Printf("Result: %d\n ", increaseCounter)
}

func Day1Part1() {
	file, err := ioutil.ReadFile("day-1.txt")
	if err != nil {
		log.Fatal(err)
	}
	splitted := strings.Split(string(file), "\n")
	var increaseCounter = 0
	var previous = 0xFFFFFFFF
	for index, i := range splitted {
		value, _ := strconv.Atoi(i)
		if index > 0 && previous < value {
			increaseCounter++
		}
		previous = value
	}
	fmt.Printf("Result: %d\n", increaseCounter)
}
