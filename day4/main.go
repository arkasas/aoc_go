package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type field struct {
	number   string
	isActive bool
}

func open(filename string) *os.File {
	data, error := os.Open(filename)
	if error != nil {
		log.Fatal(error)
	}

	return data
}

func markNumberInBoards(boards [][][]field, number string) [][][]field {
	newBoards := boards
	for _, board := range newBoards {
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board[i]); j++ {
				if board[i][j].isActive == false {
					board[i][j].isActive = board[i][j].number == number
				}
			}
		}
	}

	return newBoards
}

func checkWinInBoard(board [][]field) bool {
	//Check horizontal
	for i := 0; i < len(board); i++ {
		win := board[i][0].isActive
		for j := 0; j < len(board[i]); j++ {
			win = board[i][j].isActive
			if board[i][j].isActive == false {
				win = false
				break
			}
		}
		if win == true {
			return true
		}
	}

	//Vertical
	for i := 0; i < len(board); i++ {
		win := board[0][i].isActive
		for j := 0; j < len(board[i]); j++ {
			win = board[j][i].isActive
			if board[j][i].isActive == false {
				win = false
				break
			}
		}
		if win == true {
			return true
		}
	}

	return false
}

func findSumInBoard(board [][]field) int {
	sum := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j].isActive == false {
				number, _ := strconv.Atoi(board[i][j].number)
				sum += number
			}
		}
	}

	return sum
}

func readFirstLine(scanner *bufio.Scanner) []string {
	scanner.Scan()
	return strings.Split(scanner.Text(), ",")
}

func readBoards(scanner *bufio.Scanner) [][][]field {
	scanner.Scan()

	boards := [][][]field{}
	singleBoard := [][]field{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			boards = append(boards, singleBoard)
			singleBoard = [][]field{}
		} else {
			line := strings.Join(strings.Fields(line), " ")
			arr := strings.Split(line, " ")
			fieldRow := []field{}
			for _, number := range arr {
				fieldRow = append(fieldRow, field{number, false})
			}
			singleBoard = append(singleBoard, fieldRow)
		}
	}
	boards = append(boards, singleBoard)
	return boards
}

func taskA(input string) int {
	scanner := bufio.NewScanner(open(input))
	numbers := readFirstLine(scanner)
	boards := readBoards(scanner)
	for _, number := range numbers {
		boards = markNumberInBoards(boards, string(number))
		for _, board := range boards {
			if checkWinInBoard(board) {
				sum := findSumInBoard(board)
				num, _ := strconv.Atoi(string(number))
				return sum * num
			}
		}
	}

	return 0
}

func contains(elems []int, value int) bool {
	for _, number := range elems {
		if number == value {
			return true
		}
	}

	return false
}

func taskB(input string) int {
	winingBoards := []int{}
	scanner := bufio.NewScanner(open(input))
	numbers := readFirstLine(scanner)
	boards := readBoards(scanner)
	for _, number := range numbers {
		boards = markNumberInBoards(boards, string(number))
		for index, board := range boards {
			if checkWinInBoard(board) {
				if contains(winingBoards, index) == false {
					winingBoards = append(winingBoards, index)
				}
				if len(winingBoards) == len(boards) {
					sum := findSumInBoard(board)
					num, _ := strconv.Atoi(string(number))
					return sum * num
				}
			}
		}
	}

	return 0
}

func testBoardWinCondition_notWin() {
	testBoard := [][]field{
		{field{"2", false}, field{"2", false}, field{"2", false}},
		{field{"2", false}, field{"2", false}, field{"2", false}},
		{field{"2", false}, field{"2", false}, field{"2", false}},
	}

	fmt.Println(checkWinInBoard(testBoard) == false)
}
func testBoardWinCondition_horizontalWin() {
	testBoard := [][]field{
		{field{"2", false}, field{"2", false}, field{"2", false}},
		{field{"2", false}, field{"2", false}, field{"2", false}},
		{field{"2", true}, field{"2", true}, field{"2", true}},
	}

	fmt.Println(checkWinInBoard(testBoard) == true)
}
func testBoardWinCondition_verticalWin() {
	testBoard := [][]field{
		{field{"2", false}, field{"2", false}, field{"2", true}},
		{field{"2", false}, field{"2", false}, field{"2", true}},
		{field{"2", false}, field{"2", false}, field{"2", true}},
	}

	fmt.Println(checkWinInBoard(testBoard) == true)
}

func runTaskA() {
	fmt.Println(taskA("test_input") == 4512)
	fmt.Println(taskA("input"))
}

func runTaskB() {
	fmt.Println(taskB("test_input") == 1924)
	fmt.Println(taskB("input"))
}

func main() {
	runTaskB()
}
