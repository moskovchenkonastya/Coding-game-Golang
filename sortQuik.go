package main

import "fmt"

func quikSort(arr []int, left int, right int) []int{

	n := len(arr)
	l := 0
	r := n - 1
	mid := arr[(l + r) / 2 ]

	for l <= r {

		for arr[r] > mid {
			r--
		}

		for arr[l] < mid {
			l++
		}

		if l <= r {
			arr[r],arr[l] = arr[l],arr[r]
			l++
			r--
		}
	}

	if r > left{
		quikSort(arr, left, r)
	}

	if l > right{
		quikSort(arr, l, right)
	}


	return arr
}


func main (){

	arr := []int{1, 34, 67, 987, 9, 765, 4, 7, 9}

	arr = quikSort(arr, 0, (len(arr) - 1))

	fmt.Println(arr)
}


