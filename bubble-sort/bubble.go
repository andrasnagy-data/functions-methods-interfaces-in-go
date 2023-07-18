package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var intStr string

	fmt.Println(
		"Enter max 10 comma separated integers (for example: 1,3,4,5): ")
	fmt.Scan(&intStr)
	strSlc := strings.Split(intStr, ",")

	intSlc := make([]int, 0)

	for _, numStr := range strSlc {

		integer, _ := strconv.Atoi(numStr)

		intSlc = append(intSlc, integer)

	}
	BubbleSort(intSlc)
	fmt.Println(intSlc)

}

func BubbleSort(slc []int) {
	for i := 0; i < len(slc)-1; i++ {
		for j := 0; j < len(slc)-i-1; j++ {
			Swap(slc, j)
		}
	}
}

func Swap(slc []int, idx int) {
	if slc[idx] > slc[idx+1] {
		slc[idx], slc[idx+1] = slc[idx+1], slc[idx]
	}
}
