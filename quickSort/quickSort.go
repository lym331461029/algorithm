package main

import "fmt"

func quickSort(arr []int, start int, end int) []int {
	if start+1 >= end {
		return arr
	}

	if start+2 == end {
		if arr[start] > arr[start+1] {
			arr[start], arr[start+1] = arr[start+1], arr[start]
		}
		return arr
	}

	center := (start + end) / 2
	if arr[start] > arr[center] {
		arr[start], arr[center] = arr[center], arr[start]
	}

	if arr[start] > arr[end-1] {
		arr[start], arr[end-1] = arr[end-1], arr[start]
	}

	if arr[center] > arr[end-1] {
		arr[center], arr[end-1] = arr[end-1], arr[center]
	}

	arr[end-2], arr[center] = arr[center], arr[end-2]
	i := start + 1
	j := end - 3
	for {

		for ; arr[j] > arr[end-2]; j-- {

		}
		for ; arr[i] < arr[end-2]; i++ {

		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		} else {
			break
		}
	}

	arr[i], arr[end-2] = arr[end-2], arr[i]
	quickSort(arr, start, i)
	quickSort(arr, i+1, end)
	return arr
}

func main() {
	a := []int{3, 9, 20, 21, 7, 4, 58, 13}
	ax := quickSort(a, 0, 8)
	fmt.Println(ax)
	b := []int{7, 4, 5}
	quickSort(b, 0, 3)
	fmt.Println(b)

	c := []int{7, 6, 5, 4}
	quickSort(c, 0, 4)
	fmt.Println(c)

}
