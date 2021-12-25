package main

import "fmt"

func insertSort(input_list []int, size int) []int {
	var keyword int
	var index int
	for i := 1; i < size; i++ {
		keyword = input_list[i]
		index = i - 1
		for index >= 0 && input_list[index] > keyword {
			input_list[index+1] = input_list[index]
			index--
		}
		input_list[index+1] = keyword
	}
	return input_list
}

func main() {
	input_list := []int{12, 222, 3, 42, 5}
	fmt.Println(insertSort(input_list, len(input_list)))
}
