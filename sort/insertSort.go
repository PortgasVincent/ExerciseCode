package main

import "fmt"

// func InsertSort(values []int) {
// 	length := len(values)
// 	if length <= 1 {
// 		return
// 	}

// 	for i := 1; i < length; i++ {
// 		tmp := values[i] // 从第二个数开始，从左向右依次取数
// 		key := i - 1     // 下标从0开始，依次从左向右

// 		// 每次取到的数都跟左侧的数做比较，如果左侧的数比取的数大，就将左侧的数右移一位，直至左侧没有数字比取的数大为止
// 		for key >= 0 && tmp < values[key] {
// 			values[key+1] = values[key]
// 			key--
// 			//fmt.Println(values)
// 		}

// 		// 将取到的数插入到不小于左侧数的位置
// 		if key+1 != i {
// 			values[key+1] = tmp
// 		}
// 		//fmt.Println(values)
// 	}
// }

func main() {
	arr := []int{3, 2, 23, 45, 4, 6, 3, 4, 9, 5, 3, 2, 1, 1, 6, 6}
	InsertSort(arr)
	fmt.Println(arr)
}

func InsertSort(arr []int) {
	length := len(arr)
	for i := 1; i < length; i++ {
		temp := arr[i]
		key := i - 1
		for key >= 0 && arr[key] > temp {
			arr[key+1] = arr[key]
			key--
		}
		if key+1 != i {
			arr[key+1] = temp
		}
	}
	return
}
