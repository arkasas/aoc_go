package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const FORWARD = "forward"
const DOWN = "down"
const UP = "up"

func open(filename string) *os.File {
	data, error := os.Open(filename)
	if error != nil {
		log.Fatal(error)
	}

	return data
}

func taskA(input string) int {
	depthPosition := 0
	horizontalPosition := 0

	scanner := bufio.NewScanner(open(input))
	for scanner.Scan() {
		move := strings.Split(scanner.Text(), " ")
		value, _ := strconv.Atoi(move[1])
		switch move[0] {
		case FORWARD:
			horizontalPosition += value
		case DOWN:
			depthPosition += value
		case UP:
			depthPosition -= value
		}
	}
	return depthPosition * horizontalPosition
}

func taskB(input string) int {
	depthPosition := 0
	horizontalPosition := 0
	aim := 0

	scanner := bufio.NewScanner(open(input))
	for scanner.Scan() {
		move := strings.Split(scanner.Text(), " ")
		value, _ := strconv.Atoi(move[1])
		switch move[0] {
		case FORWARD:
			horizontalPosition += value
			depthPosition += aim * value
		case DOWN:
			aim += value
		case UP:
			aim -= value
		}
	}
	return depthPosition * horizontalPosition
}

func runTaskA() {
	fmt.Println(taskA("test_input") == 150)
	fmt.Println(taskA("input"))
}

func runTaskB() {
	fmt.Println(taskB("test_input") == 900)
	fmt.Println(taskB("input"))
}
