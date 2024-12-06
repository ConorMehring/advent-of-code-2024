package main

import (
	"os"
	"strconv"
	"strings"
)

func ReadFile() {
	fileName := "../input.txt"
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	result := strings.Split(string(data), "\n")

	var left []int
	var right []int
	for _, row := range result {
		ietm := strings.Split(row, "   ")
		l, err := strconv.Atoi(ietm[0])
		if err != nil {
			panic(err)
		}
		r, err := strconv.Atoi(ietm[1])
		if err != nil {
			panic(err)
		}
		left = append(left, l)
		right = append(right, r)
	}

	left = MergeSort(left)
	right = MergeSort(right)

	println("similarity: ", similarity(left, right))
	// var distance = 0
	// for index, l := range left {
	// 	r := right[index]
	// 	if r < l {
	// 		distance += l - r
	// 	} else {
	// 		distance += r - l
	// 	}
	// }

	// println("Distance: ", distance)
}

func MergeSort(input []int) []int {
	if len(input) < 2 {
		return input
	}
	left := MergeSort(input[:len(input)/2])
	right := MergeSort(input[len(input)/2:])
	return merge(left, right)
}

func merge(left []int, right []int) []int {
	var final = []int{}
	l := 0
	r := 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			final = append(final, left[l])
			l++
		} else {
			final = append(final, right[r])
			r++
		}
	}
	for ; l < len(left); l++ {
		final = append(final, left[l])
	}
	for ; r < len(right); r++ {
		final = append(final, right[r])
	}
	return final
}

func similarity(left []int, right []int) int {
	totalSimilarity := 0
	for _, l := range left {
		count := 0
		for _, r := range right {
			if r == l {
				count++
			}
		}
		totalSimilarity += l * count
	}
	return totalSimilarity
}

func main() {
	ReadFile()
}
