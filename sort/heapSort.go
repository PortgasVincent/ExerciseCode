package sort

import "fmt"

func swap(a int, b int) (int, int) {
	return b, a
}

//func Adjust(val []int, i int, length int) {
//	temp := val[i]
//	for k := i*2 + 1; k < length; k = k*2 + 1 {
//		if k+1 < length && val[k+1] > val[k] {
//			k++
//		}
//		if val[k] > temp {
//			val[i] = val[k]
//			i = k
//		}
//	}
//	val[i] = temp
//}
//
//func HeapSort(val []int) {
//	length := len(val)
//	if length <= 1 {
//		return
//	}
//	for i := length/2 - 1; i >= 0; i-- {
//		Adjust(val, i, length)
//	}
//	for i := length - 1; i >= 0; i-- {
//		val[0], val[i] = swap(val[0], val[i])
//		Adjust(val, 0, i)
//	}
//}
//
//func adjust(arr []int, p, length int){
//	temp := arr[p]
//	for k := 2*p+1; k<length; k = 2*k+1{
//		if k+1<length && arr[k+1]>arr[k]{
//			k++
//		}
//		if arr[k] > temp{
//			arr[p] = arr[k]
//			p = k
//		}
//	}
//	arr[p] = temp
//
//}
//
//func HeapSort(arr []int){
//	l := len(arr)
//	if l <= 1{
//		return
//	}
//	for i := l/2 - 1; i >= 0; i--{
//		adjust(arr, i, l-1)
//	}
//	for j := l-1; j >= 0; j--{
//		arr[0], arr[j] = swap(arr[0], arr[j])
//		adjust(arr, 0, j)
//	}
//}

func main() {
	arr := []int{1, 5, 8, 56, 4, 6, 3, 9, 45}
	HeapSort(arr)
	fmt.Println(arr)
}

func HeapSort(arr []int){
	length := len(arr)
	if length <= 1{
		return
	}
	for i := length/2-1;i >= 0;i--{
		adjust(arr, i, length)
	}
	for  i:= length -1; i >= 0; i--{
		arr[0], arr[i] = swap(arr[0], arr[i])
		adjust(arr, 0, i)
	}
}

func adjust(arr []int, i, l int){
	temp := arr[i]
	for k := 2*i+1; k < l; k = 2*k+1{
		if k+1 < l && arr[k+1]>arr[k]{
			k += 1
		}
		if arr[k]>temp{
			arr[i] = arr[k]
			i = k
		}

	}
	arr[i] = temp
}