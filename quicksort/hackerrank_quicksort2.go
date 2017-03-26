package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	length := 0
	fmt.Scanf("%v", &length)
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		v := 0
		fmt.Scanf("%v", &v)
		arr[i] = v
	}
	quickSort(arr)
}

func quickSort(arr []int) {
	pivot(arr)
}

func printArr(arr []int) {
	for _, v := range arr {
		fmt.Printf("%v ", v)
	}
	fmt.Print("\n")
}

func pivot(arr []int) {

	greater := []int{}
	lesser := []int{}
	// select our pivot point
	p := arr[0]
	//split our arrays up as per hackerank question
	for i := 0; i < len(arr); i++ {
		if arr[i] > p {
			greater = append(greater, arr[i])
		} else if arr[i] < p {
			lesser = append(lesser, arr[i])
		}
	}
	// still not sorted if there is more than one val in any single array
	if len(lesser) > 1 {
		pivot(lesser)
	}
	if len(greater) > 1 {
		pivot(greater)
	}
	//once recursion is complete start reconstituing our array
	index := 0
	// lowest to highest with the pivot in the middle
	for i := 0; i < len(lesser); i++ {
		arr[index] = lesser[i]
		index++
	}
	arr[index] = p
	index++
	for i := 0; i < len(greater); i++ {
		arr[index] = greater[i]
		index++
	}
	printArr(arr)

}
