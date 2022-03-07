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

func count(value string, outputMap map[int][]int) {
	for i, letter := range value {
		intValue, _ := strconv.Atoi(string(letter))
		outputMap[i] = append(outputMap[i], intValue)
	}
}

func sum(array []int) int {
	sum := 0
	for _, value := range array {
		sum += value
	}

	return sum
}

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func reverseArray(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func removeElementsWhere(outputMap map[int][]int, prediction int, positionToRemove int) {
	indexToRemove := make([]int, 0)
	for i := 0; i < len(outputMap[positionToRemove]); i++ {
		if outputMap[positionToRemove][i] == prediction {
			indexToRemove = append(indexToRemove, i)
		}
	}

	indexToRemove = reverseArray(indexToRemove)
	for i := 0; i < len(outputMap); i++ {
		for j := 0; j < len(indexToRemove); j++ {
			outputMap[i] = RemoveIndex(outputMap[i], indexToRemove[j])
		}
	}
}

func taskA(input string) int64 {
	scanner := bufio.NewScanner(open(input))
	m := make(map[int][]int)
	for scanner.Scan() {
		count(scanner.Text(), m)
	}

	gammaBinary := ""
	epsilonBinary := ""
	for i := 0; i < len(m); i++ {
		if sum(m[i]) >= len(m[i])/2 {
			gammaBinary += string("1")
			epsilonBinary += string("0")
		} else {
			gammaBinary += string("0")
			epsilonBinary += string("1")
		}
	}

	gammaRate, _ := strconv.ParseInt(gammaBinary, 2, 64)
	epsilonRate, _ := strconv.ParseInt(epsilonBinary, 2, 64)
	return gammaRate * epsilonRate
}

func copyMap(originalMap map[int][]int) map[int][]int {
	targetMap := make(map[int][]int)
	for i := 0; i < len(originalMap); i++ {
		targetMap[i] = originalMap[i]
	}

	return targetMap
}

func taskB(input string) int64 {
	scanner := bufio.NewScanner(open(input))
	oxygenMap := make(map[int][]int)
	scrubberMap := make(map[int][]int)
	for scanner.Scan() {
		text := scanner.Text()
		count(text, oxygenMap)
		count(text, scrubberMap)
	}

	for i := 0; i < len(oxygenMap); i++ {
		if len(oxygenMap[i]) == 1 {
			break
		}
		if sum(oxygenMap[i]) >= len(oxygenMap[i])-sum(oxygenMap[i]) {
			removeElementsWhere(oxygenMap, 0, i)
		} else {
			removeElementsWhere(oxygenMap, 1, i)
		}
	}

	for i := 0; i < len(scrubberMap); i++ {

		if len(scrubberMap[i]) == 1 {
			break
		}
		if sum(scrubberMap[i]) >= len(scrubberMap[i])-sum(scrubberMap[i]) {
			removeElementsWhere(scrubberMap, 1, i)
		} else {
			removeElementsWhere(scrubberMap, 0, i)
		}
	}

	oxygenBinary := ""
	scrubberBinary := ""

	for i := 0; i < len(oxygenMap); i++ {
		oxygenBinary += strconv.Itoa(oxygenMap[i][0])
		scrubberBinary += strconv.Itoa(scrubberMap[i][0])
	}

	oxygenRate, _ := strconv.ParseInt(oxygenBinary, 2, 64)
	scrubberRate, _ := strconv.ParseInt(scrubberBinary, 2, 64)
	return oxygenRate * scrubberRate
}

func runTaskA() {
	fmt.Println(taskA("test_input") == 198)
	fmt.Println(taskA("input"))
}

func runTaskB() {
	fmt.Println(taskB("test_input") == 230)
	fmt.Println(taskB("input"))
}

func main() {
	runTaskB()
}
