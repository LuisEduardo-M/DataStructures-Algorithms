package BubbleSort

import "fmt"

func BubbleSort(arr []int) []int {
	swap := true

	for swap {
		swap = false

		for i := 0; i < len(arr)-1; i++ {
			// fmt.Println(arr[i] > arr[i+1])
			if arr[i] > arr[i+1] {
				swap = true
				arr[i], arr[i+1] = arr[i+1], arr[i]
				fmt.Println(arr)
			}
		}
	}
	return arr
}

// func main() {
// 	arr := []int{6, 8, 1, 4, 10, 7, 8, 9, 3, 2, 5}

// 	r := BubbleSort(arr)
// 	fmt.Println(r)
// }

/*
iteration 0: 6, 8, 1, 4, 10, 7, 8, 9, 3, 2, 5
iteration 1: 6, 1, 4, 8, 7, 8, 9, 3, 2, 5, 10
iteration 2: 1, 4, 6, 7, 8, 8, 3, 2, 5, 9, 10
sorted list: 1, 2, 3, 4, 5, 6, 7, 8, 8, 9, 10
*/
