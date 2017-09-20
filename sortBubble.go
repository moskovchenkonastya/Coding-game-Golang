package main

import "fmt"

func sortBubble(arr []int) []int {

	length := len(arr)
	for i := 0; i < length; i++{
		for j := length - 1; j > i; j--{
			if(arr[j - 1] > arr[j]) {
				arr[j], arr[j - 1] = arr[j - 1], arr[j]
			}
		}
	}
	return arr
}


func main() {

	arr := []int{1, 34, 67, 987, 9, 765, 4, 7, 9}
	arr = sortBubble(arr)
	fmt.Println(arr)
}
