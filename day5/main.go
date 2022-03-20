package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	Horizontal int = 0
	Vertical       = 1
	Diagonal       = 2
	Unknown        = 3
)

type vent struct {
	x1        int
	y1        int
	x2        int
	y2        int
	direction int
}

type coord struct {
	x int
	y int
}

func open(filename string) *os.File {
	data, error := os.Open(filename)
	if error != nil {
		log.Fatal(error)
	}

	return data
}

func parseToVent(line string) vent {
	splited := strings.Split(line, " -> ")
	val1 := strings.Split(splited[0], ",")
	val2 := strings.Split(splited[1], ",")

	x1, _ := strconv.Atoi(val1[0])
	y1, _ := strconv.Atoi(val1[1])
	x2, _ := strconv.Atoi(val2[0])
	y2, _ := strconv.Atoi(val2[1])

	return vent{x1, y1, x2, y2, Unknown}
}

func maxValueInVent(v vent) int {
	max := -1
	if v.x1 > max {
		max = v.x1
	}
	if v.x2 > max {
		max = v.x2
	}
	if v.y1 > max {
		max = v.y1
	}
	if v.y2 > max {
		max = v.y2
	}
	return max
}

func findMaxValueInVents(vents []vent) int {
	max := -1
	for _, vent := range vents {
		if maxValueInVent(vent) > max {
			max = maxValueInVent(vent)
		}
	}

	return max
}

func removeNotMatching(vents []vent) []vent {
	ventsToReturn := []vent{}
	for _, vent := range vents {
		if vent.x1 == vent.x2 {
			vent.direction = Horizontal
			ventsToReturn = append(ventsToReturn, vent)
		} else if vent.y1 == vent.y2 {
			vent.direction = Vertical
			ventsToReturn = append(ventsToReturn, vent)
		}
	}

	return ventsToReturn
}

func count(matrix [][]int, size int) int {
	sum := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if matrix[i][j] >= 2 {
				sum++
			}
		}
	}

	return sum
}

func findMathEq(v vent) {

}

func taskA(input string) int {
	scanner := bufio.NewScanner(open(input))
	vents := []vent{}
	for scanner.Scan() {
		vents = append(vents, parseToVent(scanner.Text()))
	}
	vents = removeX(vents)
	fmt.Println(vents)
	maxValue := findMaxValueInVents(vents) + 1
	matrix := make([][]int, maxValue)
	for i := range matrix {
		matrix[i] = make([]int, maxValue)
	}

	for i := 0; i < maxValue; i++ {
		for j := 0; j < maxValue; j++ {
			matrix[i][j] = 0
		}
	}

	for _, vent := range vents {
		if vent.direction == Horizontal {
			smaller := vent.y1
			higher := vent.y2
			if vent.y2 < vent.y1 {
				smaller = vent.y2
				higher = vent.y1
			}
			for i := smaller; i <= higher; i++ {
				matrix[i][vent.x1] += 1
			}
		} else if vent.direction == Vertical {
			smaller := vent.x1
			higher := vent.x2
			if vent.x2 < vent.x1 {
				smaller = vent.x2
				higher = vent.x1
			}
			for i := smaller; i <= higher; i++ {
				matrix[vent.y1][i] += 1
			}
		}
	}
	fmt.Println(matrix)

	return count(matrix, maxValue)
}

func removeX(vents []vent) []vent {
	ventsToReturn := []vent{}
	for _, vent := range vents {
		if vent.x1 == vent.x2 {
			vent.direction = Horizontal
			ventsToReturn = append(ventsToReturn, vent)
		} else if vent.y1 == vent.y2 {
			vent.direction = Vertical
			ventsToReturn = append(ventsToReturn, vent)
		} else if math.Abs(float64(vent.x1)-float64(vent.x2)) == math.Abs(float64(vent.y1)-float64(vent.y2)) {
			vent.direction = Diagonal
			ventsToReturn = append(ventsToReturn, vent)
		}
	}

	return ventsToReturn
}

func pointsForVent(v vent) []coord {
	a := v.y1 - v.y2
	b := v.x2 - v.x1
	c := (v.x1-v.x2)*v.y1 + (v.y2-v.y1)*v.x1
	fmt.Println(a, b, c)

	startX := v.x1
	endX := v.x2
	if (v.x1*v.x1 + v.y1 + v.y1) > (v.x2*v.x2 + v.y2 + v.y2) {
		startX = v.x2
		endX = v.x1
	}

	output := []coord{}
	for i := startX; i <= endX; i++ {
		yVal := (a*i + c) / (-b)
		output = append(output, coord{x: i, y: yVal})
	}

	return output
}

func taskB(input string) int {
	scanner := bufio.NewScanner(open(input))
	vents := []vent{}
	for scanner.Scan() {
		vents = append(vents, parseToVent(scanner.Text()))
	}
	vents = removeX(vents)
	fmt.Println(vents)
	maxValue := findMaxValueInVents(vents) + 1
	matrix := make([][]int, maxValue)
	for i := range matrix {
		matrix[i] = make([]int, maxValue)
	}

	for i := 0; i < maxValue; i++ {
		for j := 0; j < maxValue; j++ {
			matrix[i][j] = 0
		}
	}

	for _, vent := range vents {
		if vent.direction == Horizontal {
			smaller := vent.y1
			higher := vent.y2
			if vent.y2 < vent.y1 {
				smaller = vent.y2
				higher = vent.y1
			}
			for i := smaller; i <= higher; i++ {
				matrix[i][vent.x1] += 1
			}
		} else if vent.direction == Vertical {
			smaller := vent.x1
			higher := vent.x2
			if vent.x2 < vent.x1 {
				smaller = vent.x2
				higher = vent.x1
			}
			for i := smaller; i <= higher; i++ {
				matrix[vent.y1][i] += 1
			}
		} else if vent.direction == Diagonal {
			poi := pointsForVent(vent)
			for _, p := range poi {
				matrix[p.y][p.x] += 1
			}
		}
	}

	return count(matrix, maxValue)
}

func runTaskA() {
	fmt.Println(taskA("test_input") == 5)
	fmt.Println(taskA("input"))
}

func runTaskB() {
	fmt.Println(taskB("test_input") == 12)
	fmt.Println(taskB("input"))
}

func main() {
	runTaskB()
}
