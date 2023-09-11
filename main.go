package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	x := input([]int64{}, nil)
	//fmt.Println(x)

	fmt.Println("\n--- Unsorted --- \n\n", x)
	fmt.Println("\n--- Sorted ---\n\n", MergeSort(x))
}

func input(inputArray []int64, err error) []int64 {
	if err != nil {
		return inputArray
	}
	var value string
	fmt.Println("Please enter the integer digits you want to include \n[Seperate each digit with COMMA(,) as delimiter]:")
	fmt.Scan(&value)
	values := strings.Split(value, ",")

LOOP:
	// converting string to int64 and/or taking error 'err' as input
	for i := 0; i < len(values); i++ {
		n, err := strconv.ParseInt(values[i], 10, 64)

		if err != nil {

			fmt.Println(err.Error(), " INVALID INPUT")
			//log.Fatalf("\nstrconv.Atoi() failed with %s \n", err)
			//fmt.Println("INVALID INPUT")
			fmt.Println("Please enter VALID integer digits \nSeperate each digit with COMMA(,)as delimiter]:")
			fmt.Scan(&value)
			values = strings.Split(value, ",")

			goto LOOP
		}
		inputArray = append(inputArray, n)
	}
	return inputArray

}

// Splitting the user input inputArray in to two arrays
func MergeSort(items []int64) []int64 {
	var num = int64(len(items))

	if num == 1 {
		return items
	}

	middle := int64(num / 2)
	var (
		left  = make([]int64, middle)
		right = make([]int64, num-middle)
	)
	for i := int64(0); i < num; i++ {
		if i < middle {
			left[i] = items[i]
		} else {
			right[i-middle] = items[i]
		}
	}

	return merge(MergeSort(left), MergeSort(right))
}

// joins the two sorted arrays to one single array
func merge(left, right []int64) []int64 {
	result := make([]int64, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return result
}
