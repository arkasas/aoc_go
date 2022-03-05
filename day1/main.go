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

func runTaskA(input string) int {
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

func runTaskB(input string) int {
	scanner := bufio.NewScanner(open(input))
	counter := 0
	scanner.Scan()
	valueA, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	valueB, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	valueC, _ := strconv.Atoi(scanner.Text())
	prevSum := valueA + valueB + valueC
	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		nextSum := valueB + valueC + value
		if nextSum > prevSum {
			counter++
		}
		prevSum = nextSum
		valueA = valueB
		valueB = valueC
		valueC = value
	}

	return counter
}

func taskA() {
	fmt.Println(runTaskA("test_input") == 7)
	fmt.Println(runTaskA("input"))
}

func taskB() {
	fmt.Println(runTaskB("test_input") == 5)
	fmt.Println(runTaskB("input"))
}

func main() {
	taskB()
}
