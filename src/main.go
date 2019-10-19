package main

import (
	"errors"
	"fmt"
	"sort"
)

func fibSlice(x int) ([]int, error) {
	fibArray := []int{}
	var newIndex int
	if x < 0 {
		return fibArray, errors.New("You cannot insert a negative number")
	}
	for i := 0; i < x+1; i++ {
		if i < 2 {
			newIndex = i
		} else {
			newIndex = fibArray[i-2] + fibArray[i-1]
		}
		fibArray = append(fibArray, newIndex)
	}
	return fibArray, nil
}
func revertSlice(slice []int) []int {
	revertedSlice := slice
	sort.Sort(sort.Reverse(sort.IntSlice(revertedSlice)))
	return revertedSlice
}
func findValue(slice []int, value int) (int, error) {
	pos := sort.SearchInts(slice, value)
	if pos < len(slice) && value == slice[pos] {
		return pos, nil
	} else {
		return 0, errors.New("value NOT found")
	}
}
func main() {
	var index int
	fmt.Println("Choose an index for the Fibonnaci sequence:")
	_, err := fmt.Scan(&index)
	fibresult, err := fibSlice(index)

	if err != nil {
		fmt.Println(err)
	} else {
		lastElement := fibresult[len(fibresult)-1]
		fmt.Println("The element you chose is:", lastElement)
		fmt.Println("Your sequence is the following:", fibresult)
	}
	fmt.Println("Your Fibonaci sequence reverted looks like this:", revertSlice(fibresult))

	var value int
	sort.Ints(fibresult)
	fmt.Println("Choose an value to search in your Fibonnaci sequence:")
	_, err = fmt.Scan(&value)
	findresult, err := findValue(fibresult, value)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value found in index:", findresult)
	}
}
