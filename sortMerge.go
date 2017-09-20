package main

import "fmt"

func sortMerge(left []int, right [] int) []int {

	lst := make([]int, 0)
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0]{
			lst = append(lst, left[0])
			left = left[1:]
		} else {
			lst = append(lst, right[0])
			right = right[1:]
		}

		if len(left) > 0{
			lst = append(lst, left...)
		}

		if len(right) > 0{
			lst = append(lst, right...)
		}
	}

	return arr
}

func merge(arr []int) []int {

	if len(arr) >= 2 {
		mid := len(arr) / 2;
		arr = sortMerge(merge(arr[:mid]), merge(arr[mid:]))
	}

	return arr
}


func main() {

	arr := []int{1, 34, 67, 987, 9, 765, 4, 7, 9}

	arr = merge(arr)

	fmt.Println(arr)
}
