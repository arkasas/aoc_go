package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func open(filename string) *os.File {
	data, error := os.Open(filename)
	if error != nil {
		log.Fatal(error)
	}

	return data
}

func runTask(input string) int {
	scanner := bufio.NewScanner(open(input))
	counter := 0
	scanner.Scan()
	previousValue, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		intValue, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		if intValue > previousValue {
			counter++
		}
		previousValue = intValue
	}

	return counter
}

func main() {
	fmt.Println(runTask("test_input") == 7)
	fmt.Println(runTask("input"))
}
