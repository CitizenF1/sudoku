package main

import (
	"fmt"
)

var sudoku = [9][9]int{
	{8, 0, 7, 0, 0, 0, 0, 0, 0},
	{0, 3, 1, 0, 0, 2, 4, 0, 0},
	{0, 4, 0, 0, 0, 0, 0, 5, 2},
	{9, 6, 0, 4, 1, 0, 8, 7, 0},
	{1, 0, 0, 7, 0, 3, 9, 2, 0},
	{0, 0, 4, 9, 0, 8, 1, 0, 0},
	{4, 0, 6, 1, 0, 7, 2, 3, 0},
	{7, 5, 3, 0, 0, 0, 0, 9, 1},
	{0, 1, 0, 0, 0, 6, 5, 0, 0},
}

// var a [9][9]int = sudoku

func draw() {
	for _, row := range sudoku {
		fmt.Println(row)
	}
}
func canPut(x int, y int, value int) bool {
	return !alredyInVertical(x, y, value) &&
		!alredyInHorizontal(x, y, value) &&
		!alredyInSqare(x, y, value)
}

func alredyInVertical(x int, y int, value int) bool {
	for i := range [9]int{} {
		if sudoku[i][x] == value {
			return true
		}
	}
	return false
}

func alredyInHorizontal(x int, y int, value int) bool {
	for i := range [9]int{} {
		if sudoku[y][i] == value {
			return true
		}
	}
	return false
}

func alredyInSqare(x int, y int, value int) bool {
	sx, sy := int(x/3)*3, int(y/3)*3
	for dy := range [3]int{} {
		for dx := range [3]int{} {
			if sudoku[sy+dy][sx+dx] == value {
				return true
			}
		}
	}
	return false
}

func next(x int, y int) (int, int) {
	nextX, nextY := (x+1)%9, y
	if nextX == 0 {
		nextY = y + 1
	}
	return nextX, nextY
}

func solve(x int, y int) bool {
	if y == 9 {
		return true
	}
	if sudoku[y][x] != 0 {
		return solve(next(x, y))
	} else {
		for i := range [9]int{} {
			var v = i + 1
			if canPut(x, y, v) {
				sudoku[y][x] = v
				if solve(next(x, y)) {
					return true
				}
				sudoku[y][x] = 0
			}
		}
		return false
	}
}

func main() {
	draw()
	if solve(0, 0) {
		fmt.Println("Found")
		draw()
	} else {
		fmt.Println("Not found")
	}
}
