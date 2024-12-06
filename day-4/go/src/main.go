package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ReadFile() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var wordSearch = [][]string{}
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "")
		wordSearch = append(wordSearch, row)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	println("XMAS count: ", CrossMasSearch(wordSearch))
}

func CrossMasSearch(wordSearch [][]string) int {
	hight := 0
	for _, _ = range wordSearch {
		hight++
	}
	count := 0
	for x, row := range wordSearch {
		for y, cell := range row {
			if cell == "A" {
				if x+1 < hight && x-1 >= 0 && y+1 < hight && y-1 >= 0 {
					matches := 0
					if (wordSearch[x+1][y+1] == "M" && wordSearch[x-1][y-1] == "S") || (wordSearch[x+1][y+1] == "S" && wordSearch[x-1][y-1] == "M") {
						matches++
					}
					if (wordSearch[x-1][y+1] == "M" && wordSearch[x+1][y-1] == "S") || (wordSearch[x-1][y+1] == "S" && wordSearch[x+1][y-1] == "M") {
						matches++
					}
					if matches == 2 {
						count++
					}
				}
			}
		}
	}
	return count
}

func XmasSearch(wordSearch [][]string) int {
	transformations := [][]int{{-1, -1}, {0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}}
	hight := 0
	for _, _ = range wordSearch {
		hight++
	}
	count := 0
	for x, row := range wordSearch {
		for y, cell := range row {
			if cell == "X" {
				for _, transformation := range transformations {
					if x+(transformation[0]*3) < hight && x+(transformation[0]*3) >= 0 && y+(transformation[1]*3) < hight && y+(transformation[1]*3) >= 0 {
						if wordSearch[x+transformation[0]][y+transformation[1]] == "M" && wordSearch[x+(transformation[0]*2)][y+(transformation[1]*2)] == "A" && wordSearch[x+(transformation[0]*3)][y+(transformation[1]*3)] == "S" {
							count++
						}
					}
				}
			}
		}
	}
	return count
}

func main() {
	ReadFile()
}
