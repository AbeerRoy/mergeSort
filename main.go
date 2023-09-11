package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/AbeerRoy/merge_sort"
)

func main() {

	x := input([]int64{}, nil)
	//fmt.Println(x)

	fmt.Println("\n--- Unsorted --- \n\n", x)
	fmt.Println("\n--- Sorted ---\n\n", merge_sort.MergeSort(x))
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
